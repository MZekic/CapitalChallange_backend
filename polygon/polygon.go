package polygon

import (
	"capital-challenge-server/models"
	"capital-challenge-server/utils"
	"context"
	"os"
	"time"

	polygon "github.com/polygon-io/client-go/rest"
	polygonModels "github.com/polygon-io/client-go/rest/models"
	"github.com/u2takey/go-utils/uuid"
	"golang.org/x/exp/slices"
)

var PolygonClient *polygon.Client

func StartPolygonClient() {
	c := polygon.New(os.Getenv("POLYGON_API_KEY"))
	PolygonClient = c
}

func GetCompanyInfoByTicker(ticker string) (models.Companies, error) {

	var res models.Companies

	params := polygonModels.GetTickerDetailsParams{Ticker: ticker}
	resP, err := PolygonClient.GetTickerDetails(context.Background(), &params)
	if err != nil {
		return res, err
	}

	res.Name = resP.Results.Name
	res.Ticker = resP.Results.Ticker
	res.LogoURL = resP.Results.Branding.LogoURL
	res.HomepageURL = resP.Results.HomepageURL
	res.Locale = resP.Results.Locale
	res.Market = resP.Results.Market
	res.PrimaryExchange = resP.Results.PrimaryExchange
	res.TotalEmployees = float32(resP.Results.TotalEmployees)
	res.Type = resP.Results.Type
	res.Description = resP.Results.Description
	res.CurrencyName = resP.Results.CurrencyName

	return res, nil

}

func GetDailyCompanyStockInfo() ([]models.CompanyStock, error) {
	var res []models.CompanyStock
	tickers := utils.Tickers()
	lastDay := time.Now().AddDate(0, 0, -2)
	year := lastDay.Year()
	month := lastDay.Month()
	day := lastDay.Day()

	params := polygonModels.GetGroupedDailyAggsParams{Locale: polygonModels.US, MarketType: polygonModels.Stocks, Date: polygonModels.Date(time.Date(year, month, day, 0, 0, 0, 0, time.Local))}
	resP, err := PolygonClient.GetGroupedDailyAggs(context.Background(), &params)
	if err != nil {
		return res, err
	}
	var isWantedTicker int
	if resP.ResultsCount > 0 {
		for _, data := range resP.Results {
			var companyStock models.CompanyStock
			companyStock.ID = uuid.NewUUID()
			companyStock.ClosePrice = float32(data.Close)
			companyStock.HighestPrice = float32(data.High)
			companyStock.LowestPrice = float32(data.Low)
			companyStock.NumberOfTransactions = int(data.Transactions)
			companyStock.OTC = data.OTC
			companyStock.OpenPrice = float32(data.Open)
			companyStock.Ticker = data.Ticker
			companyStock.TradingVolume = int(data.Volume)
			companyStock.VolumeWeightedAveragePrice = float32(data.VWAP)
			companyStock.Date = &lastDay

			isWantedTicker = slices.Index(tickers, companyStock.Ticker)

			if isWantedTicker >= 0 {
				res = append(res, companyStock)
			}
		}
	}
	return res, nil
}



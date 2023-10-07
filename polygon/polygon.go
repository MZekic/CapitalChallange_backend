package polygon

import (
	"capital-challenge-server/models"
	"context"
	"os"

	polygon "github.com/polygon-io/client-go/rest"
	polygonModels "github.com/polygon-io/client-go/rest/models"
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

package v1Companies

import (
	dbhelper "capital-challenge-server/dbHelper"
	"capital-challenge-server/polygon"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/u2takey/go-utils/uuid"
)

// GetCompanyInfo godoc
// @Summary      Get Company info
// @Description  get company info by Ticker
// @Tags         companies
// @Accept       json
// @Produce      json
// @Param        ticker   path      string  true  "Ticker"
// @Success      200  {object}  GetCompanyInfoResponse
// @Failure      400
// @Failure      404
// @Failure      500
// @Router       /companies/{ticker} [get]
func GetCompanyInfo(c *gin.Context) {
	var req GetCompanyInfoRequest
	req.Ticker = c.Param("ticker")

	var res GetCompanyInfoResponse
	companyStock, err := dbhelper.GetCompanyStockLatestInfoByTicker(req.Ticker)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	res.CompanyStock = companyStock

	company, err := dbhelper.GetCompanyInfoByTicker(req.Ticker)
	if err == sql.ErrNoRows {
		company, err := polygon.GetCompanyInfoByTicker(req.Ticker)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		company.ID = uuid.NewUUID()
		res.Company = company

		c.JSON(http.StatusOK, res)
		go dbhelper.InsertCompany(company)
		return
	} else if err != sql.ErrNoRows && err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	res.Company = company

	c.JSON(http.StatusOK, res)
}

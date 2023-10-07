package v1Companies

import (
	dbhelper "capital-challenge-server/dbHelper"
	"capital-challenge-server/polygon"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetCompanyInfo godoc
// @Summary      Get Company info
// @Description  get company info by Ticker
// @Tags         companies
// @Accept       json
// @Produce      json
// @Param        ticker   path      string  true  "Ticker"
// @Success      200  {object}  models.Companies
// @Failure      400
// @Failure      404
// @Failure      500
// @Router       /companies/{ticker} [get]
func GetCompanyInfo(c *gin.Context) {
	var req GetCompanyInfoRequest
	req.Ticker = c.Param("ticker")

	_, err := dbhelper.GetCompanyInfoByTicker(req.Ticker)
	if err == sql.ErrNoRows {
		company, err := polygon.GetCompanyInfoByTicker(req.Ticker)
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		//todo add to database company info

		c.JSON(http.StatusOK, company)
		return
	} else if err != sql.ErrNoRows && err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
}

/*
func GetCompanyInfo(c *gin.Context) {
	var req GetCompanyInfoRequest
	if err := c.BindJSON(&req); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}
	err := validateGetCompanyInfoRequest(req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
	}

	c.JSON(http.StatusOK, req)
}
*/

package v1Companies

import (
	dbhelper "capital-challenge-server/dbHelper"
	"capital-challenge-server/polygon"
	utils "capital-challenge-server/utils"
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
// @Success      200  {object}  models.Companies
// @Failure      400
// @Failure      404
// @Failure      500
// @Router       /companies/{ticker} [get]
func GetCompanyInfo(c *gin.Context) {
	var req GetCompanyInfoRequest
	req.Ticker = c.Param("ticker")

	res, err := dbhelper.GetCompanyInfoByTicker(req.Ticker)
	if err == sql.ErrNoRows {
		company, err := polygon.GetCompanyInfoByTicker(req.Ticker)
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			utils.Log(c, http.StatusInternalServerError)
			return
		}
		company.ID = uuid.NewUUID()

		c.JSON(http.StatusOK, company)
		go dbhelper.InsertCompany(company)
		utils.Log(c, http.StatusOK)
		return
	} else if err != sql.ErrNoRows && err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		utils.Log(c, http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, res)
	utils.Log(c, http.StatusOK)
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

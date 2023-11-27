package cron

import (
	"capital-challenge-server/dbHelper"
	"capital-challenge-server/polygon"
	"log"

	"github.com/robfig/cron/v3"
)
func GetDailyStocks(){

	c := cron.New()
	c.AddFunc("CRON_TZ=America/Los_Angeles 0 1 * * *",func () {res, err := polygon.GetDailyCompanyStockInfo()
		if err != nil {
			log.Println("Cron job error: polygon API error")
			return
		}

		err = dbHelper.BatchInsertCompanyStock(res)
		if err != nil {
			log.Println("Cron job error: Fail to insert to DB")
			return
		}} )
	
	c.Start()
}




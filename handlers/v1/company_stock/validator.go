package v1companystock

import (
	"fmt"
	"strings"
)

func validateBuyOrSellCompanyStocksRequest(req BuyOrSellCompanyStocksRequest) error {
	var errors []string
	if len(req.CompanyStockID) <= 0 {
		errors = append(errors, "company_stock_id is required")
	}

	if req.Quantity <= 0 {
		errors = append(errors, "quantity must be higher than 0")
	}

	if req.BuyOrSell != "buy" && req.BuyOrSell != "sell" {
		errors = append(errors, "buy or sell field must is required")
	}

	if len(errors) > 0 {
		return fmt.Errorf("invalid BuyCompanyStocks request: %s", errors)
	} else {
		return nil
	}
}

func validateBuyCompanyStocksRequest(req BuyCompanyStockRequest) error {
	var errors []string
	if len(req.CompanyStockID) <= 0 {
		errors = append(errors, "company_stock_id is required")
	}

	if req.Quantity <= 0 {
		errors = append(errors, "quantity must be higher than 0")
	}

	if len(errors) > 0 {
		return fmt.Errorf("invalid BuyCompanyStocks request: %s", errors)
	} else {
		return nil
	}
}

func validateSellCompanyStocksRequest(req SellCompanyStockRequest) error {
	var errors []string
	if len(req.CompanyStockID) <= 0 {
		errors = append(errors, "company_stock_id is required")
	}

	if req.Quantity <= 0 {
		errors = append(errors, "quantity must be higher than 0")
	}

	if len(errors) > 0 {
		return fmt.Errorf("invalid BuyCompanyStocks request: %s", strings.Join(errors, ", "))
	} else {
		return nil
	}
}

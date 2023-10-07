package v1Companies

import "fmt"

func validateGetCompanyInfoRequest(req GetCompanyInfoRequest) error {
	var errors string
	if len(req.Ticker) <= 0 {
		errors = "ticker is missing"
	}

	if len(errors) > 0 {
		return fmt.Errorf("invalid GetCompanyInfo request: %s", errors)
	} else {
		return nil
	}
}

package v1Users

import "fmt"

func validateUserRegistrationRequest(req UserRegistrationRequest) error {
	var errors []string
	if len(req.Username) <= 0 {
		errors = append(errors, "Username is required")
	}

	if len(req.Email) <= 0 {
		errors = append(errors, "Email is required")
	}

	if len(req.Password) <= 0 {
		errors = append(errors, "Password is required")
	}

	if len(errors) > 0 {
		return fmt.Errorf("invalid Registration request: %s", errors)
	} else {
		return nil
	}
}

func validateUserLoginRequest(req UserLoginRequest) error {
	var errors []string
	if len(req.Email) <= 0 {
		errors = append(errors, "Email is required")
	}

	if len(req.Password) <= 0 {
		errors = append(errors, "Password is required")
	}

	if len(errors) > 0 {
		return fmt.Errorf("invalid login request: %s", errors)
	} else {
		return nil
	}
}
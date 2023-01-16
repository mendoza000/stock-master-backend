package helpers

import "github.com/mendoza000/stockmaster/models"

func ValidateBodyProduct(body models.Product) (message string) {
	if body.Title == "" {
		return ("Title is required")
	}
	if body.Brand == "" {
		return ("Brand is required")
	}
	if !body.Stock {
		return ("The stock to register a product must be true")
	}
	if body.Price == 0.0 {
		return ("Price is required")
	}
	if body.Details == "" {
		return ("Details is required")
	}
	if body.Amount == 0 {
		return ("Amount is required")
	}

	return ""
}

func ValidateBodyUser(body models.User) (message string) {
	if body.Username == "" {
		return "Username is required"
	}
	if body.Mail == "" {
		return "Mail is required"
	}
	if body.Password == "" {
		return "Password is required"
	}
	if body.UserRange == 0 {
		return "UserRange is required"
	}
	return ""
}

func ValidateBodyUserLogin(body models.UserLogin) (message string)  {
	if body.Username == "" {
		return "Username is required"
	}
	if body.Password == "" {
		return "Password is required"
	}
	return ""
}
package main

import (
	"github.com/go-playground/validator"
	"github.com/gumanzanoo/email-service-provider/internal/domain/campaign"
)

func main() {
	contacts := []campaign.Contact{{Email: "fusca"}, {Email: ""}}
	campaign := campaign.Campaign{Contacts: contacts}
	validate := validator.New()
	err := validate.Struct(campaign)

	if err == nil {
		println("nenhum erro")
	} else {
		validationErrors := err.(validator.ValidationErrors)
		for _, v := range validationErrors {
			println(v.StructField() + "is invalid: " + v.Tag())
		}
	}
}

package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/gumanzanoo/email-service-provider/internal/domain/campaign"
	"github.com/gumanzanoo/email-service-provider/internal/infraestructure/database"
	"github.com/gumanzanoo/email-service-provider/internal/routes"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	campaignService := campaign.Service{
		Repository: &database.CampaignRepository{},
	}

	handler := routes.Handler{
		CampaignService: campaignService,
	}

	r.Post("/campaign", routes.ErrorHandler(handler.CampaignPost))
	r.Get("/campaign", routes.ErrorHandler(handler.CampaignGet))

	err := http.ListenAndServe(":3001", r)
	if err != nil {
		return
	}
}

package routes

import "github.com/gumanzanoo/email-service-provider/internal/domain/campaign"

type Handler struct {
	CampaignService campaign.Service
}

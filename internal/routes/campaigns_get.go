package routes

import (
	"net/http"
)

func (h *Handler) CampaignGet(_ http.ResponseWriter, _ *http.Request) (interface{}, int, error) {
	campaigns, err := h.CampaignService.Repository.Get()
	return campaigns, 200, err

}

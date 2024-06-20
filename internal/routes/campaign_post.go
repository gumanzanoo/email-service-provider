package routes

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/gumanzanoo/email-service-provider/internal/contract"
)

func (h *Handler) CampaignPost(_ http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	var request contract.NewCampaign
	_ = render.DecodeJSON(r.Body, &request)
	id, err := h.CampaignService.Create(request)
	return map[string]string{"id": id}, 200, err
}

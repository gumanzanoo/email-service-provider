package campaign

import (
	"github.com/gumanzanoo/email-service-provider/internal/contract"
	"github.com/gumanzanoo/email-service-provider/internal/exceptions"
)

type Service struct {
	Repository Repository
}

func (s *Service) Create(newCampaign contract.NewCampaign) (string, error) {
	campaign, err := NewCampaign(newCampaign.Name, newCampaign.Content, newCampaign.Emails)
	if err != nil {
		return "", err
	}

	err = s.Repository.Save(campaign)
	if err != nil {
		return "", exceptions.InternalErr
	}

	return campaign.Id, nil
}

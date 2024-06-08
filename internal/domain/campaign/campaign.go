package campaign

import (
	"time"

	"github.com/gumanzanoo/email-service-provider/internal/exceptions"
	"github.com/rs/xid"
)

type Contact struct {
	Email string `validate:"email"`
}

type Campaign struct {
	Id        string    `validate:"required"`
	Name      string    `validate:"min=5,max=24"`
	Content   string    `validate:"required,min=5,max=1024"`
	Contacts  []Contact `validate:"min=1,dive"`
	CreatedAt time.Time `validate:"required"`
	UpdateAt  time.Time `validate:"required"`
}

func NewCampaign(name string, content string, emails []string) (*Campaign, error) {
	contacts := make([]Contact, len(emails))
	for i, email := range emails {
		contacts[i].Email = email
	}

	campaign := &Campaign{
		Id:        xid.New().String(),
		Name:      name,
		Content:   content,
		Contacts:  contacts,
		CreatedAt: time.Now(),
		UpdateAt:  time.Now(),
	}

	err := exceptions.ValidateStruct(campaign)
	if err == nil {
		return campaign, nil
	}
	return nil, err
}

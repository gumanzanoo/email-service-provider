package campaign

import (
	"testing"
	"time"

	"github.com/jaswdr/faker/v2"
	"github.com/stretchr/testify/assert"
)

var (
	name     = "Campanha X"
	content  = "body body body"
	contacts = []string{"email@e.com", "email2@e.com"}
	fake     = faker.New()
)

func Test_NewCampaign_CreateCampaign(t *testing.T) {
	asrt := assert.New(t)
	campaign, _ := NewCampaign(name, content, contacts)
	asrt.Equal(campaign.Name, name)
	asrt.Equal(campaign.Content, content)
	asrt.Equal(len(campaign.Contacts), len(contacts))
}

func Test_NewCampaign_IDIsNotNill(t *testing.T) {
	asrt := assert.New(t)
	campaign, _ := NewCampaign(name, content, contacts)
	asrt.NotNil(campaign.Id)
}

// func Test_NewCampaign_MustStatusStartWithPending(t *testing.T) {
// 	assert := assert.New(t)
// 	campaign, _ := NewCampaign(name, content, contacts)
// 	assert.Equal(Pending, campaign.Status)
// }

func Test_NewCampaign_CreatedOnMustBeNow(t *testing.T) {
	asrt := assert.New(t)
	now := time.Now().Add(-time.Minute)
	campaign, _ := NewCampaign(name, content, contacts)
	asrt.Greater(campaign.CreatedAt, now)
}

func Test_NewCampaign_MustValidateNameMin(t *testing.T) {
	asrt := assert.New(t)
	_, err := NewCampaign("", content, contacts)
	asrt.Equal("name must have 5 characters at least", err.Error())
}

func Test_NewCampaign_MustValidateNameMax(t *testing.T) {
	asrt := assert.New(t)
	_, err := NewCampaign(fake.Lorem().Text(30), content, contacts)
	asrt.Equal("name reached max of 24 characters", err.Error())
}

func Test_NewCampaign_MustValidateContentMin(t *testing.T) {
	asrt := assert.New(t)
	_, err := NewCampaign(name, "", contacts)
	asrt.Equal("content is required", err.Error())
}

func Test_NewCampaign_MustValidateContentMax(t *testing.T) {
	asrt := assert.New(t)
	_, err := NewCampaign(name, fake.Lorem().Text(1040), contacts)
	asrt.Equal("content reached max of 1024 characters", err.Error())
}

func Test_NewCampaign_MustValidateContactsMin(t *testing.T) {
	asrt := assert.New(t)
	_, err := NewCampaign(name, content, nil)
	asrt.Equal("contacts must have 1 characters at least", err.Error())
}

func Test_NewCampaign_MustValidateContacts(t *testing.T) {
	asrt := assert.New(t)
	_, err := NewCampaign(name, content, []string{"email_invalid"})
	asrt.Equal("email is not valid", err.Error())
}

// func Test_NewCampaign_MustValidateCreatedBy(t *testing.T) {
// 	asrt := assert.New(t)
// 	_, err := NewCampaign(name, content, contacts)
// 	asrt.Equal("createdby is invalid", err.Error())
// }

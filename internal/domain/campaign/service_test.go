package campaign

import (
	"errors"
	"testing"

	"github.com/gumanzanoo/email-service-provider/internal/contract"
	"github.com/gumanzanoo/email-service-provider/internal/exceptions"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type repositoryMock struct {
	mock.Mock
}

func (r *repositoryMock) Save(campaign *Campaign) error {
	args := r.Called(campaign)
	return args.Error(0)
}

func (r *repositoryMock) Get() []Campaign {
	//args := r.Called(campaign)
	return nil
}

var (
	newCampaign = contract.NewCampaign{
		Name:    "Teste Y",
		Content: "Body Body",
		Emails:  []string{"teste@e.com"},
	}
	service = Service{}
)

func Test_Create_Campaign(t *testing.T) {
	asrt := assert.New(t)
	repositoryMock := new(repositoryMock)
	repositoryMock.On("Save", mock.Anything).Return(nil)
	service.Repository = repositoryMock
	id, err := service.Create(newCampaign)
	asrt.NotNil(id)
	asrt.Nil(err)
	repositoryMock.AssertExpectations(t)
}

func Test_Save_Campaign(t *testing.T) {
	asrt := assert.New(t)
	repositoryMock := new(repositoryMock)
	repositoryMock.On("Save", mock.Anything).Return(errors.New("error while save into database"))
	service.Repository = repositoryMock
	_, err := service.Create(newCampaign)
	asrt.True(errors.Is(exceptions.InternalErr, err))
	repositoryMock.AssertExpectations(t)
}

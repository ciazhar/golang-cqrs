package postgres

import (
	"errors"
	"github.com/ciazhar/golang-cqrs/common"
	"github.com/ciazhar/golang-cqrs/common/env"
	"github.com/ciazhar/golang-cqrs/common/validator"
	"github.com/ciazhar/golang-cqrs/write-service-postgres/app"
	"github.com/ciazhar/golang-cqrs/write-service-postgres/internal/mocks"
	"github.com/ciazhar/golang-cqrs/write-service-postgres/pkg/social/model"
	"github.com/ciazhar/golang-cqrs/write-service-postgres/pkg/social/repository/postgres"
	"github.com/stretchr/testify/assert"
	"testing"
)

var Application *app.Application

func init() {
	application, err := app.SetupAppWithPath(env.GetEnvPath() + "/write-service-postgres/config.json")
	if err != nil {
		panic(err)
	}
	Application = application
}

func NewSocial() model.Social {
	var social model.Social
	common.ToStruct("write-service-postgres/testdata/social/actual.1.golden", &social)
	repo := postgres.NewSocialPostgresRepository(Application)
	_ = repo.Store(&social)
	return social
}

type SocialExample struct {
	SocialId string `valid:"socialMustExist"`
}

func NewPostExample() SocialExample {
	return SocialExample{SocialId: NewSocial().Id}
}

func TestSocialPostgresValidatorInit(t *testing.T) {

	t.Run("default", func(t *testing.T) {
		dummy := NewPostExample()
		repo := new(mocks.SocialPostgresRepository)
		repo.On("GetByID", dummy.SocialId).Return(NewSocial(), nil)
		NewSocialPostgresValidator(repo)

		err := validator.Struct(dummy)
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		dummy := NewPostExample()
		repo := new(mocks.SocialPostgresRepository)
		repo.On("GetByID", dummy.SocialId).Return(NewSocial(), errors.New("not found"))
		NewSocialPostgresValidator(repo)

		err := validator.Struct(dummy)
		assert.Error(t, err)
	})

}

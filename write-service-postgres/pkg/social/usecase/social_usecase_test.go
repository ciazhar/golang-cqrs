package usecase

import (
	"bou.ke/monkey"
	"errors"
	"github.com/ciazhar/golang-cqrs/common"
	"github.com/ciazhar/golang-cqrs/common/env"
	"github.com/ciazhar/golang-cqrs/write-service-postgres/app"
	"github.com/ciazhar/golang-cqrs/write-service-postgres/internal/mocks"
	"github.com/ciazhar/golang-cqrs/write-service-postgres/pkg/social/model"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var Application *app.Application

func init() {
	application, err := app.SetupAppWithPath(env.GetEnvPath() + "/write-service-postgres/config.json")
	if err != nil {
		panic(err)
	}
	Application = application
}

func NewActual() model.Social {
	var social model.Social
	common.ToStruct("social/actual.1.golden", &social)
	return social
}

func NewActual2() model.Social {
	var social model.Social
	common.ToStruct("social/actual.2.golden", &social)
	return social
}

func TestSocialUseCase_Store(t *testing.T) {
	repo := new(mocks.SocialPostgresRepository)
	uc := NewSocialUseCase(Application, repo)
	testCases := []struct {
		name        string
		social      model.Social
		returnError error
	}{
		{"default", NewActual(), nil},
		{"default2", NewActual2(), nil},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			repo.On("Store", &testCase.social).Return(testCase.returnError)

			err := uc.Store(&testCase.social)

			assert.NoError(t, err)
			repo.AssertExpectations(t)
		})
	}
}

func TestSocialUseCase_Update(t *testing.T) {
	repo := new(mocks.SocialPostgresRepository)
	uc := NewSocialUseCase(Application, repo)
	actual := NewActual()
	actual.Id = "100"
	testCases := []struct {
		name        string
		social      model.Social
		returnError error
	}{
		{"default", NewActual(), nil},
		{"default2", actual, errors.New("not found")},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			repo.On("GetByID", testCase.social.Id).Return(testCase.social, testCase.returnError)
			repo.On("Update", &testCase.social).Return(testCase.returnError)

			err := uc.Update(&testCase.social)

			assert.Equal(t, err, testCase.returnError)
			repo.AssertExpectations(t)
		})
	}
}

func TestSocialUseCase_Delete(t *testing.T) {
	repo := new(mocks.SocialPostgresRepository)
	uc := NewSocialUseCase(Application, repo)
	actual := NewActual()
	actual.Id = "100"
	testCases := []struct {
		name        string
		social      model.Social
		returnError error
	}{
		{"default", NewActual(), nil},
		{"error", actual, errors.New("not found")},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			wayback := time.Date(1974, time.May, 19, 1, 2, 3, 4, time.UTC)
			patch := monkey.Patch(time.Now, func() time.Time { return wayback })
			defer patch.Unpatch()

			repo.On("GetByID", testCase.social.Id).Return(testCase.social, testCase.returnError)

			testCase.social.DeletedAt = time.Now()
			repo.On("Update", &testCase.social).Return(testCase.returnError)

			err := uc.Delete(testCase.social.Id)

			assert.Equal(t, err, testCase.returnError)
			//repo.AssertExpectations(t)
		})
	}
}

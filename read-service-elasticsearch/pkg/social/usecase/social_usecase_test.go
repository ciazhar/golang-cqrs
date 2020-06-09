package usecase

import (
	"bou.ke/monkey"
	"errors"
	"github.com/ciazhar/golang-cqrs/common"
	"github.com/ciazhar/golang-cqrs/common/rest"
	"github.com/ciazhar/golang-cqrs/read-service-elasticsearch/internal/mocks"
	"github.com/ciazhar/golang-cqrs/read-service-elasticsearch/pkg/social/model"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func NewActual() model.Social {
	var social model.Social
	common.ToStruct("write-service-postgres/testdata/social/actual.1.golden", &social)
	return social
}

func NewActual2() model.Social {
	var social model.Social
	common.ToStruct("write-service-postgres/testdata/social/actual.2.golden", &social)
	return social
}

func TestSocialUseCase_Store(t *testing.T) {
	repo := new(mocks.SocialPostgresRepository)
	uc := NewSocialUseCase(repo)
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

func TestSocialUseCase_Fetch(t *testing.T) {
	repo := new(mocks.SocialPostgresRepository)
	uc := NewSocialUseCase(repo)
	testCases := []struct {
		name         string
		offset       int
		limit        int
		returnSocial []model.Social
		returnError  error
	}{
		{"default", 0, 10, []model.Social{NewActual(), NewActual2()}, nil},
		{"default2", 0, 5, []model.Social{NewActual(), NewActual2()}, nil},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			param := rest.NewParam()
			param.Offset = 1
			param.Limit = 10

			repo.On("Fetch", param).Return(testCase.returnSocial, testCase.returnError)

			expected, err := uc.Fetch(param)

			assert.NotEmpty(t, expected)
			assert.NoError(t, err)
			assert.Len(t, expected, len(testCase.returnSocial))
			repo.AssertExpectations(t)
		})
	}
}

func TestSocialUseCase_GetByID(t *testing.T) {
	repo := new(mocks.SocialPostgresRepository)
	uc := NewSocialUseCase(repo)
	testCases := []struct {
		name         string
		id           string
		returnSocial model.Social
		returnError  error
	}{
		{"default", "1", NewActual(), nil},
		{"default2", "2", NewActual(), nil},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			repo.On("GetByID", testCase.id).Return(testCase.returnSocial, testCase.returnError)

			expected, err := uc.GetByID(testCase.id)

			assert.NoError(t, err)
			assert.NotNil(t, expected)
			repo.AssertExpectations(t)
		})
	}
}

func TestSocialUseCase_Update(t *testing.T) {
	repo := new(mocks.SocialPostgresRepository)
	uc := NewSocialUseCase(repo)
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
	uc := NewSocialUseCase(repo)
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

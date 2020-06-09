package rest

import (
	"errors"
	"github.com/ciazhar/golang-cqrs/common"
	"github.com/ciazhar/golang-cqrs/write-service-postgres/internal/mocks"
	"github.com/ciazhar/golang-cqrs/write-service-postgres/pkg/social/model"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
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

func NewActualReader(payload model.Social) *strings.Reader {
	return common.ToReader(payload)
}

func TestSocialController_Store(t *testing.T) {
	testCases := []struct {
		name        string
		payload     model.Social
		reader      io.Reader
		returnError error
		httpStatus  int
	}{
		{"default", NewActual(), NewActualReader(NewActual()), nil, http.StatusOK},
		{"bad-request", NewActual(), strings.NewReader("{"), errors.New(""), http.StatusBadRequest},
		{"error", NewActual2(), NewActualReader(NewActual2()), errors.New(""), http.StatusInternalServerError},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			uc := new(mocks.SocialUseCase)
			ctrl := NewSocialController(uc)
			w := httptest.NewRecorder()
			router := gin.New()
			router.POST("/social", ctrl.Store)

			r, err := http.NewRequest(http.MethodPost, "/social", testCase.reader)
			assert.NoError(t, err)
			uc.On("Store", &testCase.payload).Return(testCase.returnError)

			router.ServeHTTP(w, r)

			assert.Equal(t, testCase.httpStatus, w.Code)
		})
	}
}

func TestSocialController_Update(t *testing.T) {
	testCases := []struct {
		name        string
		payload     model.Social
		reader      io.Reader
		returnError error
		httpStatus  int
	}{
		{"default", NewActual(), NewActualReader(NewActual()), nil, http.StatusOK},
		{"bad-request", NewActual(), strings.NewReader("{"), errors.New(""), http.StatusBadRequest},
		{"error", NewActual2(), NewActualReader(NewActual2()), errors.New(""), http.StatusInternalServerError},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			uc := new(mocks.SocialUseCase)
			ctrl := NewSocialController(uc)
			w := httptest.NewRecorder()
			router := gin.New()
			router.PUT("/social", ctrl.Update)

			r, err := http.NewRequest(http.MethodPut, "/social", testCase.reader)
			assert.NoError(t, err)
			uc.On("Update", &testCase.payload).Return(testCase.returnError)

			router.ServeHTTP(w, r)

			assert.Equal(t, testCase.httpStatus, w.Code)
		})
	}
}

func TestSocialController_Delete(t *testing.T) {
	testCases := []struct {
		name        string
		id          string
		returnError error
		httpStatus  int
	}{
		{"default", "1", nil, http.StatusOK},
		{"error", "10", errors.New(""), http.StatusInternalServerError},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			uc := new(mocks.SocialUseCase)
			ctrl := NewSocialController(uc)
			w := httptest.NewRecorder()

			router := gin.New()
			router.DELETE("/social/:id", ctrl.Delete)

			r, err := http.NewRequest(http.MethodDelete, "/social/"+testCase.id, nil)
			assert.NoError(t, err)
			uc.On("Delete", testCase.id).Return(testCase.returnError)

			router.ServeHTTP(w, r)

			assert.Equal(t, testCase.httpStatus, w.Code)
			uc.AssertExpectations(t)
		})
	}
}

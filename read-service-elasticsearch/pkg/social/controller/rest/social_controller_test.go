package rest

import (
	"errors"
	"github.com/ciazhar/golang-cqrs/common"
	"github.com/ciazhar/golang-cqrs/common/rest"
	"github.com/ciazhar/golang-cqrs/read-service-elasticsearch/internal/mocks"
	"github.com/ciazhar/golang-cqrs/read-service-elasticsearch/pkg/social/model"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strconv"
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

func TestSocialController_Fetch(t *testing.T) {
	testCases := []struct {
		name         string
		offset       int
		limit        int
		returnSocial []model.Social
		returnError  error
		httpStatus   int
	}{
		{"default", 0, 10, []model.Social{NewActual(), NewActual2()}, nil, http.StatusOK},
		{"error", 1, 5, []model.Social{NewActual(), NewActual2()}, errors.New(""), http.StatusInternalServerError},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			uc := new(mocks.SocialUseCase)
			ctrl := NewSocialController(uc)

			w := httptest.NewRecorder()

			router := gin.New()
			router.GET("/social", ctrl.Fetch)

			r, err := http.NewRequest(http.MethodGet, "/social?offset="+strconv.Itoa(testCase.offset)+"&limit="+strconv.Itoa(testCase.limit), nil)
			assert.NoError(t, err)
			param := rest.NewParam()
			param.Limit = testCase.limit
			param.Offset = testCase.offset
			uc.On("Fetch", param).Return(testCase.returnSocial, testCase.returnError)

			router.ServeHTTP(w, r)

			assert.Equal(t, testCase.httpStatus, w.Code)
			uc.AssertExpectations(t)
		})
	}

}

func TestSocialController_GetByID(t *testing.T) {
	testCases := []struct {
		name         string
		id           string
		returnSocial model.Social
		returnError  error
		httpStatus   int
	}{
		{"default", "1", NewActual(), nil, http.StatusOK},
		//{"bad-request","bukan-id",model.Social{}, errors.New(""), http.StatusBadRequest},
		{"internal-server-error", "10", model.Social{}, errors.New(""), http.StatusInternalServerError},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			uc := new(mocks.SocialUseCase)
			ctrl := NewSocialController(uc)
			w := httptest.NewRecorder()

			router := gin.New()
			router.GET("/social/:id", ctrl.GetByID)
			r, err := http.NewRequest(http.MethodGet, "/social/"+testCase.id, nil)
			assert.NoError(t, err)
			uc.On("GetByID", testCase.id).Return(testCase.returnSocial, testCase.returnError)

			router.ServeHTTP(w, r)

			assert.Equal(t, testCase.httpStatus, w.Code)
			uc.AssertExpectations(t)
		})
	}
}

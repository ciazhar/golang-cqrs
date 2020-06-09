package rest

import (
	"github.com/ciazhar/golang-cqrs/common/logger"
	"github.com/ciazhar/golang-cqrs/common/rest"
	"github.com/ciazhar/golang-cqrs/read-service-elasticsearch/pkg/social/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

type SocialController interface {
	Fetch(c *gin.Context)
	GetByID(c *gin.Context)
}

type socialController struct {
	SocialUseCase usecase.SocialUseCase
}

func (it socialController) Fetch(c *gin.Context) {
	param := rest.NewParamGin(c)
	param.Paging()

	payload, err := it.SocialUseCase.Fetch(param.Param)
	if err != nil {
		c.JSON(http.StatusInternalServerError, logger.ErrorS(err))
		return
	}
	c.JSON(http.StatusOK, payload)
}

func (it socialController) GetByID(c *gin.Context) {
	id := rest.RequestPathVariableString(c, "id")

	payload, err := it.SocialUseCase.GetByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, logger.ErrorS(err))
		return
	}
	c.JSON(http.StatusOK, payload)
}

func NewSocialController(SocialUseCase usecase.SocialUseCase) SocialController {
	return socialController{SocialUseCase: SocialUseCase}
}

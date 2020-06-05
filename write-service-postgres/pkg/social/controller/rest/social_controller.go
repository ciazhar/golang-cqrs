package rest

import (
	"github.com/ciazhar/golang-cqrs/common/logger"
	"github.com/ciazhar/golang-cqrs/common/rest"
	"github.com/ciazhar/golang-cqrs/write-service-postgres/pkg/social/model"
	"github.com/ciazhar/golang-cqrs/write-service-postgres/pkg/social/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

type SocialController interface {
	Store(c *gin.Context)
	Delete(c *gin.Context)
	Update(c *gin.Context)
}

type socialController struct {
	SocialUseCase usecase.SocialUseCase
}

func (it socialController) Store(ctx *gin.Context) {
	var payload model.Social
	if err := ctx.BindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, logger.WarnS(err))
		return
	}

	err := it.SocialUseCase.Store(&payload)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, logger.ErrorS(err))
		return
	}

	ctx.JSON(http.StatusOK, payload)
}

func (it socialController) Update(ctx *gin.Context) {
	var payload model.Social
	if err := ctx.BindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, logger.WarnS(err))
		return
	}

	err := it.SocialUseCase.Update(&payload)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, logger.ErrorS(err))
		return
	}

	ctx.JSON(http.StatusOK, payload)
}

func (it socialController) Delete(c *gin.Context) {
	id := rest.RequestPathVariableString(c, "id")

	if err := it.SocialUseCase.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, logger.ErrorS(err))
		return
	}
	c.JSON(http.StatusOK, nil)
}

func NewSocialController(SocialUseCase usecase.SocialUseCase) SocialController {
	return socialController{SocialUseCase: SocialUseCase}
}

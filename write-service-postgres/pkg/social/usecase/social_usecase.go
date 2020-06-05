package usecase

import (
	"errors"
	"github.com/ciazhar/golang-cqrs/common/logger"
	"github.com/ciazhar/golang-cqrs/common/validator"
	"github.com/ciazhar/golang-cqrs/write-service-postgres/app"
	"github.com/ciazhar/golang-cqrs/write-service-postgres/pkg/social/model"
	"github.com/ciazhar/golang-cqrs/write-service-postgres/pkg/social/repository/postgres"
	"github.com/imdario/mergo"
	"time"
)

type SocialUseCase interface {
	Store(req *model.Social) error
	Update(req *model.Social) error
	Delete(id string) error
}

type socialUseCase struct {
	Application      *app.Application
	SocialRepository postgres.SocialPostgresRepository
}

func (c socialUseCase) Update(req *model.Social) error {
	oldReq, err := c.SocialRepository.GetByID(req.Id)
	if err != nil {
		return logger.WithError(err)
	}

	if err := mergo.Merge(req, oldReq); err != nil {
		return logger.WithError(err)
	}
	if err := validator.Struct(req); err != nil {
		return logger.WithError(err)
	}

	req.CreatedAt = oldReq.CreatedAt
	req.UpdatedAt = time.Now()
	req.DeletedAt = oldReq.DeletedAt

	return c.SocialRepository.Update(req)
}

func (c socialUseCase) Delete(id string) error {
	payload, err := c.SocialRepository.GetByID(id)
	if err != nil {
		return logger.WithError(err)
	}
	if !payload.DeletedAt.IsZero() {
		return logger.WithError(errors.New("not found"))
	}
	payload.DeletedAt = time.Now()
	return c.SocialRepository.Update(&payload)
}

func (c socialUseCase) Store(req *model.Social) error {
	if err := validator.Struct(req); err != nil {
		return logger.WithError(err)
	}
	req.CreatedAt = time.Now()
	req.UpdatedAt = time.Now()

	publisher, err := c.Application.RabbitMQBroker.CreatePublisher("testing")
	if err != nil {
		return logger.WithError(err)
	}
	if err := publisher.Publish(req, req.Id); err != nil {
		return logger.WithError(err)
	}

	return c.SocialRepository.Store(req)
}

func NewSocialUseCase(application *app.Application, socialRepository postgres.SocialPostgresRepository) SocialUseCase {
	return socialUseCase{
		Application:      application,
		SocialRepository: socialRepository,
	}
}

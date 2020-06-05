package usecase

import (
	"errors"
	"github.com/ciazhar/golang-cqrs/common"
	"github.com/ciazhar/golang-cqrs/common/logger"
	"github.com/ciazhar/golang-cqrs/common/validator"
	"github.com/ciazhar/golang-cqrs/write-service-postgres/app"
	"github.com/ciazhar/golang-cqrs/write-service-postgres/pkg/social/model"
	"github.com/ciazhar/golang-cqrs/write-service-postgres/pkg/social/repository/postgres"
	"github.com/imdario/mergo"
	uuid "github.com/satori/go.uuid"
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
	StorePublisher   common.Publisher
	UpdatePublisher  common.Publisher
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

	if err := c.SocialRepository.Update(req); err != nil {
		return logger.WithError(err)
	}

	if err := c.UpdatePublisher.Publish(req, uuid.NewV1().String()); err != nil {
		return logger.WithError(err)
	}

	return nil
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

	if err := c.SocialRepository.Update(&payload); err != nil {
		return logger.WithError(err)
	}

	if err := c.UpdatePublisher.Publish(payload, uuid.NewV1().String()); err != nil {
		return logger.WithError(err)
	}

	return nil
}

func (c socialUseCase) Store(req *model.Social) error {
	if err := validator.Struct(req); err != nil {
		return logger.WithError(err)
	}
	req.CreatedAt = time.Now()
	req.UpdatedAt = time.Now()

	if err := c.SocialRepository.Store(req); err != nil {
		return logger.WithError(err)
	}

	if err := c.StorePublisher.Publish(req, uuid.NewV1().String()); err != nil {
		return logger.WithError(err)
	}

	return nil
}

func NewSocialUseCase(application *app.Application, socialRepository postgres.SocialPostgresRepository) (SocialUseCase, error) {
	r := socialUseCase{}

	storePublisher, err := application.RabbitMQBroker.CreatePublisher("social_store")
	if err != nil {
		return r, logger.WithError(err)
	}

	updatePublisher, err := application.RabbitMQBroker.CreatePublisher("social_update")
	if err != nil {
		return r, logger.WithError(err)
	}

	return socialUseCase{
		Application:      application,
		SocialRepository: socialRepository,
		StorePublisher:   storePublisher,
		UpdatePublisher:  updatePublisher,
	}, nil
}

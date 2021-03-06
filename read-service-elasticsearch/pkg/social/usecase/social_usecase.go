package usecase

import (
	"errors"
	"github.com/ciazhar/golang-cqrs/common/logger"
	"github.com/ciazhar/golang-cqrs/common/rest"
	"github.com/ciazhar/golang-cqrs/common/validator"
	"github.com/ciazhar/golang-cqrs/read-service-elasticsearch/pkg/social/model"
	"github.com/ciazhar/golang-cqrs/read-service-elasticsearch/pkg/social/repository/elastic"
	"github.com/imdario/mergo"
	"time"
)

type SocialUseCase interface {
	Fetch(param rest.Param) ([]model.Social, error)
	GetByID(id string) (model.Social, error)
	Store(req *model.Social) error
	Update(req *model.Social) error
	Delete(id string) error
}

type socialUseCase struct {
	SocialRepository elastic.SocialPostgresRepository
}

func (c socialUseCase) GetByID(id string) (model.Social, error) {
	return c.SocialRepository.GetByID(id)
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
	payload, err := c.GetByID(id)
	if err != nil {
		return logger.WithError(err)
	}
	if !payload.DeletedAt.IsZero() {
		return logger.WithError(errors.New("not found"))
	}
	payload.DeletedAt = time.Now()
	return c.SocialRepository.Update(&payload)
}

func (c socialUseCase) Fetch(param rest.Param) ([]model.Social, error) {
	return c.SocialRepository.Fetch(param)
}

func (c socialUseCase) Store(req *model.Social) error {
	if err := validator.Struct(req); err != nil {
		return logger.WithError(err)
	}
	req.CreatedAt = time.Now()
	req.UpdatedAt = time.Now()
	return c.SocialRepository.Store(req)
}

func NewSocialUseCase(SocialRepository elastic.SocialPostgresRepository) SocialUseCase {
	return socialUseCase{SocialRepository: SocialRepository}
}

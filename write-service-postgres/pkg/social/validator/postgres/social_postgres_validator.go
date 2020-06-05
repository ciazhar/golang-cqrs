package postgres

import (
	"github.com/asaskevich/govalidator"
	"github.com/ciazhar/golang-cqrs/write-service-postgres/pkg/social/repository/postgres"
)

type SocialPostgresValidator interface {
	SocialMustExist()
}

type socialPostgresValidator struct {
	SocialRepository postgres.SocialPostgresRepository
}

func (r socialPostgresValidator) SocialMustExist() {
	govalidator.TagMap["socialMustExist"] = func(str string) bool {
		return r.validatePostId(str)
	}
}

func (r socialPostgresValidator) validatePostId(postId string) bool {
	if postId != "" && govalidator.IsUUIDv4(postId) {
		if _, err := r.SocialRepository.GetByID(postId); err != nil {
			return false
		}
	}
	return true
}

func NewSocialPostgresValidator(SocialRepository postgres.SocialPostgresRepository) {
	validator := socialPostgresValidator{SocialRepository: SocialRepository}
	validator.SocialMustExist()
}

package postgres

import (
	"github.com/ciazhar/golang-cqrs/common/logger"
	"github.com/ciazhar/golang-cqrs/write-service-postgres/app"
	"github.com/ciazhar/golang-cqrs/write-service-postgres/pkg/social/model"
	"github.com/go-pg/pg/v9/orm"
	uuid "github.com/satori/go.uuid"
)

type SocialPostgresRepository interface {
	GetByID(id string) (model.Social, error)
	Store(req *model.Social) error
	Update(req *model.Social) error
}

type repository struct {
	app *app.Application
}

func (r repository) GetByID(id string) (model.Social, error) {
	social := model.Social{Id: id}
	if err := r.app.Postgres.Select(&social); err != nil {
		return social, logger.WithError(err)
	}
	return social, nil
}

func (r repository) Store(req *model.Social) error {
	id := uuid.Must(uuid.NewV4(), nil)
	req.Id = id.String()
	return r.app.Postgres.Insert(req)
}

func (r repository) Update(req *model.Social) error {
	return r.app.Postgres.Update(req)
}

func NewSocialPostgresRepository(app *app.Application) SocialPostgresRepository {
	r := repository{
		app: app,
	}
	if err := r.app.Postgres.CreateTable((*model.Social)(nil), &orm.CreateTableOptions{
		IfNotExists: true,
		Temp:        false,
	}); err != nil {
		panic(err)
	}

	return r
}

package pg

import (
	"context"
	"github.com/ciazhar/golang-cqrs/common/env"
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v9"
	"github.com/rs/zerolog/log"
)

type dbLogger struct{}

func (d dbLogger) BeforeQuery(c context.Context, _ *pg.QueryEvent) (context.Context, error) {
	return c, nil
}

func (d dbLogger) AfterQuery(_ context.Context, q *pg.QueryEvent) error {
	query, err := q.FormattedQuery()
	if err != nil {
		return err
	}
	log.Debug().Msg(query)
	return nil
}

func InitPG(environment *env.Environtment) *pg.DB {
	db := pg.Connect(&pg.Options{
		User:     environment.Get("postgres.username"),
		Password: environment.Get("postgres.password"),
		Database: environment.Get("postgres.database"),
		Addr:     environment.Get("postgres.host") + ":" + environment.Get("postgres.port"),
	})
	if gin.Mode() == gin.DebugMode {
		db.AddQueryHook(dbLogger{})
	}
	return db
}

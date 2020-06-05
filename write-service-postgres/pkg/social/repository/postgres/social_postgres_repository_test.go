package postgres

import (
	"github.com/ciazhar/golang-cqrs/common"
	"github.com/ciazhar/golang-cqrs/common/env"
	"github.com/ciazhar/golang-cqrs/write-service-postgres/app"
	"github.com/ciazhar/golang-cqrs/write-service-postgres/pkg/social/model"
	"github.com/go-pg/pg/v9/orm"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var Application *app.Application

func init() {
	application, err := app.SetupAppWithPath(env.GetEnvPath() + "/write-service-postgres/config.json")
	if err != nil {
		panic(err)
	}
	Application = application
_:
	application.Postgres.DropTable((*model.Social)(nil), &orm.DropTableOptions{
		IfExists: true,
		Cascade:  true,
	})
_:
	application.Postgres.CreateTable((*model.Social)(nil), nil)
}

var ID string

func NewActual() model.Social {
	var social model.Social
	common.ToStruct("social/actual.1.golden", &social)
	social.CreatedAt = time.Now()
	social.UpdatedAt = time.Now()
	return social
}

func NewActual2() model.Social {
	var social model.Social
	common.ToStruct("social/actual.2.golden", &social)
	social.CreatedAt = time.Now()
	social.UpdatedAt = time.Now()
	return social
}

func TestRepositoryStore(t *testing.T) {
	actual := NewActual()
	actual2 := NewActual2()
	repo := NewSocialPostgresRepository(Application)

	t.Run("default", func(t *testing.T) {
		err := repo.Store(&actual)
		assert.NoError(t, err)
		ID = actual.Id
	})
	t.Run("default2", func(t *testing.T) {
		err := repo.Store(&actual2)
		assert.NoError(t, err)
	})
}

func TestRepositoryGetByID(t *testing.T) {
	repo := NewSocialPostgresRepository(Application)

	t.Run("default", func(t *testing.T) {
		expected, err := repo.GetByID(ID)

		assert.NotNil(t, expected)
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		expected, err := repo.GetByID("100")

		assert.NotNil(t, expected)
		assert.Error(t, err)
	})
}

func TestRepositoryUpdate(t *testing.T) {
	actual := NewActual()
	repo := NewSocialPostgresRepository(Application)

	t.Run("default", func(t *testing.T) {
		actual.Id = ID
		err := repo.Update(&actual)
		assert.NoError(t, err)
	})
}

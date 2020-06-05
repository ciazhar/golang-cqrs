package model

import "time"

type Social struct {
	tableName struct{}  `pg:"social"`
	Id        string    `json:"id"  pg:",pk"`
	Name      string    `json:"name" pg:",unique"`
	Detail    string    `json:"detail"`
	CreatedAt time.Time `json:"created_at" `
	UpdatedAt time.Time `json:"updated_at" `
	DeletedAt time.Time `json:"deleted_at" `
}

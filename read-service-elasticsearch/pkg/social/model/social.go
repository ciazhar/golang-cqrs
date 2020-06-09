package model

import "time"

type Social struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Detail    string    `json:"detail"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

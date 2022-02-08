package model

type Record interface {
	Id() string
	SetId(id string)
}

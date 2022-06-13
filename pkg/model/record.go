package model

type Record interface {
	Id() string
	IdFieldName() string
}

type IdAutoGenerator interface {
	Record
	SetId(id string)
}

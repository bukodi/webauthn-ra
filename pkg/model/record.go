package model

type Record interface {
	Id() string
	SetId(id string)
	MarshalJSON() ([]byte, error)
	UnmarshalJSON(b []byte) error
}

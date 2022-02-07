package model

type Record interface {
	Id() string
	MarshalJSON() ([]byte, error)
	UnmarshalJSON(b []byte) error
}

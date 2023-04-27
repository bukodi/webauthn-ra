package cfggorm

import (
	"github.com/bukodi/webauthn-ra/pkg/internal/repo"
	"time"
)

type LatestConfig struct {
	AlwaysOne     int `gorm:"primary_key;column:always_one"`
	LatestStateId string
	Created       time.Time
}

func (l LatestConfig) Id() string {
	return "1"
}

func (l LatestConfig) IdFieldName() string {
	return "always_one"
}

var _ repo.Record = &LatestConfig{}

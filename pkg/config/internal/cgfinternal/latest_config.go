package cgfinternal

import (
	"github.com/bukodi/webauthn-ra/pkg/config"
	"time"
)

type LatestConfig struct {
	LatestStateId config.StateId
	Created       time.Time
}

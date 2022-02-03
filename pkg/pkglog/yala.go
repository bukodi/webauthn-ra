package pkglog

import (
	"context"
	"fmt"
	"github.com/elgopher/yala/logger"
)

// define global logger, no need to initialize it (by default nothing is logged)
var log logger.Global

func Logf(ctx context.Context, format string, a ...interface{}) {
	log.Info(ctx, fmt.Sprintf(format, a))
}

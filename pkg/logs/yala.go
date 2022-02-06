package logs

import (
	"context"
	"fmt"
	"github.com/elgopher/yala/logger"
)

var yalaLog logger.Global

func Debuf(ctx context.Context, format string, a ...interface{}) {
	yalaLog.Debug(ctx, fmt.Sprintf(format, a))
}

func Infof(ctx context.Context, format string, a ...interface{}) {
	yalaLog.Info(ctx, fmt.Sprintf(format, a))
}

func Errorf(ctx context.Context, format string, a ...interface{}) {
	yalaLog.Error(ctx, fmt.Sprintf(format, a))
}

func Error(ctx context.Context, err error) {
	yalaLog.WithError(err).Error(ctx, err.Error())
}

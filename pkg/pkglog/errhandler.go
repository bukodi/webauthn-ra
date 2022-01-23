package pkglog

import (
	"context"
	"log"
)

func Handle(ctx context.Context, err error) error {
	return err
}

func LogError(ctx context.Context, err error) {
	log.Println(err)
}

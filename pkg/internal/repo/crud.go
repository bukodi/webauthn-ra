package repo

import (
	"context"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"github.com/bukodi/webauthn-ra/pkg/errlog"
	"github.com/oklog/ulid/v2"
	"time"
)

func Create[R Record](ctx context.Context, r R) error {
	err := NewWriteTx(ctx, func(ctx context.Context) error {
		if r.Id() == "" {
			if iag, ok := any(r).(IdAutoGenerator); ok {
				id, err := ulid.New(uint64(time.Now().UnixMilli()), rand.Reader)
				if err != nil {
					return errlog.Handle(ctx, err)
				}
				iag.SetId(id.String())
			} else {
				return errlog.Handle(ctx, fmt.Errorf("id isn't set")) // TODO replace with Err constant
			}
		}

		jsonBytes, err := json.Marshal(r)
		errlog.Debugf(ctx, "Json: %s", string(jsonBytes))

		tx, err := RequiresWriteTx(ctx)
		if err != nil {
			return errlog.Handle(ctx, err)
		}

		tx.Create(r)
		if tx.Error != nil {
			return errlog.Handle(ctx, err)
		}

		return nil
	})
	return errlog.Handle(ctx, err)
}

func FindById[R Record](ctx context.Context, obj R, id string) error {
	tx := dbInstance.First(obj, "id = ?", id)
	if tx.Error != nil {
		return errlog.Handle(ctx, tx.Error)
	}
	return nil
}

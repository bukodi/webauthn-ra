package config

import (
	"context"
	"fmt"
	"github.com/bukodi/webauthn-ra/pkg/config/internal/cfginternal"
	"github.com/bukodi/webauthn-ra/pkg/errlog"
	"github.com/bukodi/webauthn-ra/pkg/internal/repo"
	"sync"
)

type Options struct {
}

func Init(ctx context.Context, opts *Options) error {
	if err := repo.RegisterType[*cfginternal.ChangeTx](); err != nil {
		return errlog.Handle(ctx, err)
	}
	if err := repo.RegisterType[*cfginternal.LatestConfig](); err != nil {
		return errlog.Handle(ctx, err)
	}
	return nil
}

var configTypes = sync.Map{}

func RegisterType[R repo.Record]() error {
	if err := repo.RegisterType[R](); err != nil {
		return errlog.Handle(context.TODO(), err)
	}

	var r R
	typeName := fmt.Sprintf("%T", r)

	if _, swapped := configTypes.Swap(typeName, typeName); swapped {
		return errlog.Handle(context.TODO(), fmt.Errorf("config type already registered: %s", typeName))
	} else {
		configTypes.Store(typeName, typeName)
	}

	return nil
}

func TypeNames() []string {
	var names []string
	configTypes.Range(func(key, value interface{}) bool {
		names = append(names, key.(string))
		return true
	})
	return names
}

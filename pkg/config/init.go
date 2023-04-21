package config

import (
	"context"
	"fmt"
	"github.com/bukodi/webauthn-ra/pkg/errlog"
	"github.com/bukodi/webauthn-ra/pkg/model"
	"github.com/bukodi/webauthn-ra/pkg/repo"
	"sync"
)

type Options struct {
}

func Init(ctx context.Context, opts *Options) error {
	if err := repo.RegisterType[*ChangeTx](); err != nil {
		return errlog.Handle(ctx, err)
	}
	return nil
}

var configTypes = sync.Map{}

func RegisterType[R model.Record]() error {
	if err := repo.RegisterType[R](); err != nil {
		return errlog.Handle(context.TODO(), err)
	}

	var r R
	typeName := fmt.Sprintf("%T", r)

	// TODO: this is not atomic, use CompareAndSwap
	if _, found := configTypes.Load(typeName); found {
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

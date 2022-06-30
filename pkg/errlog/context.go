package errlog

// See this article: https://blog.khanacademy.org/statically-typed-context-in-go/

import (
	"context"
	"reflect"
)

type ContextAware[T any] interface {
	Exec(fn func())
	Current(ctx context.Context) T
}

type UserCtx struct {
	name string
}

func FromContext[T any](ctx context.Context, pt *T) {
	key := reflect.TypeOf(*pt)
	x := ctx.Value(key)
	if x != nil {
		if t, ok := x.(T); ok {
			*pt = t
		}
	}
}

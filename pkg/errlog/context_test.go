package errlog

import (
	"context"
	"reflect"
	"testing"
)

func TestUserCtx(t *testing.T) {
	uctx := &UserCtx{
		name: "John Doe",
	}

	ctx := context.WithValue(context.Background(), reflect.TypeOf(uctx), uctx)
	var uctx2 UserCtx
	FromContext(ctx, &uctx2)
	t.Log(uctx2)
}

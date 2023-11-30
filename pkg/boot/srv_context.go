package boot

import (
	"context"
	"time"
)

type Runtime struct {
	parent     context.Context
	cancelFunc context.CancelCauseFunc
}

func (rt *Runtime) StartServe() error {
	rt.parent, rt.cancelFunc = context.WithCancelCause(rt.parent)
	return nil
}

func (rt *Runtime) StopServe() error {
	rt.cancelFunc(context.Canceled)
	return nil
}

func (rt *Runtime) Deadline() (deadline time.Time, ok bool) {
	return rt.parent.Deadline()
}

func (rt *Runtime) Done() <-chan struct{} {
	return rt.parent.Done()
}

func (rt *Runtime) Err() error {
	return rt.parent.Err()
}

func (rt *Runtime) Value(key any) any {
	return rt.parent.Value(key)
}

// Type guard
var _ context.Context = (*Runtime)(nil)

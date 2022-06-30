package errlog

import (
	"fmt"
)

func ToError(recovered interface{}) error {
	if recovered == nil {
		return nil
	}
	err, isErr := recovered.(error)
	if isErr {
		return err
	} else {
		return fmt.Errorf("%v", recovered)
	}
}

func CatchPanicToHandler(fn func(err error)) {
	recoveredErr := ToError(recover())
	if recoveredErr != nil {
		fn(recoveredErr)
	}
}

func CatchPanicToVar(destErr *error) {
	recoveredErr := ToError(recover())
	if recoveredErr != nil {
		*destErr = recoveredErr
	}
}

func Check1st(err error) {
	if err != nil {
		panic(err)
	}
}

func Check2nd[A any](a A, err error) A {
	if err != nil {
		panic(err)
	}
	return a
}

func Check3drd[A, B any](a A, b B, err error) (A, B) {
	if err != nil {
		panic(err)
	}
	return a, b
}

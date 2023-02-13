package errlog

import (
	"fmt"
	"strings"
	"testing"
)

func TestCheck2nd(t *testing.T) {

	tests := []struct {
		x, y, z int
		wantErr bool
	}{
		{2, 0, 0, true},
		{2, 2, 1, false},
		{3, 2, 0, true},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%ddiv%d", tt.x, tt.y), func(t *testing.T) {
			if z, err := fnSafeDiv(tt.x, tt.y); (err != nil) != tt.wantErr || z != tt.z {
				t.Errorf("fnSafeDiv() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}

}

func fnSafeDiv(x, y int) (z int, err error) {
	defer CatchPanicToVar(&err)
	z = Check2nd(fnUnsafeDiv(x, y))
	return z, nil
}

func fnUnsafeDiv(x, y int) (int, error) {
	if x%y != 0 {
		return 0, fmt.Errorf("remainder isn't 0")
	}
	z := x / y
	return z, nil
}

func TestCatchPanicToVar(t *testing.T) {
	var err error
	if err = fnResourceWrite("f"); err != nil {
		t.Log(err)
	}
	if err = fnResourceWrite("cant_open"); err != nil {
		t.Log(err)
	}
	if err = fnResourceWrite("cant_close"); err != nil {
		t.Log(err)
	}
	if err = fnResourceWrite("cant_close, cant_write"); err != nil {
		t.Log(err)
	}
	if err = fnResourceWrite("cant_write"); err != nil {
		t.Log(err)
	}
}

func fnResourceWrite(resourceName string) (err error) {
	defer CatchPanicToVar(&err)

	if strings.Contains(resourceName, "cant_open") {
		return fmt.Errorf("can't open: %s", resourceName)
	}
	defer func() {
		if strings.Contains(resourceName, "cant_close") {
			panic(fmt.Errorf("can't close: %s", resourceName))
		}
	}()

	if strings.Contains(resourceName, "cant_write") {
		return fmt.Errorf("can't write: %s", resourceName)
	}

	return nil
}

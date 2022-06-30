package errlog

import (
	"fmt"
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

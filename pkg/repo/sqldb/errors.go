package sqldb

import "fmt"

type ErrUnsupportedDriver struct {
	driver string
}

func (e *ErrUnsupportedDriver) Error() string { return fmt.Sprintf("unsupported driver: %s", e.driver) }

package preventsuspend

import (
	"errors"
	"fmt"
)

var (
	ErrSuspendFunctionUnset = errors.New("suspend func was not set")
)

type Observer struct {
	suspendFunc func(bool)
}

func (o *Observer) Observe() error {
	if o.suspendFunc == nil {
		return fmt.Errorf("%s", ErrSuspendFunctionUnset.Error())
	}

	sigs := make()

	for {

	}
}

func (o *Observer) SetSuspendStatus(func(bool)) {

}

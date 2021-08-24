package concurrent_map

import "fmt"

type IllegalParameterError struct {
	msg string
}

func newIllegalParameterError(errMsg string) IllegalParameterError {
	return IllegalParameterError{
		msg: fmt.Sprintf("concurrent map: illegal parameter: %s", errMsg),
	}
}

func (e IllegalParameterError) Error() string {
	return e.msg
}
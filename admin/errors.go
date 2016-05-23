package admin

import "fmt"

// returned if app with given name is already registered
type ErrAlreadyRegistered struct {
	ftype string
	name string
}

func (t *ErrAlreadyRegistered) Error() string {
	return fmt.Sprintf("%s with name '%s' is already registered", t.ftype, t.name)
}


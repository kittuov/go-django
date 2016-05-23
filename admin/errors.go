package admin

import "fmt"

// returned if app with given name is already registered
type AlreadyRegistered struct {
	ftype string
	name string
}

func (t AlreadyRegistered) Error() string {
	return fmt.Sprintf("%s with name '%s' is already registered", t.ftype, t.name)
}


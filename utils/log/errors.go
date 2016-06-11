package log

import "fmt"

type LevelDoesNotExist int

func (l LevelDoesNotExist) Error() string {
	return fmt.Sprintf("Log module doesnot contain level %2d ", l)
}

package godjadmin

import (
	"os"
	"fmt"
)

type Command struct {
	Cmd       string
	ShortDesc string
	LongDesc  string
	Func      func(...string)
}

var (
	commands = []Command{
		{
			Cmd:"Run",
			ShortDesc:"Hello Bro",
		},
	}
)

func main() {
	args := os.Args[1:]
	fmt.Print(args)
	fmt.Print(commands)
}

func Serve()
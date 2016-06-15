/*
Use this package's methods to write logs any where
*/
package log

import (
	"log"
	"os"
	"github.com/fatih/color"
)

const (
	VERB = 0
	DEBU = 1
	INFO = 2
	NOTI = 3
	WARN = 4
	ERRO = 5
	CRIT = 6
	default_level = INFO
)

var (
	log_min_level = 0
	levels = logLevels{
		&logLevel{
			Level:VERB,
			Prefix:"VERBOSE",
			ColorFunc:color.New(color.FgBlue).SprintFunc(),
		},
		&logLevel{
			Level:DEBU,
			Prefix:"DEBUG",
			ColorFunc:color.New(color.FgCyan).SprintFunc(),
		},
		&logLevel{
			Level:INFO,
			Prefix:"INFO",
			ColorFunc:color.New(color.FgWhite).SprintFunc(),
		},
		&logLevel{
			Level:NOTI,
			Prefix:"NOTICE",
			ColorFunc:color.New(color.FgGreen).SprintFunc(),
		},
		&logLevel{
			Level:WARN,
			Prefix:"WARNING",
			ColorFunc: color.New(color.FgYellow).SprintFunc(),
		},
		&logLevel{
			Level:ERRO,
			Prefix:"ERROR",
			ColorFunc: color.New(color.FgRed).SprintFunc(),
		},
		&logLevel{
			Level:CRIT,
			Prefix:"CRITICAL",
			ColorFunc: color.New(color.FgMagenta).SprintFunc(),
		},
	}
)

/*
Initialization methods. they are run when running the program
*/
func init() {
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Lshortfile + log.Ltime)
}


/*
Log takes in level of the log and data to log.
if the logging fails or unable to log,
it does not panic or return error.
Best practice is to use the levels set in this package
eg., log.INFO, log.DEBU, log.WARN etc.,
*/
func Log(level int, v ...interface{}) {
	baseLog(level, v)
}

func baseLog(level int, v ...interface{}) {
	// Don't log if the level is less than
	if level < log_min_level {
		return
	}

	// Get the log level data
	log_level := levels.getLevel(level)

	// Set the new prefix
	prefix := log.Prefix()
	coloured_prefix := log_level.ColorFunc(log_level.Prefix + " ► " + prefix)
	log.SetPrefix(coloured_prefix)

	// Log the data
	coloured_str := log_level.ColorFunc(v...)
	log.Output(3, coloured_str)

	// Reset the prefix
	log.SetPrefix(prefix)
}

/*
SetMin sets the minimum level to log into the system
*/
func SetMin(level int) error {
	if levels.getLevel(level) != nil {
		log_min_level = level
		return nil
	}
	return LevelDoesNotExist(level)
}
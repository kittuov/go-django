/*
Use this package's methods to write logs any where
*/
package log

import (
	"log"
	"os"
)

const (
	VERBOSE = 0
	DEBUG = 1
	INFO = 2
	WARN = 3
	ERROR = 4
	FATAL = 5
)

var (
	verbose_logger *log.Logger
	debug_logger *log.Logger
	info_logger *log.Logger
	warn_logger *log.Logger
	error_logger *log.Logger
	fatal_logger *log.Logger
)

func init() {
	verbose_logger = log.New(os.Stdout, "VERBOSE:", log.Ldate + log.Ltime)
	debug_logger = log.New(os.Stdout, "DEBUG:", log.Ldate + log.Ltime)
	info_logger = log.New(os.Stdout, "INFO:", log.Ldate + log.Ltime)
	warn_logger = log.New(os.Stdout, "WARN:", log.Ldate + log.Ltime)
	error_logger = log.New(os.Stdout, "ERROR:", log.Ldate + log.Ltime)
	fatal_logger = log.New(os.Stdout, "FATAL:", log.Ldate + log.Ltime)
}

// Prints debug level statements
func Verbose(v ...interface{}) {
	verbose_logger.Print(v...)
}

// same as Debug function with formatted string
func Verbosef(format string, v ...interface{}) {
	verbose_logger.Printf(format, v...)
}

// Prints debug level statements
func Debug(v ...interface{}) {
	debug_logger.Print(v...)
}

// same as Debug function with formatted string
func Debugf(format string, v ...interface{}) {
	debug_logger.Printf(format, v...)
}

// Prints info level statements
func Info(v ...interface{}) {
	info_logger.Print(v...)
}

// same as Info function with formatted string
func Infof(format string, v ...interface{}) {
	info_logger.Printf(format, v...)
}

// Prints warn level statements
func Warn(v ...interface{}) {
	warn_logger.Print(v...)
}

// same as Warn function with formatted string
func Warnf(format string, v ...interface{}) {
	warn_logger.Printf(format, v...)
}

// Prints error level statements
func Error(v ...interface{}) {
	error_logger.Print(v...)
}

// same as Error function with formatted string
func Errorf(format string, v ...interface{}) {
	error_logger.Printf(format, v...)
}

// Prints fatal level statements
func Fatal(v ...interface{}) {
	fatal_logger.Print(v...)
}

// same as Fatal function with formatted string
func Fatalf(format string, v ...interface{}) {
	fatal_logger.Printf(format, v...)
}


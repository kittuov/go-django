package log

import "fmt"

// Verbose Methods
func Verbose(v ...interface{}) {
	baseLog(VERB, v...)
}

func Verbosef(format string, v ...interface{}) {
	baseLog(VERB, fmt.Sprintf(format, v...))
}

// Debug Methods
func Debug(v ...interface{}) {
	baseLog(DEBU, v...)
}

func Debugf(format string, v ...interface{}) {
	baseLog(DEBU, fmt.Sprintf(format, v...))
}

// Info logs
func Info(v ...interface{}) {
	baseLog(INFO, v...)
}

func Infof(format string, v ...interface{}) {
	baseLog(INFO, fmt.Sprintf(format, v...))
}

// Notice logs
func Notice(v ...interface{}) {
	baseLog(NOTI, v...)
}

func Noticef(format string, v ...interface{}) {
	baseLog(NOTI, fmt.Sprintf(format, v...))
}

// Notice logs
func Warn(v ...interface{}) {
	baseLog(NOTI, v...)
}

func Warnf(format string, v ...interface{}) {
	baseLog(WARN, fmt.Sprintf(format, v...))
}

// Error logs
func Error(v ...interface{}) {
	baseLog(ERRO, v...)
}

func Errorf(format string, v ...interface{}) {
	baseLog(ERRO, fmt.Sprintf(format, v...))
}

// Critical logs
func Critical(v ...interface{}) {
	baseLog(CRIT, v...)
}

func Criticalf(format string, v ...interface{}) {
	baseLog(CRIT, fmt.Sprintf(format, v...))
}

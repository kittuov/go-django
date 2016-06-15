package log

import "fmt"

// Verbose Methods
func Verbose(v...interface{}) {
	Log(VERB, v...)
}

func Verbosef(format string, v...interface{}) {
	Verbose(fmt.Sprintf(format, v...))
}

// Debug Methods
func Debug(v...interface{}) {
	Log(DEBU, v...)
}

func Debugf(format string, v...interface{}) {
	Debug(fmt.Sprintf(format, v...))
}

// Info logs
func Info(v...interface{}) {
	Log(INFO, v...)
}

func Infof(format string, v...interface{}) {
	Info(fmt.Sprintf(format, v...))
}

// Notice logs
func Notice(v...interface{}) {
	Log(NOTI, v...)
}

func Noticef(format string, v...interface{}) {
	Notice(fmt.Sprintf(format, v...))
}

// Notice logs
func Warn(v...interface{}) {
	Log(NOTI, v...)
}

func Warnf(format string, v...interface{}) {
	Warn(fmt.Sprintf(format, v...))
}

// Error logs
func Error(v...interface{}) {
	Log(ERRO, v...)
}

func Errorf(format string, v...interface{}) {
	Error(fmt.Sprintf(format, v...))
}

// Critical logs
func Crit(v...interface{}) {
	Log(CRIT, v...)
}

func Critf(format string, v...interface{}) {
	Crit(fmt.Sprintf(format, v...))
}


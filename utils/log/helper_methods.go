package log

import "fmt"

// Verbose Methods
func Verb(v...interface{}) {
	Log(VERB, v...)
}

func Verbf(format string, v...interface{}) {
	Verb(fmt.Sprintf(format, v...))
}

// Debug Methods
func Debu(v...interface{}) {
	Log(DEBU, v...)
}

func Debuf(format string, v...interface{}) {
	Debu(fmt.Sprintf(format, v...))
}

// Info logs
func Info(v...interface{}) {
	Log(INFO, v...)
}

func Infof(format string, v...interface{}) {
	Info(fmt.Sprintf(format, v...))
}

// Notice logs
func Noti(v...interface{}) {
	Log(NOTI, v...)
}

func Notif(format string, v...interface{}) {
	Noti(fmt.Sprintf(format, v...))
}

// Error logs
func Erro(v...interface{}) {
	Log(ERRO, v...)
}

func Errof(format string, v...interface{}) {
	Erro(fmt.Sprintf(format, v...))
}

// Critical logs
func Crit(v...interface{}) {
	Log(CRIT, v...)
}

func Critf(format string, v...interface{}) {
	Crit(fmt.Sprintf(format, v...))
}


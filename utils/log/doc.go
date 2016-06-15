/*
Package log aims at providing level based logging while still supporting
all the code written in go's native log package.

In short

log.Log method lets you log the data with levels.
log.SetMinLevel lets you set minimum level to log.
There are 7 levels. There are also helper methods for you to run those logs.

	log.VERB
	log.DEBU
	log.INFO
	log.NOTI
	log.WARN
	log.ERRO
	log.CRIT

Compatibility

The package is also compatible with go's built in log pkg. it respects any prefix
you set, it will output to any output writer you set. Also, any code written using
the package works with this package seamlessly. it contains all the helper
funcs present in go's log pkg. By default it logs into INFO level

BUG(kittuov@gmail.com): adding file name flag just prints out log file number
line rather than the calling file line number.

~~BUG(kittuov@gmail.com): No color for flags in log eg: time, short_file, date etc.~~
*/
package log

package log

/*
	import (
		"log"
	}

	and 

	import (
		"go.wit.com/log"
	}

	Should work exactly the same.

	These are golang log functions that are not changed
	at all. The arguments are all just passed straight through
	so this package appears to work exactly like the original ones
*/

// TODO: fill in the other functions from "log". Is there a way to automagically do that?
// the full list is:
/*
	type Logger

		// NEED THESE
		func (l *Logger) Fatal(v ...any)
		func (l *Logger) Fatalf(format string, v ...any)
		func (l *Logger) Fatalln(v ...any)
		func (l *Logger) Panic(v ...any)
		func (l *Logger) Panicf(format string, v ...any)
		func (l *Logger) Panicln(v ...any)
		func (l *Logger) Print(v ...any)
		func (l *Logger) Printf(format string, v ...any)
		func (l *Logger) Println(v ...any)

		func Default() *Logger
		func New(out io.Writer, prefix string, flag int) *Logger

		// what are these?
		func (l *Logger) Flags() int
		func (l *Logger) SetFlags(flag int)
		func (l *Logger) Prefix() string
		func (l *Logger) SetPrefix(prefix string)

		// probably not this stuff
		func (l *Logger) SetOutput(w io.Writer)
		func (l *Logger) Output(calldepth int, s string) error
		func (l *Logger) Writer() io.Writer
*/


import 	(
	origlog "log"
)

func Println(a ...any) {
	if ! PRINTLN.Ok() { return }
	if ! PRINTLN.b { return }
	origlog.Println(a...)
}

func Printf(s string, a ...any) {
	if ! PRINTLN.Ok() { return }
	if ! PRINTLN.b { return }
	origlog.Printf(s, a...)
}

func Fatalf(s string, a ...any) {
	origlog.Fatalf(s, a...)
}

func Fatal(s string, a ...any) {
	origlog.Fatalf(s, a...)
}

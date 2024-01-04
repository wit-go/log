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

import 	(
	origlog "log"
)

func Println(a ...any) {
	if ! PRINTLN.B { return }
	origlog.Println(a...)
}

func Printf(s string, a ...any) {
	if ! PRINTLN.B { return }
	origlog.Printf(s, a...)
}

func Fatalf(s string, a ...any) {
	origlog.Fatalf(s, a...)
}

func Fatal(s string, a ...any) {
	origlog.Fatalf(s, a...)
}

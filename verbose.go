package log

import 	(
	golanglog "log"
)

func Verbose(a ...any) {
	if ! VERBOSE.B { return }
	golanglog.Println(a...)
}

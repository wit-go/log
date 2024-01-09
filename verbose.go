package log

import 	(
	golanglog "log"
)

func Verbose(a ...any) {
	if ! VERBOSE.Ok() { return }
	if ! VERBOSE.b { return }
	golanglog.Println(a...)
}

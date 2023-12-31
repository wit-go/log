package log

import 	(
	golanglog "log"
)

func Verbose(a ...any) {
	if ! VERBOSE { return }
	golanglog.Println(a...)
}

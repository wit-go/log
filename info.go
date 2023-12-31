package log

import 	(
	golanglog "log"
)

func Info(a ...any) {
	if ! INFO { return }
	golanglog.Println(a...)
}

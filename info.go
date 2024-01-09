package log

import 	(
	golanglog "log"
)

func Info(a ...any) {
	if ! INFO.Ok() { return }
	if ! INFO.b { return }
	golanglog.Println(a...)
}

func Infof(s string, a ...any) {
	if ! INFO.Ok() { return }
	if ! INFO.b { return }
	golanglog.Printf(s, a...)
}

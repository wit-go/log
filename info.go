package log

import 	(
	golanglog "log"
)

func Info(a ...any) {
	if ! INFO.B { return }
	golanglog.Println(a...)
}

func Infof(s string, a ...any) {
	if ! INFO.B { return }
	golanglog.Printf(s, a...)
}

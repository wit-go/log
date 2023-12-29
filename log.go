package log

import 	(
	origlog "log"
)

func Println(a ...any) {
	origlog.Println(a...)
}

func Fatalf(s string, a ...any) {
	origlog.Fatalf(s, a...)
}

func Fatal(s string, a ...any) {
	origlog.Fatalf(s, a...)
}

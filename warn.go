package log

import 	(
	origlog "log"
)

func Warn(a ...any) {
	if ! WARN.B { return }
	origlog.Println(a...)
}

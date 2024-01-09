package log

import 	(
	origlog "log"
)

func Warn(a ...any) {
	if ! WARN.Ok() { return }
	if ! WARN.b { return }
	origlog.Println(a...)
}

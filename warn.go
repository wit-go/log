package log

import 	(
	origlog "log"
)

func Warn(a ...any) {
	if ! WARN { return }
	origlog.Println(a...)
}

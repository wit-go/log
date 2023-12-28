package log

import 	(
	origlog "log"
)

func Warn(a ...any) {
	origlog.Println(a...)
}

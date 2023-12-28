package log

import 	(
	origlog "log"
)

func Error(a ...any) {
	origlog.Println(a...)
}

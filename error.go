package log

import 	(
	origlog "log"
)

func Error(err error, a ...any) {
	origlog.Println("Error:", err)
	origlog.Println(a...)
}

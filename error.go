package log

import 	(
	origlog "log"
)

func Error(err error, a ...any) {
	if ! ERROR.B { return }
	origlog.Println("Error:", err)
	origlog.Println(a...)
}

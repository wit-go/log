package log

import 	(
	origlog "log"
)

func Println(a ...any) {
	origlog.Println(a...)
}

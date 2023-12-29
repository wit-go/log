package log

import 	(
	golanglog "log"
)

func Bool(b bool, a ...any) {
	if ! b {return}
	golanglog.Println(a...)
}

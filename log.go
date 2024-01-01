package log

import 	(
	origlog "log"
)

/*

a simple way turn logging messages on and off. The gui config
window will let you enable & disable logging while your app is
running.

Example:

	log.Log(NETWARN, "socket connection failed to 127.0.0.1:8080")

In your package, register NETWARN:

	var NETWARN bool
	log.Register("myNetPkg", "NETWARN", &NETWARN)
*/

func Log(x any, a ...any) {
	if x == nil { return }
	switch x.(type) {
	case bool:
		if ! x.(bool) {
			return
		}
		origlog.Println(a...)
	case LogFlag:
		var f LogFlag
		f = x.(LogFlag)
		if ! f.B {
			return
		}
		a = append([]any{f.Subsystem}, a...)
		origlog.Println(a...)
	default:
		a = append([]any{x}, a...)
		origlog.Println(a...)
	}
}

func Logf(x any, s string, a ...any) {
	if x == nil { return }
	switch x.(type) {
	case bool:
		if ! x.(bool) {
			return
		}
	case LogFlag:
		var f LogFlag
		f = x.(LogFlag)
		if ! f.B {
			return
		}
	}
	origlog.Printf(s, a...)
}

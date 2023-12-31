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

func Log(b bool, a ...any) {
	if ! b { return }
	origlog.Println(a...)
}

func Logf(b bool, s string, a ...any) {
	if ! b { return }
	origlog.Printf(s, a...)
}

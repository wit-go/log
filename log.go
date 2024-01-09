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

	var NETWARN *log.LogFlag
	NETWARN = log.NewFlag("NETWARN", true, "go.wit.com/log", "log", "network warnings!")

*/

func Log(f *LogFlag, a ...any) {
	if ! f.Ok() { 
		// if the flag is NULL, notify the user they didn't initialize the flag
		a = append([]any{"FLAG = NULL"}, a...)
		origlog.Println(a...)
		return
	}
	if ! f.Get() { return }
	a = append([]any{f.short}, a...)
	origlog.Println(a...)
}

func Logf(f *LogFlag, s string, a ...any) {
	if ! f.Get() { return }
	s = f.short + " " + s
	origlog.Printf(s, a...)
}

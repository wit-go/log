package log

import 	(
	"errors"
	origlog "log"
)

var SPEW bool = false
var INFO bool = true
var WARN bool = true
var ERROR bool = true

func All(b bool) {
	Set("SPEW", b)
	Set("INFO", b)
	Set("WARN", b)
	Set("ERROR", b)
}

func Set(flag string, b bool) {
	switch flag {
	case "INFO":
		INFO = b
	case "WARN":
		WARN = b
	case "SPEW":
		SPEW = b
	case "ERROR":
		ERROR = b
	default:
		Error(errors.New("unknown flag"), "Flag name sent:", flag)
	}
}

func Get(flag string) bool {
	switch flag {
	case "INFO":
		return INFO
	case "WARN":
		return WARN
	case "SPEW":
		return SPEW
	case "ERROR":
		return ERROR
	default:
		Error(errors.New("unknown flag"), "Flag name sent:", flag)
	}
	return false
}

func Println(a ...any) {
	origlog.Println(a...)
}

func Fatalf(s string, a ...any) {
	origlog.Fatalf(s, a...)
}

func Fatal(s string, a ...any) {
	origlog.Fatalf(s, a...)
}

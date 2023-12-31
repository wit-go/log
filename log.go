package log

import 	(
	"errors"
	origlog "log"
)

var INFO bool = true
var WARN bool = true
var ERROR bool = true

var VERBOSE bool = false
var SPEW bool = false

func All(b bool) {
	Set("SPEW", b)
	Set("INFO", b)
	Set("WARN", b)
	Set("ERROR", b)
	Set("VERBOSE", b)
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
	case "VERBOSE":
		VERBOSE = b
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
	case "VERBOSE":
		return VERBOSE
	default:
		Error(errors.New("unknown flag"), "Flag name sent:", flag)
	}
	return false
}

// a simple way turn logging messages on and off
func Log(b bool, a ...any) {
	if ! b { return }
	origlog.Println(a...)
}

func Logf(b bool, s string, a ...any) {
	if ! b { return }
	origlog.Printf(s, a...)
}

func Println(a ...any) {
	origlog.Println(a...)
}

func Printf(s string, a ...any) {
	origlog.Printf(s, a...)
}

func Fatalf(s string, a ...any) {
	origlog.Fatalf(s, a...)
}

func Fatal(s string, a ...any) {
	origlog.Fatalf(s, a...)
}

package log

import 	(
	"errors"
)

var INFO bool = true
var WARN bool = true
var ERROR bool = true

var VERBOSE bool = false
var SPEW bool = false

var registered map[string][]string

func init() {
	registered = make(map[string][]string)
}

func All(b bool) {
	Set("SPEW", b)
	Set("INFO", b)
	Set("WARN", b)
	Set("ERROR", b)
	Set("VERBOSE", b)
}

func ListAll() map[string][]string {
	registered["all"] = []string{"SPEW","INFO", "WARN", "ERROR", "VERBOSE"}
	return registered
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

// register a variable name from a subsystem
// this is used for custom log flags
func Register(subsystem string, name string, b *bool) {
	Info("log.Register() got subsystem", subsystem, "with name =", name, "bool value =", b)
	registered[subsystem] = append(registered[subsystem], name)
}

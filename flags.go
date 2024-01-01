package log

import 	(
	"errors"
)

var INFO bool = true
var WARN bool = true
var ERROR bool = true

var VERBOSE bool = false
// var SPEW bool = false

var SPEW LogFlag

type LogFlag struct {
	B	bool
	Name	string
	Subsystem	string
	Desc	string
}

var registered map[string][]string
var flags []*LogFlag

func init() {
	registered = make(map[string][]string)

	SPEW.B = false
	SPEW.Name = "SPEW"
	SPEW.Subsystem = "log"
	SPEW.Desc = "Enable log.Spew()"

	// register the default flags used by this log package
	registered["log"] = []string{"SPEW","INFO", "WARN", "ERROR", "VERBOSE"}
}

func All(b bool) {
	Set("SPEW", b)
	Set("INFO", b)
	Set("WARN", b)
	Set("ERROR", b)
	Set("VERBOSE", b)
	for _, f := range flags {
		Warn("All() Setting Name =", f.Name, "Subsystem =", f.Subsystem, "to b =", b, "vs", f.B)
		f.B = b
		Warn("All() f.B is now", f.B)
	}
}

func ListFlags() map[string][]string {
	Info("ListFlags() registered =", registered)
	for _, f := range flags {
		Info("ListFlags() flag B =", f.B, "Name =", f.Name, "Subsystem =", f.Subsystem, "Description:", f.Desc)
	}

	return registered
}

func Set(flag string, b bool) {
	switch flag {
	case "INFO":
		INFO = b
	case "WARN":
		WARN = b
	case "SPEW":
		SPEW.B = b
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
		return SPEW.B
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
// inspired by Alex Flint
func (f *LogFlag) Register() {
	Info("log.Register() ", f)
	flags = append(flags,f)
}

// register a variable name from a subsystem
// this is used for custom log flags
func Register(subsystem string, name string, b *bool) {
	Info("log.Register() got subsystem", subsystem, "with name =", name, "bool value =", b)
	registered[subsystem] = append(registered[subsystem], name)
}

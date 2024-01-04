package log

/*
	Handles the on/off flags for things like log.Info() and log.Warn()
*/

import 	(
	"sync"
)

var INFO LogFlag
var VERBOSE LogFlag
var SPEW LogFlag
var WARN LogFlag
var ERROR LogFlag
var PRINTLN LogFlag

// writeMutex protects locks the write process
var flagsMutex sync.Mutex

type LogFlag struct {
	B	bool
	Name	string
	Subsystem	string
	Desc	string
}

var flags []*LogFlag

func init() {
	INFO.B = false
	INFO.Name = "INFO"
	INFO.Subsystem = "log"
	INFO.Desc = "Enable log.Info()"
	INFO.Register()

	SPEW.B = false
	SPEW.Name = "SPEW"
	SPEW.Subsystem = "log"
	SPEW.Desc = "Enable log.Spew()"
	SPEW.Register()

	VERBOSE.B = false
	VERBOSE.Name = "VERBOSE"
	VERBOSE.Subsystem = "log"
	VERBOSE.Desc = "Enable log.Verbose()"
	VERBOSE.Register()

	WARN.B = true
	WARN.Name = "WARN"
	WARN.Subsystem = "log"
	WARN.Desc = "Enable log.Warn()"
	WARN.Register()

	ERROR.B = true
	ERROR.Name = "ERROR"
	ERROR.Subsystem = "log"
	ERROR.Desc = "Enable log.Error()"
	ERROR.Register()

	PRINTLN.B = true
	PRINTLN.Name = "PRINTLN"
	PRINTLN.Subsystem = "log"
	PRINTLN.Desc = "Enable log.Println()"
	PRINTLN.Register()
}

// set all the flags
func SetAll(b bool) {
	flagsMutex.Lock()
	defer flagsMutex.Unlock()
	for _, f := range flags {
		f.B = b
	}
}


// this bypasses all checks and _always_ logs the info to STDOUT
// is this a bad idea? Probably not....
// TODO: returning []*LogFlag is not safe and access must be locked
// but this is only used by the log debugging window at this time
func ShowFlags() []*LogFlag {
	flagsMutex.Lock()
	defer flagsMutex.Unlock()
	for _, f := range flags {
		Log(true, "ShowFlags() ", "(" + f.Subsystem + ")", f.Name, "=", f.B, ":", f.Desc)
	}

	return flags
}

// TODO, switch to this 
func ProcessFlags(callback func(*LogFlag)) {
	flagsMutex.Lock()
	defer flagsMutex.Unlock()
	for _, f := range flags {
		Log(true, "ProcessFlags() run callback(f) here on", f)
		callback(f)
	}

}


// register a variable name from a subsystem
// inspired by Alex Flint
func (f *LogFlag) Register() {
	flagsMutex.Lock()
	defer flagsMutex.Unlock()
	Info("log.Register() ", f)
	flags = append(flags,f)
}

func (f *LogFlag) Set(b bool) {
	flagsMutex.Lock()
	defer flagsMutex.Unlock()
	Info("Set() ", "(" + f.Subsystem + ")", f.Name, "=", f.B, ":", f.Desc)
	f.B = b
	Info("Set() f.B is now", f.B)
}

func Set(subsystem string, name string, b bool) {
	flagsMutex.Lock()
	defer flagsMutex.Unlock()
	Info("Set() TODO find var:", "(" + subsystem + ")", name, "=", b)
	for _, f := range flags {
		Log(true, "Set() ", "(" + f.Subsystem + ")", f.Name, "=", f.B, ":", f.Desc)
		if (subsystem == f.Subsystem) && (name == f.Name) {
			Log(true, "Set() FOUND ", f)
			f.B = b
			return
		}
	}

}

func Get(subsystem string, name string) bool {
	flagsMutex.Lock()
	defer flagsMutex.Unlock()
	Info("Get() TODO find var:", "(" + subsystem + ")", name)
	for _, f := range flags {
		Log(true, "Get() ", "(" + f.Subsystem + ")", f.Name, "=", f.B, ":", f.Desc)
		if (subsystem == f.Subsystem) && (name == f.Name) {
			Log(true, "Get() FOUND ", f)
			return f.B
		}
	}
	return false
}

package log

/*
	Handles the on/off flags for things like log.Info() and log.Warn()
*/

/*
	The original log flags:

	log.Ldate: The date in the local time zone: YYYY/MM/DD.
	log.Ltime: The time in the local time zone: HH:MM:SS.
	log.Lmicroseconds: Microsecond resolution: HH:MM:SS.microseconds.
	log.Llongfile: Full file name and line number: /a/b/c/d.go:23.
	log.Lshortfile: Final file name element and line number: d.go:23.
	log.LUTC: If Ldate or Ltime is set, use UTC rather than the local time zone.
	log.Lmsgprefix: Move the "prefix" from the beginning of the line to before the message.
	log.LstdFlags: Initial values for the standard logger (Ldate | Ltime).

	can be set this way:
	myLogger.SetFlags(log.Ldate | log.Ltime)

*/



import 	(
	"sync"
)

var INFO *LogFlag	// toggles log.Info()
var VERBOSE *LogFlag	// toggles log.Verbose()
var SPEW *LogFlag	// toggles log.Spew()

var WARN *LogFlag	// toggles log.Warn() (true by default)
var ERROR *LogFlag	// toggles log.Warn() (true by default)
var PRINTLN *LogFlag	// toggles log.Println() (true by default)

var always *LogFlag

// writeMutex protects locks the write process
var flagsMutex sync.Mutex

type LogFlag struct {
	b	bool
	orig	bool	// used as the Default value. set at the time of Registration()
	name	string
	// TODO: figure out what package is sending the Registration
	subsystem	string // probably should just be forced to be the package name
	short	string	// string actually printed on each line
	desc	string
}

var flags []*LogFlag

func init() {
	full := "go.wit.com/log"
	short := "log"

	INFO = NewFlag("INFO", false, full, short, "Enable log.Info()")
	SPEW = NewFlag("SPEW", false, full, short, "Enable log.Spew()")
	WARN = NewFlag("WARN", true,  full, short, "Enable log.Warn()")

	ERROR = NewFlag("ERROR",   true,   full, short, "Enable log.Error()")
	PRINTLN = NewFlag("PRINTLN", true,   full, short, "Enable log.Println()")
	VERBOSE = NewFlag("VERBOSE", false,  full, short, "Enable log.Verbose()")

	// internally used to bypass the possibility that all the flags are off
	always = new(LogFlag)
	always.b = true
	always.orig = true
	always.subsystem = full
	always.short = short
	always.desc = "internal only"
}

// restores flag to it's default value
func (f *LogFlag) SetDefault() {
	if ! f.Ok() {return}
	f.b = f.orig
}

// set all the flags
func SetDefaults() {
	flagsMutex.Lock()
	defer flagsMutex.Unlock()
	for _, f := range flags {
		f.SetDefault()
	}
}

// protects against panic() by making sure it exists.
func (f *LogFlag) Ok() bool {
	if f == nil {return false}
	return true
}

// set all the flags
func SetAll(b bool) {
	flagsMutex.Lock()
	defer flagsMutex.Unlock()
	for _, f := range flags {
		f.b = b
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
		Log(always, "ShowFlags() ", "(" + f.subsystem + ")", f.name, "=", f.b, ":", f.desc)
	}

	return flags
}

// TODO, switch to this. maybe.
func ProcessFlags(callback func(*LogFlag)) {
	flagsMutex.Lock()
	defer flagsMutex.Unlock()
	for _, f := range flags {
		Log(always, "ProcessFlags() run callback(f) here on", f)
		callback(f)
	}

}

// returns the value of the flag
func (f *LogFlag) Get() bool {
	if ! f.Ok() {return false}
	return f.b
}

// returns the name of the flag
func (f *LogFlag) GetName() string {
	if ! f.Ok() {return ""}
	return f.name
}

// returns the subsystem of the flag
func (f *LogFlag) GetSubsystem() string {
	if ! f.Ok() {return ""}
	return f.subsystem
}

// returns the description of the flag
func (f *LogFlag) GetDesc() string {
	if ! f.Ok() {return ""}
	return f.desc
}

// register a variable name from a subsystem
// inspired by Alex Flint
// set the Default value at the time of registration

// this is what the current log.SetFlag() function should become
func NewFlag(name string, b bool, full, short, desc string) *LogFlag {
	flagsMutex.Lock()
	defer flagsMutex.Unlock()
	f := new(LogFlag)
	Log(always, "log.SetFlag() ", full, short, name, true)
	f.b = b
	f.orig = b
	f.short = short
	f.subsystem = full
	f.name = name
	f.desc = desc
	flags = append(flags,f)
	return f
}

func (f *LogFlag) Set(b bool) {
	if ! f.Ok() {return}
	flagsMutex.Lock()
	defer flagsMutex.Unlock()
	Info("Set() ", "(" + f.subsystem + ")", f.name, "=", f.b, ":", f.desc)
	f.b = b
	Info("Set() f.b is now", f.b)
}

/*
func Set(subsystem string, name string, b bool) {
	flagsMutex.Lock()
	defer flagsMutex.Unlock()
	Verbose("log.Set() TODO find var:", "(" + subsystem + ")", name, "=", b)
	for _, f := range flags {
		Verbose("log.Set() ", "(" + f.subsystem + ")", f.name, "=", f.b, ":", f.desc)
		if (subsystem == f.subsystem) && (name == f.name) {
			Verbose("log.Set() FOUND ", f)
			f.b = b
			return
		}
	}

}

func Get(subsystem string, name string) bool {
	flagsMutex.Lock()
	defer flagsMutex.Unlock()
	Verbose("log.Get() TODO find var:", "(" + subsystem + ")", name)
	for _, f := range flags {
		Verbose("log.Get() ", "(" + f.subsystem + ")", f.name, "=", f.b, ":", f.desc)
		if (subsystem == f.subsystem) && (name == f.name) {
			Verbose("log.Get() FOUND ", f)
			return f.b
		}
	}
	return false
}
*/

package log

import (
	arg "github.com/alexflint/go-arg"
)

//
// Attempt to switch logging to syslog on linux
//

var argLog ArgLog

// This struct can be used with the go-arg package
type ArgLog struct {
	LogTmp bool `arg:"--log-tmp" help:"send all output /tmp"`
	LogStdout bool `arg:"--log-stdout" help:"send all output to STDOUT"`
	LogQuiet bool `arg:"--log-quiet" help:"suppress all output"`
}

func init() {
	arg.Register(&argLog)
}

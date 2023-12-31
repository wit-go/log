package log
//
// version v1.2
//
// I like things to be easy.
//
// this means all the log settings are in one place. it should allow
// things to be over-ridden externally to the library
// but still allow command line --args to pass debugging settings
//
// I also have a generic sleep() and exit() in here because it's simple
//
// Usage:
//
// log("something", foo, bar)
// var DEBUG bool = true
// log(DEBUG, "something else", someOtherVariable)  # if DEBUG == false, return doing nothing
//

/*
	I've spent, am spending, too much time thinking about 'logging'. 'log', 'logrus', 'zap', whatever.
	I'm not twitter. i don't give a fuck about how many nanoseconds it takes to log. Anyway, this
	implementation is probably faster than all of those because you just set one bool to FALSE
	and it all stops.
	Sometimes I need to capture to stdout, sometimes stdout can't
	work because it doesn't exist for the user. This whole thing is a PITA. Then it's spread
	over 8 million references in every .go file. I'm tapping out and putting
	it in one place. here it is. Also, this makes having debug levels really fucking easy.
	You can define whatever level of logging you want from anywhere (command line) etc.

	log()		# doesn't do anything
	log(stuff)	# sends it to whatever log you define in a single place. here is the place
*/

import 	(
	"os"
	"time"
	"errors"
	"reflect"

	origlog "log"
)

/*
	sleep()		# you know what this does? sleeps for 1 second. yep. dump. easy.
	sleep(.1)	# you know what this does? yes, it sleeps for 1/10th of a second
*/
func Sleep(a ...any) {
	if (a == nil) {
		time.Sleep(time.Second)
		return
	}

	origlog.Println("sleep", a[0])

	switch a[0].(type) {
	case int:
		time.Sleep(time.Duration(a[0].(int)) * time.Second)
	case float64:
		time.Sleep(time.Duration(a[0].(float64) * 1000) * time.Millisecond)
	default:
		origlog.Println("sleep a[0], type = ", a[0], reflect.TypeOf(a[0]))
	}
}

/*
	exit()		# yep. exits. I guess everything must be fine
	exit(3)		# I guess 3 it is then
	exit("dont like apples")	# ok. I'll make a note of that
*/
func Exit(a ...any) {
	Error(errors.New("Exit"), a)
	//if (a) {
	//	os.Exit(a)
	//}
	os.Exit(0)
}

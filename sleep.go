package log

// a shortcut for sleep so you don't have to always change the import lines when debugging

import 	(
	"os"
	"time"
	"errors"
	"reflect"
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

	Info("sleep", a[0])

	switch a[0].(type) {
	case int:
		time.Sleep(time.Duration(a[0].(int)) * time.Second)
	case float64:
		time.Sleep(time.Duration(a[0].(float64) * 1000) * time.Millisecond)
	default:
		Info("sleep a[0], type = ", a[0], reflect.TypeOf(a[0]))
	}
}

/*
	exit()		# yep. exits. I guess everything must be fine
	exit(3)		# I guess 3 it is then
	exit("dont like apples")	# ok. I'll make a note of that
*/
func Exit(a ...any) {
	Error(errors.New("log.Exit()"), a)
	//if (a) {
	//	os.Exit(a)
	//}
	os.Exit(0)
}

package log

import 	(
	"os"
	golanglog "log"
)

// start writing all the logging to a tmp file
func SetTmp() {
	f, err := os.OpenFile("/tmp/guilogfile", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		golanglog.Fatalf("error opening file: %v", err)
	}
	// hmm. is there a trick here or must this be in main()
	// defer f.Close()

	golanglog.SetOutput(f)
	golanglog.Println("This is a test log entry")
}

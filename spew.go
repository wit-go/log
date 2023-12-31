package log

import 	(
	origlog "log"
	"github.com/davecgh/go-spew/spew"
)

func Spew(a ...any) {
	if ! SPEW { return }
	origlog.Println("SPEW:", spew.Sdump(a...))

	/*
	scs := spew.ConfigState{Indent: "\t", MaxDepth: 1}
	// Output using the ConfigState instance.
	v := map[string]int{"one": 1}
	scs.Printf("v: %v\n", v)
	scs.Dump(v)
	scs.Dump(a)
	*/
}

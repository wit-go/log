package log

import 	(
	origlog "log"
	"github.com/davecgh/go-spew/spew"
)

func Spew(b any, a ...any) {
	if ! SPEW.b { return }

	switch b.(type) {
	case bool:
		if ! b.(bool) {
			return
		}
		origlog.Println("SPEW:", spew.Sdump(a...))
	case logFlag:
		var f logFlag
		f = b.(logFlag)
		if ! f.b {
			return
		}
		origlog.Println("SPEW:", spew.Sdump(a...))
	default:
		origlog.Println("SPEW b:", spew.Sdump(b))
		origlog.Println("SPEW a:", spew.Sdump(a...))
	}
	// origlog.Println("SPEW:", spew.Sdump(a...))
	/*
	scs := spew.ConfigState{Indent: "\t", MaxDepth: 1}
	// Output using the ConfigState instance.
	v := map[string]int{"one": 1}
	scs.Printf("v: %v\n", v)
	scs.Dump(v)
	scs.Dump(a)
	*/
}

package main

// #cgo darwin LDFLAGS: -lMyLib -L${SRCDIR} -lfoo
// #include "foo.h"
import "C"

type GoFoo struct {
	foo C.Foo
}

//export add
func add() *C.char {
	return C.CString("hello")
}
func New() GoFoo {
	var ret GoFoo
	ret.foo = C.FooInit()
	return ret
}
func (f GoFoo) Free() {
	C.FooFree(f.foo)
}
func (f GoFoo) Bar() {
	C.FooBar(f.foo)
}

func main() {
	//foo := New()
	//foo.Bar()
	//foo.Free()
}

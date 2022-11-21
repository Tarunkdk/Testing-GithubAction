package main

import (
	"fmt"
	_ "golinker/inner"
	_ "runtime"
	_ "unsafe"
)

//go:linkname hello golinker/inner.hello
func hello()

//go:linkname runtimeG runtime.g
type runtimeG struct {
}

func main() {
	hello()
	r := runtimeG{}
	fmt.Println(r)
	fmt.Println(r)
}

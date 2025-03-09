package tests

import (
	"fmt"
	"runtime"
)

func testversio() {
	fmt.Printf("Go version:%s\n", runtime.Version())
}

package tests

import (
	"fmt"
	"runtime"
)

func testversion() {
	fmt.Printf("Go version:%s\n", runtime.Version())
}

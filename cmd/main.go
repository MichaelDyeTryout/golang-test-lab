package main

import (
	"fmt"
	"os"

	"github.com/michaeldye/golang-test-lab/pkg/sequence"
)

// An entrypoint so we don't unit test this. We don't need an explicit
// "//go:cover ignore" directive b/c our `go test ...` invocation ignores main
// by default.
func main() {

	numbaz, err := sequence.CountingNumbers(9)
	if err != nil {
		panic(err)
	}

	println("Values generated:")
	fmt.Printf("%v\n", *numbaz)

	// below will panic when ix == len(*numbaz)
	// note: neither golang staticcheck nor sonarqube will catch this

	// we use an instrumented integration binary to get coverage from this run
	// (see https://go.dev/doc/build-cover). !Important! We must handle a panic
	// here or we won't get generated converage data.
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Recovered from panic: %v\n", err)
		}
		os.Exit(1)
	}()

	println("Iterating over values unsafely:")
	for ix := 0; ix <= len(*numbaz); ix++ {
		fmt.Printf("%d\n", (*numbaz)[ix])
	}
}

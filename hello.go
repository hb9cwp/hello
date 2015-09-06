package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Hello World!")
	fmt.Printf("I'm running on %s with %s architecture\n", runtime.GOOS, runtime.GOARCH)
}

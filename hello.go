package main

import (
	"fmt"
	"syscall/js"
)

func main() {
	// Print to browser console
	fmt.Println("Hello, World! (logged in console)")

	// Write to the webpage
	js.Global().Get("document").Call("write", "Hello, World! (rendered on page)")
}


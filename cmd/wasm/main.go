package main

import (
	"fmt"
	"syscall/js"
)

func main() {
	c := make(chan bool)

	fmt.Println("Hello, WebAssembly!")
	fmt.Println(count)

	js.Global().Set("increment", js.FuncOf(increment))
	js.Global().Set("decrement", js.FuncOf(decrement))

	<-c
}

var count int

func increment(this js.Value, args []js.Value) interface{} {
	count++
	return count
}

func decrement(this js.Value, args []js.Value) interface{} {
	count--
	return count
}

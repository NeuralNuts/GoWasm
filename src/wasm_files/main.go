package main

import (
	"syscall/js"
)

func add(i []js.Value) {
    js.Global().Set("output", js.ValueOf(i[0].Int() + i[1].Int()))
}

func registerCallbacks() {
    js.Global().Set("addNumbers", js.NewCallback(add))
}

func main() {
    c := make(chan struct{}, 0)
    registerCallbacks()
    <-c
}

package main

import (
	"syscall/js"
)

func main() {
	quit := make(chan struct{}, 0)

	stopButton := js.Global().Get("document").Call("getElementById", "stop")
	stopButton.Set("disabled", false)
	stopButton.Set("onclick", js.NewCallback(func([]js.Value) {
		println("stopping")
		stopButton.Set("disabled", true)
		quit <- struct{}{}
	}))

	<-quit
}

package main

import (
	"syscall/js"
)

func main() {
	quit := make(chan struct{}, 0)

	stopButton := js.Global().Get("document").Call("getElementById", "stop")
	stopButton.Set("disabled", false)
	stopButton.Set("onclick", js.FuncOf(func(js.Value, []js.Value) interface{} {
		println("stopping")
		stopButton.Set("disabled", true)
		quit <- struct{}{}
		return nil
	}))

	<-quit
}

package main

import (
	"net/http"
)

func main() {
	DownloadWasmExec()
	fileServer := http.FileServer(http.Dir("."))
	http.Handle("/", fileServer)
	println("Listening on port 3000...")
	http.ListenAndServe(":3000", nil)
}

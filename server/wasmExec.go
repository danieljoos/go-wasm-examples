package main

import (
	"io"
	"net/http"
	"os"
)

const wasmExecURL = "https://raw.githubusercontent.com/golang/go/master/misc/wasm/wasm_exec.js"
const wasmExecFile = "wasm_exec.js"

func DownloadWasmExec() {
	if _, err := os.Stat(wasmExecFile); err == nil {
		return
	}
	println("Downloading wasm_exec.js...")
	out, err := os.Create(wasmExecFile)
	if err != nil {
		panic(err)
	}
	defer out.Close()
	resp, err := http.Get(wasmExecURL)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		panic(err)
	}
}

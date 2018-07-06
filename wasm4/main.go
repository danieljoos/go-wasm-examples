package main

import (
	"syscall/js"

	"github.com/shurcooL/github_flavored_markdown"
)

var document = js.Global().Get("document")

func getElementByID(id string) js.Value {
	return document.Call("getElementById", id)
}

func renderEditor(parent js.Value) js.Value {
	editorMarkup := `
		<div id="editor" style="display: flex; flex-flow: row wrap;">
			<textarea id="markdown" style="width: 50%; height: 400px"></textarea>
			<div id="preview" style="width: 50%;"></div>
			<button id="render">Render Markdown</button>
		</div>
	`
	parent.Call("insertAdjacentHTML", "beforeend", editorMarkup)
	return getElementByID("editor")
}

func main() {
	quit := make(chan struct{}, 0)

	// See example 2: Enable the stop button
	stopButton := getElementByID("stop")
	stopButton.Set("disabled", false)
	stopButton.Set("onclick", js.NewCallback(func([]js.Value) {
		println("stopping")
		stopButton.Set("disabled", true)
		quit <- struct{}{}
	}))

	// Simple markdown editor
	editor := renderEditor(document.Get("body"))
	markdown := getElementByID("markdown")
	preview := getElementByID("preview")
	renderButton := getElementByID("render")
	renderButton.Set("onclick", js.NewCallback(func([]js.Value) {
		// Process markdown input and show the result
		md := markdown.Get("value").String()
		html := github_flavored_markdown.Markdown([]byte(md))
		preview.Set("innerHTML", string(html))
	}))

	<-quit
	editor.Call("remove")
}
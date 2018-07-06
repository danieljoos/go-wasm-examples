package main

import(
	"syscall/js"

	"github.com/PaulRosset/go-hacknews"
)

func topStories() []hacknews.Post {
	// This is an example from:
	// https://github.com/PaulRosset/go-hacknews
	// Copyright (c) 2017 Paul Rosset
	// See LICENSE in the link above.
	init := hacknews.Initializer{Story: "topstories", NbPosts: 10}
	codes, err := init.GetCodesStory()
	if err != nil {
		panic(err)
	}
	posts, err := init.GetPostStory(codes)
	if err != nil {
		panic(err)
	}
	return posts
}

var document = js.Global().Get("document")

func renderPost(post hacknews.Post, parent js.Value) {
	// Creates the following HTML elements
	// <li>
	//   <a href="{URL}">{Title}</a>
	// </li>
	li := document.Call("createElement", "li")
	a := document.Call("createElement", "a")
	text := document.Call("createTextNode", post.Title)
	a.Set("href", post.Url)
	a.Call("appendChild", text)
	li.Call("appendChild", a)
	parent.Call("appendChild", li)
}

func renderPosts(posts []hacknews.Post, parent js.Value) {
	ul := document.Call("createElement", "ul")
	parent.Call("appendChild", ul)
	for _, post := range(posts) {
		renderPost(post, ul)
	}
}

func main() {
	posts := topStories()
	body := document.Get("body")
	renderPosts(posts, body)
}
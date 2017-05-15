package main

import (
	"javascript/components"

	"github.com/gopherjs/gopherjs/js"
)

func main() {
	// app := router.GetRouter()
	app := components.Home()
	root := js.Global.Get("document").Call("getElementById", "root")
	js.Global.Get("ReactDOM").Call("render", app, root)
}

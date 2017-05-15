package main

import (
	"javascript/router"

	"github.com/gopherjs/gopherjs/js"
)

func main() {
	app := router.GetRouter()
	root := js.Global.Get("document").Call("getElementById", "root")
	js.Global.Get("ReactDOM").Call("render", app, root)
}

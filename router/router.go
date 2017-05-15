package router

import (
	"javascript/components"
	"javascript/react"

	"github.com/gopherjs/gopherjs/js"
)

// GetRouter - TODO
func GetRouter() *js.Object {
	HashRouter := js.Global.Get("ReactRouterDOM").Get("HashRouter")
	Switch := js.Global.Get("ReactRouter").Get("Switch")
	Route := js.Global.Get("ReactRouter").Get("Switch")

	router :=
		react.CreateElement(HashRouter, nil, [1]*js.Object{
			react.CreateElement(Switch, nil, [1]*js.Object{
				react.CreateElement(Route, react.Props(map[string]interface{}{
					"path":      "/",
					"component": components.Home,
				}), nil),
			}),
		})

	return router
}

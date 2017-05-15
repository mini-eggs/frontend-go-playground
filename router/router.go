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
	Route := js.Global.Get("ReactRouter").Get("Route")

	router :=
		react.CreateElement(HashRouter, nil,
			react.CreateElement("div", nil, [2]*js.Object{
				react.CreateElement(components.Navigation, map[string]interface{}{
					"key": 0,
				}, nil),
				react.CreateElement(Switch, map[string]interface{}{
					"key": 1,
				}, [2]*js.Object{
					react.CreateElement(Route, map[string]interface{}{
						"key":       0,
						"exact":     true,
						"path":      "/",
						"component": components.Home,
					}, nil),
					react.CreateElement(Route, map[string]interface{}{
						"key":       1,
						"path":      "/about",
						"component": components.About,
					}, nil),
				}),
			}),
		)

	return router
}

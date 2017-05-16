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
		r.CreateElement(HashRouter, nil,
			r.CreateElement("div", nil, [2]*js.Object{
				r.CreateElement(components.Navigation, r.Props{
					"key": 0,
				}, nil),
				r.CreateElement(Switch, r.Props{
					"key": 1,
				}, [2]*js.Object{
					r.CreateElement(Route, r.Props{
						"key":       0,
						"exact":     true,
						"path":      "/",
						"component": components.Home,
					}, nil),
					r.CreateElement(Route, r.Props{
						"key":       1,
						"path":      "/about",
						"component": components.About,
					}, nil),
				}),
			}),
		)

	return router
}

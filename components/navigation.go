package components

import (
	"javascript/react"

	"github.com/gopherjs/gopherjs/js"
)

// Navigation - TODO
func Navigation() *js.Object {

	link := js.Global.Get("ReactRouterDOM").Get("Link")

	component := new(react.Component)

	component.Render = func() *js.Object {
		markup :=
			react.CreateElement("div", map[string]interface{}{
				"className": "navigation-wrapper",
			}, [3]*js.Object{
				react.CreateElement(link, map[string]interface{}{
					"key": 0,
					"to":  "/",
				}, "home"),
				react.CreateElement("br", map[string]interface{}{
					"key": 1,
				}, nil),
				react.CreateElement(link, map[string]interface{}{
					"key": 2,
					"to":  "/about",
				}, "about"),
			})
		return markup
	}

	return react.ReturnComponent(component, nil, nil)
}

package components

import (
	"javascript/react"

	"github.com/gopherjs/gopherjs/js"
)

// Navigation - TODO
func Navigation() *js.Object {

	link := js.Global.Get("ReactRouterDOM").Get("Link")

	component := new(r.Component)

	component.Render = func() *js.Object {
		return r.CreateElement("header", nil,
			r.CreateElement("nav", nil,
				r.CreateElement("ul", nil, [2]*js.Object{
					r.CreateElement("li", r.Props{"key": 0},
						r.CreateElement(link, r.Props{"to": "/"}, "home"),
					),
					r.CreateElement("li", r.Props{"key": 1},
						r.CreateElement(link, r.Props{"to": "/about"}, "about"),
					),
				}),
			),
		)
	}

	return r.ReturnComponent(component, nil, nil)
}

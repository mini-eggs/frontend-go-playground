package components

import (
	"javascript/react"

	"github.com/gopherjs/gopherjs/js"
)

// Home - TODO
func Home() *js.Object {

	component := new(r.Component)

	component.Render = func() *js.Object {
		markup :=
			r.CreateElement("div", nil, [3]*js.Object{
				r.CreateElement("i", r.Props{"key": 0}, "hello"),
				r.CreateElement("br", r.Props{"key": 1}, nil),
				r.CreateElement("i", r.Props{"key": 2}, "world"),
			})
		return markup
	}

	return r.ReturnComponent(component, nil, nil)
}

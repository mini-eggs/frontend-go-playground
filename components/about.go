package components

import (
	"javascript/react"

	"github.com/gopherjs/gopherjs/js"
)

// About - TODO
func About() *js.Object {

	component := new(r.Component)

	component.Render = func() *js.Object {
		markup :=
			r.CreateElement("div", nil, r.CreateElement("i", nil, "about"))
		return markup
	}

	return r.ReturnComponent(component, nil, nil)
}

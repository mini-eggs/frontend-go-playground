package components

import (
	"javascript/react"

	"github.com/gopherjs/gopherjs/js"
)

// About - TODO
func About() *js.Object {

	component := new(react.Component)

	component.Render = func() *js.Object {
		markup :=
			react.CreateElement("div", nil, react.CreateElement("i", nil, "about"))
		return markup
	}

	return react.ReturnComponent(component, nil, nil)
}

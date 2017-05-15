package components

import (
	"javascript/react"

	"github.com/gopherjs/gopherjs/js"
)

// Home - TODO
func Home() *js.Object {

	component := new(react.Component)

	component.Render = func() *js.Object {
		markup :=
			react.CreateElement("div", nil, [3]*js.Object{
				react.CreateElement("i", map[string]interface{}{"key": 0}, "hello"),
				react.CreateElement("br", map[string]interface{}{"key": 1}, nil),
				react.CreateElement("i", map[string]interface{}{"key": 2}, "world"),
			})
		return markup
	}

	return react.ReturnComponent(component, nil, nil)
}

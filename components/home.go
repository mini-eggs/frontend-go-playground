package components

import (
	"javascript/react"

	"fmt"

	"github.com/gopherjs/gopherjs/js"
)

// Home - TODO
func Home() *js.Object {

	component := new(react.Component)

	component.ComponentDidMount = func() {
		fmt.Println("I have been mounted - home")
	}

	component.Render = func() *js.Object {
		return react.CreateElement("div", nil, [3]*js.Object{
			react.CreateElement("i", map[string]interface{}{"key": 0}, "hello"),
			react.CreateElement("br", map[string]interface{}{"key": 1}, nil),
			react.CreateElement("i", map[string]interface{}{"key": 2}, "world"),
		})
	}

	return react.ReturnComponent(component, nil, nil)
}

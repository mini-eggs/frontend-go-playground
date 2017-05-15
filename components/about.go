package components

import (
	"javascript/react"

	"fmt"

	"github.com/gopherjs/gopherjs/js"
)

// About - TODO
func About() *js.Object {

	component := new(react.Component)

	component.ComponentDidMount = func() {
		fmt.Println("I have been mounted - about")
	}

	component.Render = func() *js.Object {
		return react.CreateElement("div", nil, react.CreateElement("i", nil, "about"))
	}

	return react.ReturnComponent(component, nil, nil)
}

package components

import (
	"javascript/react"

	"fmt"

	"github.com/gopherjs/gopherjs/js"
)

// Home - TODO
func Home() *js.Object {

	testComponent := new(react.Component)

	testComponent.ComponentDidMount = func() {
		fmt.Println("I have been mounted")
	}

	testComponent.Render = func() *js.Object {
		return react.CreateElement("i", react.Props(map[string]interface{}{"key": 0}), "world")
	}

	test := react.CreateElement(js.Global.Get("window").Call("CreateComponent", testComponent), nil, nil)

	home :=
		react.CreateElement("div", react.Props(map[string]interface{}{
			"onClick": func() {
				fmt.Println("I have been clicked")
			},
			"style": map[string]interface{}{
				"display": "inline-block",
			},
		}), [2]*js.Object{
			react.CreateElement("div", react.Props(map[string]interface{}{"key": 0}), [1]*js.Object{
				react.CreateElement("i", react.Props(map[string]interface{}{"key": 0}), "hello"),
			}),
			react.CreateElement("div", react.Props(map[string]interface{}{"key": 1}), [1]*js.Object{
				react.CreateElement("i", react.Props(map[string]interface{}{"key": 0}), test),
			}),
		})

	return home
}

package r

import (
	"github.com/gopherjs/gopherjs/js"
)

// Props- TODO
type Props map[string]interface{}

// CreateElement - TODO
func CreateElement(elementType interface{}, props interface{}, children interface{}) *js.Object {
	return js.Global.Get("React").Call("createElement", elementType, props, children)
}

// Component - TODO
type Component struct {
	ComponentDidMount func()
	Render            func() *js.Object
}

// ReturnComponent - TODO
func ReturnComponent(component *Component, props *js.Object, children interface{}) *js.Object {
	return CreateElement(js.Global.Get("window").Call("CreateComponent", component), props, children)
}

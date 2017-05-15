package react

import (
	"github.com/gopherjs/gopherjs/js"
)

// Component - TODO
type Component struct {
	ComponentDidMount func()
	Render            func() *js.Object
}

// CreateElement - TODO
func CreateElement(elementType interface{}, props interface{}, children interface{}) *js.Object {
	return js.Global.Get("React").Call("createElement", elementType, props, children)
}

// Props - TODO
func Props(someProps map[string]interface{}) *js.Object {
	return js.Global.Get("Object").New(someProps)
}

// Initial - TODO
func Initial() *js.Object {
	return Props(map[string]interface{}{})
}

// ReturnComponent - TODO
func ReturnComponent(component *Component, props *js.Object, children interface{}) *js.Object {
	return CreateElement(js.Global.Get("window").Call("CreateComponent", component), props, children)
}

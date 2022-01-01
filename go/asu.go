package asu

import "github.com/gopherjs/gopherjs/js"

func NewJSObject() *js.Object {
	g := js.Global

	if g == nil {
		return nil
	}

	return g.Get("Object").New()
}

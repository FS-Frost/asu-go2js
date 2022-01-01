package asu

import "github.com/gopherjs/gopherjs/js"

type baseTag struct {
	*js.Object
	jsName       interface{} `js:"name"`
	jsFromArgs   interface{} `js:"fromArgs"`
	jsFromString interface{} `js:"fromString"`
	jsToString   interface{} `js:"toString"`
}

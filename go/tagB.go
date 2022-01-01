package asu

import (
	"fmt"

	"github.com/gopherjs/gopherjs/js"
	"github.com/oriser/regroup"
)

const regexTagB = `(?P<tag>\\b\s*(?P<arg>[0-1]))`

type TagB struct {
	baseTag
	Argument int `regroup:"arg,required" js:"argument"`
}

func (t *TagB) Name() string {
	return string(TAG_B)
}

func NewTagB() *TagB {
	t := &TagB{}
	t.Object = NewJSObject()

	if t.Object == nil {
		return t
	}

	t.Argument = 0
	t.jsFromString = t.FromString
	t.jsToString = t.ToString

	return t
}

// func (t *TagB) toJS() *js.Object {
// 	tJS := &TagB{}
// 	tJS.Object = t.Object
// 	tJS.Argument = t.Argument
// 	tJS.jsName = t.Name
// 	tJS.jsFromArgs = t.FromArgs
// 	tJS.jsFromString = t.FromString
// 	tJS.jsToString = t.ToString
// 	return tJS.Object
// }

func NewTagBJS() *js.Object {
	t := NewTagB()
	return t.Object
}

func (t *TagB) FromArgs(n int) *TagB {
	t.Argument = n
	return t
}

func (t *TagB) FromString(text string) (*TagB, error) {
	re := regroup.MustCompile(regexTagB)
	err := re.MatchToTarget(text, t)

	if err != nil {
		return nil, formatTagParsingError(t, err)
	}

	return t, nil
}

func (t *TagB) ToString() string {
	s := fmt.Sprintf(`\%s%d`, t.Name(), t.Argument)
	return s
}

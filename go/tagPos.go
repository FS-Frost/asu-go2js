package asu

import (
	"fmt"

	"github.com/gopherjs/gopherjs/js"
	"github.com/oriser/regroup"
)

const regexPos = `(?P<tag>\\pos\s*(?P<arg>\(\s*(?P<x>-?\d+(?:\.\d+)?),\s*(?P<y>-?\d+(?:\.\d+)?)\)))`

type TagPos struct {
	baseTag
	X float64 `regroup:"x,required" js:"x"`
	Y float64 `regroup:"y,required" js:"y"`
}

func NewTagPos() *TagPos {
	t := &TagPos{}
	t.Object = NewJSObject()

	if t.Object == nil {
		return t
	}

	t.X = 0
	t.Y = 0
	t.jsName = t.Name
	t.jsFromArgs = t.FromArgs
	t.jsFromString = t.FromString
	t.jsToString = t.ToString

	return t
}

// func (t *TagPos) toJS() *js.Object {
// 	tJS := &TagPos{}
// 	tJS.Object = t.Object
// 	tJS.X = t.X
// 	tJS.Y = t.Y
// 	tJS.jsName = t.Name
// 	tJS.jsFromArgs = t.FromArgs
// 	tJS.jsFromString = t.FromString
// 	tJS.jsToString = t.ToString
// 	return tJS.Object
// }

func NewTagPosJS() *js.Object {
	t := NewTagPos()
	return t.Object
}

func (t *TagPos) Name() string {
	return string(TAG_POS)
}

func (t *TagPos) FromArgs(x float64, y float64) *TagPos {
	t.X = x
	t.Y = y
	return t
}

func (t *TagPos) FromString(text string) (*TagPos, error) {
	re := regroup.MustCompile(regexPos)
	err := re.MatchToTarget(text, t)

	if err != nil {
		return nil, formatTagParsingError(t, err)
	}

	return t, nil
}

func (t *TagPos) ToString(text string) string {
	s := fmt.Sprintf(`\%s(%g,%g)`, t.Name(), t.X, t.Y)
	return s
}

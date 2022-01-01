package asu

import (
	"fmt"

	"github.com/gopherjs/gopherjs/js"
	"github.com/oriser/regroup"
)

const regexMove = `(?P<tag>\\move\s*\((?P<x1>\s*-?\d+(?:\.\d+)?),?(?P<y1>\s*-?\d+(?:\.\d+)?),?(?P<x2>\s*-?\d+(?:\.\d+)?),?(?P<y2>\s*-?\d+(?:\.\d+)?),?(?:(?P<t1>\s*-?\d+(?:\.\d+)?),?(?P<t2>\s*-?\d+(?:\.\d+)?))?\))`

type TagMove struct {
	baseTag
	X1 float64 `regroup:"x1,required" js:"x1"`
	Y1 float64 `regroup:"y1,required" js:"y1"`
	X2 float64 `regroup:"x2,required" js:"x2"`
	Y2 float64 `regroup:"y2,required" js:"y2"`
	T1 float64 `regroup:"t1" js:"t1"`
	T2 float64 `regroup:"t2" js:"t2"`
}

func NewTagMove() *TagMove {
	t := &TagMove{}
	t.Object = NewJSObject()

	if t.Object == nil {
		return t
	}

	t.X1 = 0
	t.Y1 = 0
	t.X2 = 0
	t.Y2 = 0
	t.T1 = 0
	t.T2 = 0
	t.jsName = t.Name
	t.jsFromArgs = t.FromArgs
	t.jsFromString = t.FromString
	t.jsToString = t.ToString

	return t
}

// func (t *TagMove) toJS() *js.Object {
// 	tJS := &TagMove{}
// 	tJS.Object = t.Object
// 	tJS.X1 = t.X1
// 	tJS.Y1 = t.Y1
// 	tJS.X2 = t.X2
// 	tJS.Y2 = t.Y2
// 	tJS.T1 = t.T1
// 	tJS.T2 = t.T2
// 	tJS.jsName = t.Name
// 	tJS.jsFromArgs = t.FromArgs
// 	tJS.jsFromString = t.FromString
// 	tJS.jsToString = t.ToString
// 	return tJS.Object
// }

func NewTagMoveJS() *js.Object {
	t := NewTagMove()
	return t.Object
}

func (t *TagMove) Name() string {
	return string(TAG_MOVE)
}

func (t *TagMove) FromArgs(
	x1 float64,
	y1 float64,
	x2 float64,
	y2 float64,
	t1 float64,
	t2 float64,
) *TagMove {
	t.X1 = x1
	t.Y1 = y1
	t.X2 = x2
	t.Y2 = y2
	t.T1 = t1
	t.T2 = t2
	return t
}

func (t *TagMove) FromString(text string) (*TagMove, error) {
	re := regroup.MustCompile(regexMove)
	err := re.MatchToTarget(text, t)

	if err != nil {
		return nil, formatTagParsingError(t, err)
	}

	return t, nil
}

func (t *TagMove) ToString(text string) string {
	s := fmt.Sprintf(`\%s(%g,%g,%g,%g,%g,%g)`, t.Name(), t.X1, t.Y1, t.X2, t.Y2, t.T1, t.T2)
	return s
}

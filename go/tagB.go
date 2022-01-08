package asu

import (
	"fmt"

	"github.com/gopherjs/gopherjs/js"
	"github.com/oriser/regroup"
)

const regexTagB = `(?P<tag>\\b\s*(?P<arg>-?\d+(?:\.\d+)?))`

type TagB struct {
	baseTag
	Argument float64 `regroup:"arg,required" js:"argument"`
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

func NewTagBJS() *js.Object {
	t := NewTagB()
	return t.Object
}

func (t *TagB) FromArgs(f float64) *TagB {
	t.Argument = f
	return t
}

func (t *TagB) FromString(text string) (*TagB, error) {
	re := regroup.MustCompile(regexTagB)
	err := re.MatchToTarget(text, t)

	if err != nil {
		return nil, formatTagParsingError(t, text, err)
	}

	return t, nil
}

func (t *TagB) ToString() string {
	s := fmt.Sprintf(`\%s%g`, t.Name(), t.Argument)
	return s
}

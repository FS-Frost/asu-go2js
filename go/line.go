package asu

import (
	"fmt"
	"strconv"

	"github.com/gopherjs/gopherjs/js"
)

type Line struct {
	*js.Object
	Type           LineType    `js:"type"`
	Layer          int         `js:"layer"`
	Start          *Time       `js:"start"`
	End            *Time       `js:"end"`
	Style          string      `js:"style"`
	Actor          string      `js:"actor"`
	MarginLeft     int         `js:"marginLeft"`
	MarginRight    int         `js:"marginRight"`
	MarginVertical int         `js:"marginVertical"`
	Effect         string      `js:"effect"`
	Content        string      `js:"content"`
	jsFromString   interface{} `js:"fromString"`
	jsToString     interface{} `js:"toString"`
	jsLength       interface{} `js:"length"`
	jsEquals       interface{} `js:"equals"`
}

func NewLine() *Line {
	l := &Line{}
	l.Object = NewJSObject()
	l.Type = LINE_DIALOGUE
	l.Layer = 0
	l.Start = NewTime()
	l.End = NewTime()
	l.Style = "Default"
	l.Actor = ""
	l.MarginLeft = 0
	l.MarginRight = 0
	l.MarginVertical = 0
	l.Effect = ""
	l.Content = ""

	if l.Object == nil {
		return l
	}

	l.jsFromString = l.FromString
	l.jsToString = l.ToString
	l.jsLength = l.Length
	l.jsEquals = l.Equals

	return l
}

// func (l *Line) toJS() *js.Object {
// 	jsLine := &Line{}
// 	jsLine.Object = l.Object
// 	jsLine.Type = l.Type
// 	jsLine.Layer = l.Layer
// 	jsLine.Start = l.Start
// 	jsLine.End = l.End
// 	jsLine.Style = l.Style
// 	jsLine.Actor = l.Actor
// 	jsLine.MarginLeft = l.MarginLeft
// 	jsLine.MarginRight = l.MarginRight
// 	jsLine.MarginVertical = l.MarginVertical
// 	jsLine.Effect = l.Effect
// 	jsLine.Content = l.Content
// 	jsLine.jsFromString = l.FromString
// 	jsLine.jsToString = l.ToString
// 	return jsLine.Object
// }

func NewLineJS() *js.Object {
	l := NewLine()
	return l.Object
}

func (l *Line) FromString(text string) (*Line, error) {
	var err error
	props := ParseLineProperties(text)
	l.Type = LineType(props[LINE_TYPE])
	if !l.Type.IsValid() {
		return nil, NewError("invalid line type: '%s'", l.Type)
	}

	l.Style = props[LINE_STYLE]
	l.Actor = props[LINE_ACTOR]
	l.Effect = props[LINE_EFFECT]
	l.Content = props[LINE_CONTENT]

	l.Layer, err = strconv.Atoi(props[LINE_LAYER])

	if err != nil {
		return nil, NewError("invalid layer: %v", err)
	}

	l.Start, err = NewTime().FromString(props[LINE_START])

	if err != nil {
		return nil, NewError("invalid start time: %v", err)
	}

	l.End, err = NewTime().FromString(props[LINE_END])

	if err != nil {
		return nil, NewError("invalid end time: %v", err)
	}

	l.MarginLeft, err = strconv.Atoi(props[LINE_MARGIN_LEFT])

	if err != nil {
		return nil, NewError("invalid left margin: %v", err)
	}

	l.MarginRight, err = strconv.Atoi(props[LINE_MARGIN_RIGHT])

	if err != nil {
		return nil, NewError("invalid right margin: %v", err)
	}

	l.MarginVertical, err = strconv.Atoi(props[LINE_MARGIN_VERTICAL])

	if err != nil {
		return nil, NewError("invalid vertical margin: %v", err)
	}

	return l, nil
}

func (l *Line) ToString() string {
	return fmt.Sprintf(
		"%s: %d,%s,%s,%s,%s,%d,%d,%d,%s,%s",
		l.Type,
		l.Layer,
		l.Start.ToString(),
		l.End.ToString(),
		l.Style,
		l.Actor,
		l.MarginLeft,
		l.MarginRight,
		l.MarginVertical,
		l.Effect,
		l.Content,
	)
}

func (l *Line) Length() float64 {
	len := l.End.ToSeconds() - l.Start.ToSeconds()
	return truncate(len, 2)
}

func (l *Line) Equals(l2 *Line) bool {
	return l.ToString() == l2.ToString()
}

func ParseLineProperties(line string) map[LineProperty]string {
	result := ApplyRegex(regexLineProp, line)
	props := make(map[LineProperty]string, len(result))

	for k, v := range result {
		props[LineProperty(k)] = v
	}

	return props
}

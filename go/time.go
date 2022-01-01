package asu

import (
	"fmt"
	"strconv"

	"github.com/gopherjs/gopherjs/js"
)

const (
	MAX_TOTAL_SECONDS = 35999.99
	MAX_SECONDS       = 59.99
	MAX_MINUTES       = 59
	MAX_HOURS         = 9
)

type Time struct {
	*js.Object
	Hours            int         `js:"hours"`
	Minutes          int         `js:"minutes"`
	Seconds          float64     `js:"seconds"`
	jsFromSeconds    interface{} `js:"fromSeconds"`
	jsFromArgs       interface{} `js:"fromArgs"`
	jsFromString     interface{} `js:"fromString"`
	jsToString       interface{} `js:"toString"`
	jsToSeconds      interface{} `js:"toSeconds"`
	jsAdd            interface{} `js:"add"`
	jsSub            interface{} `js:"sub"`
	jsAdjustOverplus interface{} `js:"adjustOverplus"`
	jsEquals         interface{} `js:"equals"`
}

func NewTime() *Time {
	t := &Time{}
	t.Object = NewJSObject()

	if t.Object == nil {
		return t
	}

	t.Hours = 0
	t.Minutes = 0
	t.Seconds = 0
	t.jsAdd = t.Add
	t.jsSub = t.Sub
	t.jsAdjustOverplus = t.adjustOverplus
	t.jsFromArgs = t.FromArgs
	t.jsFromSeconds = t.FromSeconds
	t.jsFromString = t.FromString
	t.jsToSeconds = t.ToSeconds
	t.jsToString = t.ToString
	t.jsEquals = t.Equals

	return t
}

// func (t *Time) toJS() *js.Object {
// 	tJS := &Time{}
// 	tJS.Object = t.Object
// 	tJS.Hours = t.Hours
// 	tJS.Minutes = t.Minutes
// 	tJS.Seconds = t.Seconds
// 	tJS.jsFromSeconds = t.FromSeconds
// 	tJS.jsFromArgs = t.FromArgs
// 	tJS.jsFromString = t.FromString
// 	tJS.jsToString = t.ToString
// 	tJS.jsToSeconds = t.ToSeconds
// 	tJS.jsAdd = t.Add
// 	tJS.jsSub = t.Sub
// 	tJS.jsAdjustOverplus = t.adjustOverplus
// 	return tJS.Object
// }

func NewTimeJS() *js.Object {
	t := NewTime()
	return t.Object
}

func (t *Time) FromSeconds(total float64) *Time {
	h := int(total) / 3600
	total -= float64(h * 3600)

	m := int(total) / 60
	total -= float64(m * 60)

	s := total

	if total > MAX_SECONDS {
		s = MAX_SECONDS
	}

	t.Hours = h
	t.Minutes = m
	t.Seconds = truncate(s, 2)

	return t
}

func (t *Time) FromArgs(hours int, minutes int, seconds float64) *Time {
	t.Hours = hours
	t.Minutes = minutes
	t.Seconds = truncate(seconds, 2)
	return t
}

func (t *Time) FromString(text string) (*Time, error) {
	r := `(?P<h>\d+):(?P<m>[0-9]{1,2}?):(?P<s>[0-9]{1,2}(?:\.[0-9]{1,2})?)`
	result := ApplyRegex(r, text)

	h, err := strconv.Atoi(result["h"])

	if err != nil {
		return nil, NewError("invalid hours: %v", err)
	}

	m, err := strconv.Atoi(result["m"])

	if err != nil {
		return nil, NewError("invalid minutes: %v", err)
	}

	s, err := strconv.ParseFloat(result["s"], 64)

	if err != nil {
		return nil, NewError("invalid seconds: %v", err)
	}

	t.Hours = h
	t.Minutes = m
	t.Seconds = s
	t.adjustOverplus()

	return t, nil
}

func (t *Time) ToString() string {
	t.adjustOverplus()
	return fmt.Sprintf("%d:%02d:%05.2f", t.Hours, t.Minutes, t.Seconds)
}

func (t *Time) ToSeconds() float64 {
	t.adjustOverplus()
	return t.toFloat()
}

func (t *Time) toFloat() float64 {
	f := ((float64(t.Hours) * 3600) + (float64(t.Minutes) * 60) + t.Seconds)
	return f
}

func (t *Time) Add(t2 *Time) *Time {
	t.adjustOverplus()
	secondsT1 := t.ToSeconds()
	secondsT2 := t2.ToSeconds()
	total := secondsT1 + secondsT2
	return NewTime().FromSeconds(total)
}

func (t *Time) Sub(t2 *Time) *Time {
	t.adjustOverplus()
	secondsT1 := t.ToSeconds()
	secondsT2 := t2.ToSeconds()
	total := secondsT1 - secondsT2
	return NewTime().FromSeconds(total)
}

func (t *Time) adjustOverplus() {
	var newTotalSeconds float64
	currentTotalSeconds := t.toFloat()

	if currentTotalSeconds <= MAX_TOTAL_SECONDS {
		newTotalSeconds = currentTotalSeconds
	} else {
		newTotalSeconds = MAX_TOTAL_SECONDS
	}

	t2 := NewTime().FromSeconds(newTotalSeconds)
	t.Hours = t2.Hours
	t.Minutes = t2.Minutes
	t.Seconds = t2.Seconds
}

func (t *Time) Equals(t2 *Time) bool {
	return t.ToString() == t2.ToString()
}

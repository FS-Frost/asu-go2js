package asu_test

import (
	"testing"

	"github.com/FS-Frost/asu"
	"github.com/stretchr/testify/require"
)

func TestParseTimeFromString(t *testing.T) {
	text := "1:23:45.67"

	expectedTime := asu.NewTime()
	expectedTime.Hours = 1
	expectedTime.Minutes = 23
	expectedTime.Seconds = 45.67

	actualTime, err := asu.NewTime().FromString(text)
	// t.Logf("expected: %s", expectedTime.ToString())
	// t.Logf("actual: %s", actualTime.ToString())

	require.Nil(t, err)
	require.Equal(t, expectedTime, actualTime)
}

func TestPrintTime(t *testing.T) {
	time := asu.NewTime()
	time.Hours = 1
	time.Minutes = 23
	time.Seconds = 45.67

	expected := "1:23:45.67"
	actual := time.ToString()

	require.Equal(t, expected, actual)
}

func TestPrintOverplusTime(t *testing.T) {
	time := asu.NewTime()
	time.Hours = 12
	time.Minutes = 23
	time.Seconds = 45.67

	expected := "9:59:59.99"
	actual := time.ToString()

	require.Equal(t, expected, actual)
}

func TestAdjustOverplusSeconds(t *testing.T) {
	time := asu.NewTime()
	time.Hours = 0
	time.Minutes = 0
	time.Seconds = 181

	expected := "0:03:01.00"
	actual := time.ToString()

	require.Equal(t, expected, actual)
}

func TestAdjustOverplusMinutes(t *testing.T) {
	time := asu.NewTime()
	time.Hours = 0
	time.Minutes = 61
	time.Seconds = 181

	expected := "1:04:01.00"
	actual := time.ToString()

	require.Equal(t, expected, actual)
}

func TestAdjustOverplusHours(t *testing.T) {
	time := asu.NewTime()
	time.Hours = 9
	time.Minutes = 61
	time.Seconds = 181

	expected := "9:59:59.99"
	actual := time.ToString()

	require.Equal(t, expected, actual)
}

func TestSub(t *testing.T) {
	t1 := asu.NewTime()
	t1.Hours = 1
	t1.Minutes = 23
	t1.Seconds = 45.67

	t2 := asu.NewTime()
	t2.Hours = 1
	t2.Minutes = 19
	t2.Seconds = 43.66

	expected := asu.NewTime()
	expected.Hours = 0
	expected.Minutes = 4
	expected.Seconds = 2.01

	actual := t1.Sub(t2)
	// t.Logf("expected: %s", expected.ToString())
	// t.Logf("actual: %s", actual.ToString())

	require.Equal(t, expected, actual)
}

package asu_test

import (
	"testing"

	"github.com/FS-Frost/asu"
	"github.com/stretchr/testify/require"
)

func TestLineContentAssignment(t *testing.T) {
	content := "some-text"
	l := asu.NewLine()
	l.Content = content
	require.Equal(t, content, l.Content)
}

func TestNewLineFromString(t *testing.T) {
	text := `Dialogue: 3,0:00:00.00,1:23:45.67,Default,Chitanda,1,2,3,Some fx,{\b1\fs32\1c&H2F3066&}Some {\1c&H0B0B26&}text`
	expected := asu.NewLine()
	expected.Type = asu.LINE_DIALOGUE
	expected.Layer = 3
	expected.Start = asu.NewTime()
	expected.End = asu.NewTime().FromArgs(1, 23, 45.67)
	expected.Style = "Default"
	expected.Actor = "Chitanda"
	expected.MarginLeft = 1
	expected.MarginRight = 2
	expected.MarginVertical = 3
	expected.Effect = "Some fx"
	expected.Content = `{\b1\fs32\1c&H2F3066&}Some {\1c&H0B0B26&}text`
	actual, err := asu.NewLine().FromString(text)
	require.Nil(t, err)
	require.Equal(t, expected, actual)
}

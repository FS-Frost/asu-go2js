package asu_test

import (
	"fmt"
	"testing"

	"github.com/FS-Frost/asu"
	"github.com/stretchr/testify/require"
)

// func TestParseValidTagB(t *testing.T) {
// 	text := rawNewLine + `{\b1}Some text`
// 	exists := asu.TagExists(text, asu.TAG_B)
// 	require.True(t, exists)
// }

// func TestParseInvalidTagB(t *testing.T) {
// 	text := rawNewLine + `{\be1}Some text`
// 	exists := asu.TagExists(text, asu.TAG_B)
// 	require.False(t, exists)
// }

func TestFilterTagB(t *testing.T) {
	text := rawNewLine + `{\b1}Some text`
	expected := asu.NewTagB().FromArgs(1)
	actual, err := asu.NewTagB().FromString(text)
	require.Nil(t, err)
	require.Equal(t, expected, actual)
}

func TestReplaceTag(t *testing.T) {
	cases := []struct {
		tagType asu.TagType
		oldTag  string
		newTag  string
	}{
		{
			tagType: asu.TAG_B,
			oldTag:  `\b1`,
			newTag:  `\pos(10,20)`,
		},
		{
			tagType: asu.TAG_POS,
			oldTag:  `\pos(10,20)`,
			newTag:  `\move(30,40,50,60)`,
		},
	}

	for _, c := range cases {
		text := fmt.Sprintf(`{%s}Some text %s{%s}Other text`, c.oldTag, c.oldTag, c.oldTag)
		expected := fmt.Sprintf(`{%s}Some text %s{%s}Other text`, c.newTag, c.oldTag, c.oldTag)
		var actual string

		require.NotPanics(t, func() {
			actual = asu.ReplaceTag(text, c.tagType, c.newTag)
		})

		require.Equal(t, expected, actual, "replaced tag mismatch")
	}
}

// func TestAlgo(t *testing.T) {
// 	text := `{\\b1}Some text \\b1{\\b1}Other text`
// 	expected := `{\\pos(10,20)}Some text \\b1{\\b1}Other text`
// 	actual := asu.ReplaceTag(text, asu.TAG_B, )
// }

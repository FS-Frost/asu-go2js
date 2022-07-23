package asu_test

import (
	"fmt"
	"testing"

	"github.com/FS-Frost/asu"
	"github.com/stretchr/testify/require"
)

var (
	floatBaseCases = []float64{0, 1, -1, 1.00001, -1.00001, 123456789, -123456789, 123456789.00001, -123456789.0000123456789}
)

func TestBTag(t *testing.T) {
	var n float64 = 12345678921234234234243
	text := fmt.Sprintf(`\b%f`, n)
	expected := &asu.TagB{Argument: n}
	actual, err := asu.NewTagB().FromString(text)
	require.Nil(t, err, text)
	require.Equal(t, expected, actual, text)
}

func TestNewTagBFromString(t *testing.T) {
	for i, f := range floatBaseCases {
		caseName := fmt.Sprintf("case %d: %f", i, f)
		text := fmt.Sprintf(`\b%.9f`, f)
		expected := asu.NewTagB().FromArgs(f)

		actual, err := asu.NewTagB().FromString(text)
		require.Nil(t, err, caseName)
		require.Equal(t, expected, actual, caseName)
	}
}

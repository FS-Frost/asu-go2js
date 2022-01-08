package asu_test

import (
	"fmt"
	"testing"

	"github.com/FS-Frost/asu"
	"github.com/stretchr/testify/require"
)

func TestNewTagBFromString(t *testing.T) {
	type caseType struct {
		caseName string
		text     string
		expected *asu.TagB
		f        float64
	}

	cases := []caseType{}
	floatValues := []float64{}
	baseValues := []float64{0, 10}

	for _, n := range baseValues {
		f := n + 0.1000001
		floatValues = append(floatValues, n, -n, f, -f)
	}

	for _, f := range floatValues {
		newCase := caseType{
			caseName: fmt.Sprintf("%g", f),
			text:     fmt.Sprintf(`\b%g`, f),
			expected: asu.NewTagB().FromArgs(f),
			f:        f,
		}

		cases = append(cases, newCase)
	}

	for _, c := range cases {
		actual, err := asu.NewTagB().FromString(c.text)
		require.Nil(t, err, c.caseName)
		require.Equal(t, c.expected, actual, c.caseName)
	}
}

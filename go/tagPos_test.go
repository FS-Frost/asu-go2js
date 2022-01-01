package asu_test

import (
	"fmt"
	"testing"

	"github.com/FS-Frost/asu"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewTagPosFromString(t *testing.T) {
	type caseType struct {
		caseName string
		text     string
		expected *asu.TagPos
		x        float64
		y        float64
	}

	cases := []caseType{}
	floatValues := []float64{}
	baseValues := []float64{0, 10}

	for _, n := range baseValues {
		f := n + 0.1000001
		floatValues = append(floatValues, n, -n, f, -f)
	}

	for _, x := range floatValues {
		for _, y := range floatValues {
			newCase := caseType{
				caseName: fmt.Sprintf("%g, %g", x, y),
				text:     fmt.Sprintf(`\pos(%g,%g)`, x, y),
				expected: asu.NewTagPos().FromArgs(x, y),
				x:        x,
				y:        y,
			}

			cases = append(cases, newCase)
		}
	}

	for _, c := range cases {
		actual, err := asu.NewTagPos().FromString(c.text)
		require.Nil(t, err, c.caseName)
		pass := assert.Equal(t, c.expected, actual, c.caseName)

		if pass {
			continue
		}

		t.Log()
	}
}

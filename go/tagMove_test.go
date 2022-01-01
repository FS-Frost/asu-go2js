package asu_test

import (
	"fmt"
	"testing"

	"github.com/FS-Frost/asu"
	"github.com/stretchr/testify/require"
)

func TestNewTagMoveFromString(t *testing.T) {
	type testCase struct {
		text     string
		expected *asu.TagMove
		x1       float64
		y1       float64
		x2       float64
		y2       float64
		t1       float64
		t2       float64
	}

	cases := []testCase{}
	floatValues := []float64{}
	baseValues := []float64{1}

	for _, n := range baseValues {
		f := n + 0.1000001
		floatValues = append(floatValues, n, -n, f, -f)
	}

	for _, x1 := range floatValues {
		for _, y1 := range floatValues {
			for _, x2 := range floatValues {
				for _, y2 := range floatValues {
					for _, t1 := range floatValues {
						for _, t2 := range floatValues {
							newCase := testCase{
								text:     fmt.Sprintf(`\move(%g,%g,%g,%g,%g,%g)`, x1, y1, x2, y2, t1, t2),
								expected: asu.NewTagMove().FromArgs(x1, y1, x2, y2, t1, t2),
								x1:       x1,
								y1:       y1,
								x2:       x2,
								y2:       y2,
								t1:       t1,
								t2:       t2,
							}

							cases = append(cases, newCase)
						}
					}
				}
			}
		}
	}

	for _, c := range cases {
		actual, err := asu.NewTagMove().FromString(c.text)
		require.Nil(t, err, c.text)
		require.Equal(t, c.expected, actual, c.text)

		// if pass {
		// 	continue
		// }

		// t.Log()
	}
}

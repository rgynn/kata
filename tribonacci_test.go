package kata

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTribonacci(t *testing.T) {

	type testcase struct {
		Signature [3]float64
		N         int
		Expected  []float64
	}

	testcases := []testcase{
		{[3]float64{1, 1, 1}, 10, []float64{1, 1, 1, 3, 5, 9, 17, 31, 57, 105}},
		{[3]float64{0, 0, 1}, 10, []float64{0, 0, 1, 1, 2, 4, 7, 13, 24, 44}},
		{[3]float64{0, 1, 1}, 10, []float64{0, 1, 1, 2, 4, 7, 13, 24, 44, 81}},
		{[3]float64{1, 0, 0}, 10, []float64{1, 0, 0, 1, 1, 2, 4, 7, 13, 24}},
		{[3]float64{0, 0, 0}, 10, []float64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
		{[3]float64{1, 2, 3}, 10, []float64{1, 2, 3, 6, 11, 20, 37, 68, 125, 230}},
		{[3]float64{3, 2, 1}, 10, []float64{3, 2, 1, 6, 9, 16, 31, 56, 103, 190}},
		{[3]float64{1, 1, 1}, 1, []float64{1}},
		{[3]float64{300, 200, 100}, 0, []float64{}},
		{[3]float64{0.5, 0.5, 0.5}, 30, []float64{0.5, 0.5, 0.5, 1.5, 2.5, 4.5, 8.5, 15.5, 28.5, 52.5, 96.5, 177.5, 326.5, 600.5, 1104.5, 2031.5, 3736.5, 6872.5, 12640.5, 23249.5, 42762.5, 78652.5, 144664.5, 266079.5, 489396.5, 900140.5, 1655616.5, 3045153.5, 5600910.5, 10301680.5}},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprintf("%d", i+1), func(t *testing.T) {
			if result := Tribonacci(tc.Signature, tc.N); !reflect.DeepEqual(result, tc.Expected) {
				t.Fatalf("expected: %v, got: %v", tc.Expected, result)
			}
		})
	}
}

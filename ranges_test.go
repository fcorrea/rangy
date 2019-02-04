package rangedown

import (
	"reflect"
	"testing"
)

var rangetests = []struct {
	testName string
	in1      int64
	in2      int
	out      map[int][]int64
}{
	{"Get ranges", 80, 2, map[int][]int64{0: {0, 40}, 1: {41, 80}}},
	{"Get ranges grows the last chunk", 83, 2, map[int][]int64{0: {0, 41}, 1: {42, 83}}},
	{"Get ranges one chunk", 80, 1, map[int][]int64{0: {0, 80}}},
}

func TestRanges(t *testing.T) {
	for _, tt := range rangetests {
		t.Run(tt.testName, func(t *testing.T) {
			s := GetRanges(tt.in1, tt.in2)
			if !reflect.DeepEqual(s, tt.out) {
				t.Errorf("got %q, want %q", s, tt.out)
			}
		})
	}
}

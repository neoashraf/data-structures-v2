package rotation

import (
	"reflect"
	"testing"
)

func TestSimpleRotation(t *testing.T) {
	testCases := map[string]struct {
		input    []int
		n        int
		expected []int
	}{
		"simple_test": {
			input:    []int{1, 2, 3, 4, 5},
			n:        2,
			expected: []int{3, 4, 5, 1, 2},
		},
	}

	for key, testCase := range testCases {
		t.Run(key, func(t *testing.T) {
			got := SimpleRotation(testCase.input, testCase.n)
			if !reflect.DeepEqual(got, testCase.expected) {
				t.Errorf("Failed at %v\nexpected=%v\ngot=%v", key, testCase.expected, got)
			}
		})
	}
}

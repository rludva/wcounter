package counter

import (
	"reflect"
	"testing"
)

func TestWcounter(t *testing.T) {
	type testCase struct {
		lines    string
		expected map[string]int
	}
	testCases := []testCase{
		{
			"",
			map[string]int{},
		},
		{
			"lorem",
			map[string]int{
				"lorem": 1,
			},
		},
		{
			"lorem lorem",
			map[string]int{
				"lorem": 2,
			},
		},
		{
			"lorem lorem\n",
			map[string]int{
				"lorem": 2,
			},
		},
		{
			"lorem\tlorem lorem\n",
			map[string]int{
				"lorem": 3,
			},
		},
		{
			"lorem, lorem lorem\n",
			map[string]int{
				"lorem": 3,
			},
		},
		{
			"lorem. lorem lorem\n",
			map[string]int{
				"lorem": 3,
			},
		},
	}

	for _, tc := range testCases {
		if report := Wcounter(tc.lines); !reflect.DeepEqual(report, tc.expected) {
			t.Errorf("Wcounter(%q):%v expected %v", tc.lines, report, tc.expected)
		}
	}
}

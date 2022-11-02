package q02

import "testing"

type TestCase struct {
	inputA        string
	inputB        string
	isPermutation bool
}

var testCases = []TestCase{
	{inputA: "abc", inputB: "cba", isPermutation: true},
	{inputA: "abc", inputB: "def", isPermutation: false},
	{inputA: "abc", inputB: "bc", isPermutation: false},
}

func TestCheckPermutation(t *testing.T) {
	for _, tc := range testCases {
		if CheckPermutation(tc.inputA, tc.inputB) != tc.isPermutation {
			t.Errorf("'%s' and '%s': %s", tc.inputA, tc.inputB, expectedMsg(tc.isPermutation))
		}
	}
}

func TestCheckPermutation2(t *testing.T) {
	for _, tc := range testCases {
		if CheckPermutation2(tc.inputA, tc.inputB) != tc.isPermutation {
			t.Errorf("'%s' and '%s': %s", tc.inputA, tc.inputB, expectedMsg(tc.isPermutation))
		}
	}
}

// Helpers
func expectedMsg(isPermutation bool) string {
	if isPermutation {
		return "should be permutations"
	} else {
		return "should not be permutations"
	}
}

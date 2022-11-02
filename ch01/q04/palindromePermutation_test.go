package q04

import "testing"

type TestCase struct {
	input    string
	expected bool
}

var testCases = []TestCase{
	{input: "tact coa", expected: true},
	{input: "Tact coa", expected: true},
	{input: "abc", expected: false},
}

func TestIsPalindromePermutation(t *testing.T) {

	for _, tc := range testCases {
		if output := IsPalindromePermutation(tc.input); output != tc.expected {
			t.Errorf("'%s' %s", tc.input, expectedMsg(tc.expected))
		}
	}
}

func expectedMsg(expected bool) string {
	if expected {
		return "should be a palindrome permutation"
	} else {
		return "should not be a palindrome permutation"
	}
}

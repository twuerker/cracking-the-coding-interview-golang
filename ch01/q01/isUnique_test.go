package q01

import "testing"

type TestCase struct {
	input    string
	isUnique bool
}

var testCases = []TestCase{
	{input: "abc", isUnique: true},
	{input: "aaa", isUnique: false},
	{input: "abcdd", isUnique: false},
}

func TestIsUnique(t *testing.T) {
	for _, tc := range testCases {
		if IsUnique(tc.input) != tc.isUnique {
			t.Errorf("'%s' %s", tc.input, uniqueMsg(tc.isUnique))
		}
	}
}

func TestIsUnique2(t *testing.T) {
	for _, tc := range testCases {
		if IsUnique2(tc.input) != tc.isUnique {
			t.Errorf("'%s' %s", tc.input, uniqueMsg(tc.isUnique))
		}
	}
}

func TestIsUnique3(t *testing.T) {
	for _, tc := range testCases {
		if IsUnique3(tc.input) != tc.isUnique {
			t.Errorf("'%s' %s", tc.input, uniqueMsg(tc.isUnique))
		}
	}
}

func TestIsUnique4(t *testing.T) {
	for _, tc := range testCases {
		if IsUnique4(tc.input) != tc.isUnique {
			t.Errorf("'%s' %s", tc.input, uniqueMsg(tc.isUnique))
		}
	}

	//Special cases for this implementation

	if input := "abcdefghijklmnopqrstuvwxyz"; !IsUnique4(input) {
		t.Errorf("'%s' %s", input, uniqueMsg(true))
	}

	if input := "abc0101"; !IsUnique4(input) {
		t.Errorf("'%s' %s", input, uniqueMsg(true))
	}

	if input := "abc!!"; !IsUnique4(input) {
		t.Errorf("'%s' %s", input, uniqueMsg(true))
	}

	if input := "Aa"; IsUnique4(input) {
		t.Errorf("'%s' %s", input, uniqueMsg(false))
	}

}

//Helpers

func uniqueMsg(shouldBeUnique bool) string {
	if shouldBeUnique {
		return "should be unique"
	} else {
		return "should not be unique"
	}
}

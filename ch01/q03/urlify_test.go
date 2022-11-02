package q03

import "testing"

type TestCase struct {
	inputString string
	trueLength  int
	expected    string
}

var testCases = []TestCase{
	{expected: "Mr%20John%20Smith", trueLength: 13, inputString: "Mr John Smith    "},
	{expected: "Mr%20John%20Smith", trueLength: 13, inputString: "Mr John Smith"},
	{expected: "%20%20", trueLength: 2, inputString: "  "},
}

func TestURLify(t *testing.T) {
	for _, tc := range testCases {
		if output := URLify(tc.inputString, tc.trueLength); output != tc.expected {
			t.Errorf("'%s' did not equal expected '%s'", output, tc.expected)
		}
	}
}

func TestURLify2(t *testing.T) {
	for _, tc := range testCases {
		if output := URLify2(tc.inputString, tc.trueLength); output != tc.expected {
			t.Errorf("'%s' did not equal expected '%s'", output, tc.expected)
		}
	}
}

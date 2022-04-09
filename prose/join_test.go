package prose

import (
	"testing"
)

type testData struct {
	list     []string
	expected string
}

func TestJoinWithCommas(t *testing.T) {
	tests := []testData{
		{list: []string{"apple"}, expected: "apple"},
		{list: []string{"apple", "orange"}, expected: "apple and orange"},
		{list: []string{"apple", "orange", "pear"}, expected: "apple, orange, and pear"},
	}
	for _, testCase := range tests {
		output := JoinWithCommas(testCase.list)
		if output != testCase.expected {
			t.Errorf("JoinWithCommas(%#v) = \"%s\", expected \"%s\"", testCase.list, output, testCase.expected)
		}
	}
}

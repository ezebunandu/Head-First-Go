package prose

import "testing"

func TestTwoElements(t *testing.T) {
	list := []string{"apple", "orange"}
	expected := "apple and orange"
	output := JoinWithCommas(list)
	if output != expected {
		t.Errorf("JoinWithCommas(%#v) = \"%s\", expected \"%s\"", list, output, expected)
	}
}

func TestThreeElements(t *testing.T) {
	list := []string{"apple", "orange", "pear"}
	expected := "apple, orange, and pear"
	output := JoinWithCommas(list)
	if output != expected {
		t.Errorf("JoinWithCommas(%#v) = \"%s\", expected \"%s\"", list, output, expected)
	}
}

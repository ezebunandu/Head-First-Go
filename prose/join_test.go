package prose

import (
	"fmt"
	"testing"
)

func TestTwoElements(t *testing.T) {
	list := []string{"apple", "orange"}
	expected := "apple and orange"
	output := JoinWithCommas(list)
	if output != expected {
		t.Errorf(errorString(list, output, expected))
	}
}

func TestThreeElements(t *testing.T) {
	list := []string{"apple", "orange", "pear"}
	expected := "apple, orange, and pear"
	output := JoinWithCommas(list)
	if output != expected {
		t.Errorf(errorString(list, output, expected))
	}
}

func errorString(list []string, output string, expected string) string {
	return fmt.Sprintf("JoinWithCommas(%#v) = \"%s\", expected, \"%s\"", list, output, expected)
}

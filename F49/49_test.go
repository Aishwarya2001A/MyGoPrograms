package mathops

import "testing"

func TestDifference1(test *testing.T) {

	actual := Difference(4, 6)
	expected := 2
	if actual != expected {
		test.Errorf("actual value: %d, expected value: %d", actual, expected)
	} else {
		test.Logf("actual value: %d, expected value: %d", actual, expected)
	}

}

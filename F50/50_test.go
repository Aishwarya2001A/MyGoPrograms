package mathops

import "testing"

type diffTest struct {
	name                 string
	num1, num2, expected int
}

var diffTests = []diffTest{
	{"1st Test: values 10,3 ", 9, 4, 5},
	{"2nd Test: values 4,8 ", 4, 1, 30},
	{"3rd Test: values 6,9 ", 6, 1, 5},
	{"4th Test: values 10,13 ", 10, 1, 90},
}

func TestDifference(t *testing.T) {
	for _, test := range diffTests {
		t.Run(test.name, func(t *testing.T) {

			actual := Difference(test.num1, test.num2)

			if actual != test.expected {
				t.Errorf(test.name, " FAIL: actual %d not equal to expected %d", actual, test.expected)
			} else {
				t.Logf(test.name, " PASS: actual %d equal to expected %d", actual, test.expected)
			}
		})

	}
}

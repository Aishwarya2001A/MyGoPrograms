package mathops

func Difference(num1, num2 int) (difference int) {

	if num1 > num2 {
		return num1 - num2
	} else {
		return num2 - num1
	}
}

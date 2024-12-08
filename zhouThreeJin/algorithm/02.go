package IsUGuly

func isUgly(n int) bool {
	if n <= 0 {
		return false
	}
	number := n
	for number%5 == 0 {
		number = number / 5
	}
	for number%3 == 0 {
		number = number / 3
	}
	for number%2 == 0 {
		number = number / 2
	}
	if number == 1 {
		return true
	} else {
		return false
	}

}

package main

// isPowerOfTwo returns true if the given number is of two
// else it returns false

func isPowerOfTwo(num int) bool {

	if num == 0 {
		return false
	}

	for num != 1 {
		if num%2 != 0 {
			return false
		}
		num /= 2
	}

	return true
}

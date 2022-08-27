package hello

func FacItr(n int) int {
	result := 1
	for n > 0 {
		result *= n
		n--
	}
	return result
}

func FacItr2(n int) int {
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}

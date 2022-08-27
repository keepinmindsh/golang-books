package problems

func Fibonacci(loopCount int) {

	firstValue := 0
	secondValue := 1

	println(firstValue)
	println(secondValue)

	for i := 0; i < loopCount; i++ {
		firstValue, secondValue = secondValue, firstValue+secondValue

		println(secondValue)
	}

}

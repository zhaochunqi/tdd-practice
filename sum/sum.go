package sum

func Sum(numbers []int) (sum int) {

	for _, number := range numbers {
		sum += number
	}

	return
}

func SumAll(numbersToSum ...[]int) (sums []int) {

	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return
}

func SumAllTails(numbersToSumTails ...[]int) (sums []int) {
	for _, numbers := range numbersToSumTails {

		var tail []int

		if len(numbers) < 2 {
			tail = []int{0}
		} else {
			tail = numbers[1:]
		}

		sums = append(sums, Sum(tail))
	}
	return
}

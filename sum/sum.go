package sum

func Sum(nubmers [5]int) (sum int) {
	for _, number := range nubmers {
		sum += number
	}
	return
}

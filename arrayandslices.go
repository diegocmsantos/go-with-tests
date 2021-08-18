package arrayandslices

func Sum(numbers []int) int {
	total := 0
	for _, n := range numbers {
		total += n
	}
	return total
}

func SumAll(numsToSum ... []int) []int {
	lenOfArrays := len(numsToSum)
	result := make([]int, lenOfArrays)
	for i, numbers := range numsToSum {
		result[i] = Sum(numbers)
	}
	return result
}

func SumAllTails(numsToSum ... []int) []int {
	lenOfArrays := len(numsToSum)
	result := make([]int, lenOfArrays)
	for i, numbers := range numsToSum {
		if len(numbers) == 0 {
			continue
		}
		result[i] = Sum(numbers[1:])
	}
	return result
}
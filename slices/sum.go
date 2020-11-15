package slices

// Sum adds array elements
func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

// SumAll elements for each array
func SumAll(numbersToSum ...[]int) []int {

	var sums []int
	//:= make([]int, lengthOfNumbers) // create a slice with this size

	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

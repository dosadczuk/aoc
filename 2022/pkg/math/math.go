package math

// MaxInt returns the maximum value in given set of values.
func MaxInt(nums ...int) int {
	max := nums[0]

	for i := 1; i < len(nums); i++ {
		num := nums[i]
		if num > max {
			max = num
		}
	}

	return max
}

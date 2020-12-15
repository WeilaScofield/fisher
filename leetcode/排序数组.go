package leetcode

func sortArray(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}

	l, r, p := 0, len(nums)-1, nums[0]

	for i := 1; l < r; {
		if nums[i] < p {
			nums[l], nums[i] = nums[i], nums[l]
			l++
			i++
		} else {
			nums[i], nums[r] = nums[r], nums[i]
			r--
		}
	}

	sortArray(nums[:l])
	sortArray(nums[l+1:])
	return nums
}

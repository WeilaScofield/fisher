/*给定一个大小为 n 的数组，找到其中的多数元素。多数元素是指在数组中出现次数大于 ⌊ n/2 ⌋ 的元素。

你可以假设数组是非空的，并且给定的数组总是存在多数元素。*/
package leetcode

func MajorityElement(nums []int) int {
	res := 0
	set := make(map[int]int)
	limit := len(nums) / 2
	for _, v := range nums {
		set[v] += 1
	}
	for k, v := range set {
		if v > limit {
			res = k
		}
	}
	return res
}

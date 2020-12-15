/*给定一个大小为 n 的数组，找出其中所有出现超过 ⌊ n/3 ⌋ 次的元素。

说明: 要求算法的时间复杂度为 O(n)，空间复杂度为 O(1)。*/
package leetcode

func MajorityElements(nums []int) []int {
	var res []int
	set := make(map[int]int)
	limit := len(nums) / 3
	for _, v := range nums {
		set[v] += 1
	}
	for k, v := range set {
		if v > limit {
			temp := k
			res = append(res, temp)
		}
	}
	return res
}

/*输入整数数组 arr ，找出其中最小的 k 个数。例如，
输入4、5、1、6、2、7、3、8这8个数字，则最小的4个数字是1、2、3、4。*/

package leetcode

func GetLeastNumbers(arr []int, k int) []int {
	if len(arr) < k || k < 1 {
		return nil
	}

	var res []int
	for i := 0; i < k; i++ {
		selectV := arr[i]
		selectI := i
		for l := i + 1; l < len(arr); l++ {
			if arr[l] < selectV {
				selectV = arr[l]
				selectI = l
			}
		}
		arr[i], arr[selectI] = arr[selectI], arr[i]
		res = append(res, arr[i])
	}
	return res
}

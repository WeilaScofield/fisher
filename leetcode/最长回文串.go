/*给定一个包含大写字母和小写字母的字符串，找到通过这些字母构造成的最长的回文串。
在构造过程中，请注意区分大小写。比如 "Aa" 不能当做一个回文字符串。*/

package leetcode

func longestPalindrome(s string) int {
	if s == "" {
		return 0
	}
	counter := make(map[rune]int)
	for _, v := range s {
		counter[v] += 1
	}

	odd := 0
	sum := 0
	for _, v := range counter {
		if isOdd(v) {
			odd = 1
			sum += v - 1
		} else {
			sum += v
		}
	}

	return sum + odd
}

func isOdd(v int) bool {
	if v%2 == 0 {
		return false
	}
	return true
}

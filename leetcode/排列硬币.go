/*你总共有n枚硬币，你需要将它们摆成一个阶梯形状，第n行就必须正好有n枚硬币。

给定一个数字n，找出可形成完整阶梯行的总行数。

n是一个非负整数，并且在32位有符号整型的范围内。*/
package leetcode

func arrangeCoins(n int) int {
	if n == 0 {
		return 0
	}

	i := 1
	l := 1
	for {
		if l == n {
			return i
		}

		if l < n {
			i++
			l = l + i
		}

		if l > n {
			return i - 1
		}

	}

}

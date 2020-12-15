/*给定一个包含了一些 0 和 1的非空二维数组 grid , 一个 岛屿 是由四个方向 (水平或垂直) 的 1 (代表土地) 构成的组合。
你可以假设二维矩阵的四个边缘都被水包围着。

找到给定的二维数组中最大的岛屿面积。(如果没有岛屿，则返回面积为0。)*/

package leetcode

func MaxAreaOfIsland(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	if m == 0 || n == 0 {
		return 0
	}

	maxIslands := 0

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				tempIslands := 0
				dfs(grid, i, j, &tempIslands)
				maxIslands = max(maxIslands, tempIslands)
			}
		}
	}

	return maxIslands
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func dfs(grid [][]int, x, y int, maxIslands *int) {
	if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[0]) || grid[x][y] != 1 {
		return
	}

	*maxIslands += 1
	grid[x][y] = 0

	dfs(grid, x+1, y, maxIslands)
	dfs(grid, x-1, y, maxIslands)
	dfs(grid, x, y+1, maxIslands)
	dfs(grid, x, y-1, maxIslands)
}

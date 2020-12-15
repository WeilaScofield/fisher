package main

import "fmt"

func createMaze(maze *[7][7]int) {
	for i := 0; i < 7; i++ {
		maze[0][i] = 1
		maze[6][i] = 1
	}

	for i := 0; i < 7; i++ {
		maze[i][0] = 1
		maze[i][6] = 1
	}

	maze[3][1] = 1
	maze[3][2] = 1
	//maze[1][2] = 1
	//maze[2][2] = 1
}

func displayMaze(maze [7][7]int) {
	for i := 0; i < 7; i++ {
		for j := 0; j < 7; j++ {
			fmt.Print(maze[i][j], " ")
		}
		fmt.Println()
	}
	fmt.Println()
}

func findPath(maze *[7][7]int, i, j int, counter *int) bool {

	// recursive exit 1
	if maze[5][5] == 2 {
		return true
	} else {
		if maze[i][j] == 0 || maze[i][j] == 2 {
			maze[i][j] = 2
			*counter++
			//walking path:down,right,up,left
			if findPath(maze, i+1, j, counter) {
				return true
			} else if findPath(maze, i, j+1, counter) {
				return true
			} else if findPath(maze, i-1, j, counter) {
				return true
			} else if findPath(maze, i, j-1, counter) {
				return true
			} else {
				maze[i][j] = 3
				return false //go noway
			}
		} else {
			return false //meet obstacle
		}
	}

}

func main() {
	//stipulate
	//reachable = 0
	//obstacle  = 1
	//passage   = 2
	//impasses  = 3
	var maze [7][7]int
	var counter int

	createMaze(&maze)

	displayMaze(maze)

	findPath(&maze, 1, 1, &counter)

	displayMaze(maze)
	fmt.Println("the length of path:", counter)

}

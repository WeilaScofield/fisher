package leetcode

func Robot(command string, obstacles [][]int, x int, y int) bool {
	obmap := make(map[[2]int]bool)
	for _, v := range obstacles {
		obmap[[2]int{v[0], v[1]}] = true
	}

	if !isReached([]int{0, 0}, command, x, y) {
		return false
	}

	for ob, _ := range obmap {
		if ob[0] > x || ob[1] > y {
			continue
		}

		if isReached([]int{0, 0}, command, ob[0], ob[1]) {
			return false
		}
	}
	return true
}

func isReached(o []int, command string, x, y int) bool {
	for {
		for _, v := range command {
			switch string(v) {
			case "U":
				o[1]++
			case "R":
				o[0]++
			}

			if o[0] > x || o[1] > y {
				return false
			}

			if o[0] == x && o[1] == y {
				return true
			}

		}
	}
}

package main

import (
	"fmt"
	"os"
)

// chessMap store and recover

// define value node
type ValNode struct {
	row int
	col int
	val int
}

func main() {
	// create two-dimensional array
	var chessArray [11][11]int
	chessArray[2][3] = 7
	chessArray[4][6] = 1
	fmt.Println("raw array--")
	for _, v1 := range chessArray {
		for _, v2 := range v1 {
			fmt.Printf("%d\t", v2)
		}
		fmt.Println()
	}

	// transfer raw array into sparseArray
	var sparseArray []ValNode
	valNode := ValNode{
		row: 11,
		col: 11,
		val: 0,
	}

	// record the statement of raw array
	sparseArray = append(sparseArray, valNode)
	// generate sparseArray
	for i, v1 := range chessArray {
		for j, v2 := range v1 {
			if v2 != 0 {
				valNode = ValNode{
					row: i,
					col: j,
					val: v2,
				}
				sparseArray = append(sparseArray, valNode)
			}
		}
	}

	var chessMap string
	fmt.Println("sparseArray--")
	for i, v := range sparseArray {
		fmt.Printf("%d:%d,%d,%d\n", i, v.row, v.col, v.val)
		s := fmt.Sprintf("%d,%d,%d\n", v.row, v.col, v.val)
		chessMap = chessMap + s
	}

	// store the chessMap
	filepath := "./稀疏数组/chessMap.txt"
	file, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	n, err := file.WriteString(chessMap)
	if err != nil {
		fmt.Println(err)
	}

	// recover the chessMap
	file, err = os.Open(filepath)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	var bufDate []byte
	bufDate = make([]byte, n)

	n, err = file.Read(bufDate)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("bufData--")
	fmt.Println(string(bufDate))

}

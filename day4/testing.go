package main

import (
	"fmt"
	"strconv"
)

func replaceBingoNumber(board [][]string, bingonumb int) {
	bingonumbstr := strconv.Itoa(bingonumb)

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board); j++ {
			if board[i][j] == bingonumbstr {
				fmt.Println("Found")
				board[i][j] = "-1"
			}
		}
	}
}

func main() {
	values := [][]string{{"1", "2", "3", "4", "5"}, {"1", "2", "3", "4", "5"}, {"1", "2", "3", "4", "5"}}
	replaceBingoNumber(values[:][:], 2)
	fmt.Println(values)
}

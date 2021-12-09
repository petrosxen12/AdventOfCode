package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var fileName string
	if os.Getenv("DEBUG") == "1" {
		fmt.Println("In Debug Mode")
		fileName = "test_input.txt"
	} else {
		fileName = "input.txt"
	}

	file, err := os.Open(fileName)
	defer file.Close()

	if err != nil {
		fmt.Printf("Error, can't open file %s", err)
	}
	scanner := bufio.NewScanner(file)

	var bingo_numbs string
	//Get bingo numbers
	for scanner.Scan() {
		bingo_numbs = scanner.Text()
		fmt.Println(bingo_numbs)
		break
	}

	//To read boards
	fmt.Println("Boards")

	var boards = make(map[int][][]string)
	var number_board = 1

	for scanner.Scan() {
		sl := boards[number_board][:]
		for i := 0; i < 5; i++ {

			scanner.Scan()
			board := strings.Fields(scanner.Text())
			// fmt.Printf("%T\n", board)
			// fmt.Printf("%v\n", board)
			//Slice that appends to board slide
			sl = append(sl, board)
			boards[number_board] = sl
		}
		fmt.Println(boards[number_board])
		number_board++
		fmt.Println("End of board")
	}

	//For debugging purposes
	// for i, v := range boards {
	// 	fmt.Printf("Board number: %d\nContent: %v\n", i, v)
	// 	fmt.Println("Testing single value: ", v[0][1])
	// }
	var bingonumbs chan int
	var winners chan wins
	
	for 


}

type wins struct {
	boardnumber    int
	sum            int
	winning_number int
}

//Winning board
/*
1. Start by finding the sum of all unmarked numbers on that board;
2. Multiply that sum by the number that was just called when the board won, 24, to get the final score, 188 * 24 = 4512.
*/

/*
Logic:
1. Bingo numbers will be send to `bingonumbs` channel
2. Each board goroutine will consume it and change the value
3. Each board goroutine will check rows and columns after every change
4. Send back the boardnumber of winning along with sum of non-winning values*(winning number)
*/

func checkRows(board [][]string) bool {
	checker := 0 //Number of rows
	for _, row := range board {
		for _, val := range row {
			valint, _ := strconv.Atoi(val)
			checker += valint
		}
		if checker == -5 {
			return true
		}
		checker = 0
	}
	return false

}

func checkColumns(column [][]string) bool {
	checker := 0 //Number of columns

	for i := 0; i < len(column); i++ {
		for j := 0; j < len(column); j++ {
			valint, _ := strconv.Atoi(column[j][i])
			checker += valint
		}
		if checker == -5 {
			return true
		}
		checker = 0
	}
	return false
}

//Board passed in as a slice so edited in place :)
func replaceBingoNumber(board [][]string, bingonumb int) {
	bingonumbstr := strconv.Itoa(bingonumb)
	for _, row := range board {
		for _, val := range row {
			if val == bingonumbstr {
				val = "-1"
			}
		}
	}
}

func sumUnwanted(board [][]string) int {
	summius := 0
	for _, row := range board {
		for _, val := range row {
			valint, _ := strconv.Atoi(val)
			summius += valint
		}
	}
	return summius
}

func checkBoard(bingonumbs chan int, winner chan wins, board [][]string, boardnumber int) {
	bingo_numb := <-bingonumbs

	replaceBingoNumber(board[:][:], bingo_numb)
	columnwin := checkColumns(board[:][:])
	rowwin := checkRows(board[:][:])

	if columnwin == true || rowwin == true {
		sum := sumUnwanted(board[:][:])
		var win = wins{boardnumber, sum, bingo_numb}
		winner <- win
	}

}

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
		// fmt.Println(bingo_numbs)
		break
	}

	//To read boards
	// fmt.Println("Boards")

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
		// fmt.Println(boards[number_board])
		number_board++
		// fmt.Println("End of board")
	}

	//For debugging purposes
	// for i, v := range boards {
	// 	fmt.Printf("Board number: %d\nContent: %v\n", i, v)
	// 	fmt.Println("Testing single value: ", v[0][1])
	// }
	// var bingonumbs chan int
	// var winners chan wins
	var bingo_numbers []int
	bni := strings.Split(bingo_numbs, ",")

	// fmt.Printf("%T\n", bni)
	for _, bingo_number := range bni {
		bn, _ := strconv.Atoi(bingo_number)
		bingo_numbers = append(bingo_numbers, bn)
		// fmt.Println(bingo_numbers)

		// bingonumbs <- bn
	}

	wb := make(map[int]bool)

	for _, bn := range bingo_numbers {
		// for boardnumber := range boards {
		for boardnumber := 1; boardnumber < len(boards)+1; boardnumber++ {
			// fmt.Println(boardnumber)
			win := checkBoard(bn, boards[boardnumber], boardnumber)
			if win.sum != -1 {
				// wb[boardnumber] = true
				// fmt.Println(win)
				if count_trues(wb) == len(boards)-1 && wb[boardnumber] == false {
					fmt.Println("This is the last one to win!")
					fmt.Println(win)
					return
				} else {
					wb[boardnumber] = true
				}
			}
		}
	}

}

func count_trues(winb map[int]bool) int {
	counter := 0
	for _, v := range winb {
		if v == true {
			counter++
		}
	}
	return counter
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
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board); j++ {
			if board[i][j] == bingonumbstr {
				board[i][j] = "-1"
			}
		}
	}
}

func sumUnwanted(board [][]string) int {
	summius := 0
	for _, row := range board {
		for _, val := range row {
			valint, _ := strconv.Atoi(val)
			if valint != -1 {
				summius += valint
			}
		}
	}
	return summius
}

// func checkBoard(bingo_numb int, winner chan wins, board [][]string, boardnumber int) wins {
func checkBoard(bingo_numb int, board [][]string, boardnumber int) wins {
	// bingo_numb := <-bingonumbs
	// fmt.Printf("BN: %d - BingoNumb: %d\n", boardnumber, bingo_numb)
	replaceBingoNumber(board[:][:], bingo_numb)
	// fmt.Println(board)
	columnwin := checkColumns(board[:][:])
	rowwin := checkRows(board[:][:])

	// if columnwin == true || rowwin == true {
	// 	sum_winning := sumUnwanted(board[:][:]) * bingo_numb
	// 	var win = wins{boardnumber, sum_winning, bingo_numb}
	// 	winner <- win
	// } else {
	// 	var win = wins{-1, -1, -1}
	// 	winner <- win
	// }
	if columnwin == true || rowwin == true {
		sum_winning := sumUnwanted(board[:][:]) * bingo_numb
		var win = wins{boardnumber, sum_winning, bingo_numb}
		return win
	}
	var win = wins{-1, -1, -1}
	return win
}

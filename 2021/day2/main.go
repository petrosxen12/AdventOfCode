package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("mainlogger")
var backend = logging.NewLogBackend(os.Stdout, "", 0)

func main() {

	backendlevel := logging.AddModuleLevel(backend)
	backendlevel.SetLevel(logging.INFO, "")

	fmt.Println("Advent of Code Day 2 - Puzzle 1 & 2")

	// Opening file for reading
	var fileName string
	if os.Getenv("DEBUG") == "1" {
		fmt.Println("In Debug Mode")
		fileName = "test_input.txt"
	} else {
		fileName = "input.txt"
	}
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Errorf("Error, can't open file %s", err)
	}
	defer file.Close()

	x_pos := 0
	y_pos := 0
	aim := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		values := strings.Fields(scanner.Text())

		instr := values[0]
		mov, _ := strconv.Atoi(values[1])

		switch instr {
		case "forward":
			// log.Debugf("forward by %d\n",mov)
			x_pos += mov
			y_pos += (aim * mov)
			fmt.Printf("forward by %d x_pos: %d aim: %d\n", mov, x_pos, aim)

		case "down":
			// fmt.Printf("down by %d\n",mov)
			// log.Debugf("down by %d\n",mov)
			// y_pos += mov
			aim += mov
			fmt.Printf("down by %d x_pos: %d aim: %d\n", mov, y_pos, aim)

		case "up":
			// fmt.Printf("up by %d\n",mov)
			// log.Debugf("up by %d\n",mov)
			// y_pos -= mov
			aim -= mov
			fmt.Printf("up by %d x_pos: %d aim: %d\n", mov, y_pos, aim)

			// default:
			// 	fmt.Println()
		}

	}

	fmt.Printf("Final X pos: %d\nFinal Y pos: %d\n", x_pos, y_pos)
	fmt.Printf("Product: %d\n", x_pos*y_pos)
}

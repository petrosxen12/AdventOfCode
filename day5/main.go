package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type line struct {
	x1 int
	y1 int
	x2 int
	y2 int
}

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
	var coordinates_str []string
	for scanner.Scan() {
		coordinates_str = append(coordinates_str, scanner.Text())
		// fmt.Println(bingo_numbs)
	}

	var re = regexp.MustCompile(`(?m)(\d*),(\d*) -> (\d*),(\d*)`)

	// var hydrothermal_vents_map [][]int
	var coordinates []line

	for _, v := range coordinates_str {
		res := re.FindAllStringSubmatch(v, -1)
		// fmt.Println(res)
		for _, number := range res {
			fmt.Printf("x1: %s y1: %s - x2: %s y2: %s\n", number[1], number[2], number[3], number[4])

			x1, _ := strconv.Atoi(number[1])
			y1, _ := strconv.Atoi(number[2])
			x2, _ := strconv.Atoi(number[3])
			y2, _ := strconv.Atoi(number[4])
			l := line{x1, y1, x2, y2}
			coordinates = append(coordinates, l)
		}
	}

	fmt.Println(coordinates)
	fmt.Println("=====================")
	hv_coords := get_horizontal_vertical_lines(coordinates)
	fmt.Println(hv_coords)
}

func get_horizontal_vertical_lines(lines []line) []line {
	var hv_lines []line
	for _, l := range lines {
		if l.x1 == l.x2 || l.y1 == l.y2 {
			hv_lines = append(hv_lines, l)
		}
	}

	return hv_lines
}

func all_covered_points_by_line(l line) []line {
	// TODO: Implement function
}

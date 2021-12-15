package main

import (
	"bufio"
	"fmt"
	"math"
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
	fmt.Println("=========Horizontal & Vertical Lines============")
	hv_coords := get_horizontal_vertical_lines(coordinates)

	//For each coordinate show all current points
	lshold := make(map[line][]line)

	for _, v := range hv_coords {
		fmt.Println("For point: ", v)
		ls := all_covered_points_by_line(v)
		fmt.Println(ls)
		lshold[v] = ls
	}
	overlaps := 0

	overlapping := find_overlapping_points(lshold)
	for k, v := range overlapping {
		if v >= 2 {
			fmt.Println(k)
			overlaps++
		}
	}

	fmt.Println(overlaps)
}

func find_overlapping_points(l map[line][]line) map[line]int {
	ln := make(map[line]int)

	for _, v := range l {
		for _, lni := range v {
			ln[lni]++
		}
	}

	return ln
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
	var covered_points []line

	if l.x1 == l.x2 {
		biggest_val := math.Max(float64(l.y2), float64(l.y1))
		minimum_val := math.Min(float64(l.y2), float64(l.y1))

		for i := minimum_val; i <= biggest_val; i++ {
			v := int(i)
			point := line{l.x1, v, l.x2, v}
			covered_points = append(covered_points, point)
		}
	}

	if l.y1 == l.y2 {
		biggest_val := math.Max(float64(l.x2), float64(l.x1))
		minimum_val := math.Min(float64(l.x2), float64(l.x1))

		for i := minimum_val; i <= biggest_val; i++ {
			v := int(i)
			point := line{v, l.y1, v, l.y2}
			covered_points = append(covered_points, point)
		}
	}
	return covered_points
}

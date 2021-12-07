package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type dt struct {
	column     int
	maxBinary  int
	wantedRows []int
}

type kmax struct {
	key string
	val int
}

func readRow(ch chan dt, allFields []string, column int, bitcrit string) {
	var keepval string
	if bitcrit == "oxy" {
		keepval = "1"
	} else {
		keepval = "0"
	}

	m := make(map[string]int)
	// m["0"] = 0
	// m["1"] = 0
	// fmt.Println(string(val[column]))
	for _, val := range allFields {
		// fmt.Printf("Column: %d, Val: %s Type: %T\n", column, val, val)
		m[string(val[column])] += 1
	}
	// fmt.Println(m)

	var max kmax

	for k, v := range m {
		temp := kmax{k, v}

		if temp.val > max.val {
			max = temp
		} else if temp.val == max.val {
			if temp.key == keepval {
				max = temp
			}
		}
	}
	binaryWithMax, _ := strconv.Atoi(max.key)

	var rows []int
	//Find wanted rows with same maximum bit
	for i, val := range allFields {
		rowM, _ := strconv.Atoi(string(val[column]))
		if int(rowM) == int(binaryWithMax) {
			rows = append(rows, i)
		}

	}

	res := dt{column, binaryWithMax, rows}
	// fmt.Println(column)
	ch <- res
}

func epsilon_rate_calc(gamma_rate string) string {
	var complete_epsilon string
	for _, digit := range gamma_rate {
		switch string(digit) {
		case "1":
			complete_epsilon += "0"
		case "0":
			complete_epsilon += "1"
		}
	}
	return complete_epsilon
}

func make_to_string(sorted_rows []dt) string {
	// fmt.Println(sorted_rows)
	var complete_binary string

	for _, cv := range sorted_rows {
		strvalue := strconv.Itoa(cv.maxBinary)
		complete_binary += strvalue
	}
	return complete_binary
}

func binary_to_decimal(binaryval string) int {
	length_str := len(binaryval)
	var intv float64

	for i, bv := range binaryval {
		i += 1
		bv_int, _ := strconv.Atoi(string(bv))
		intv += (float64(bv_int) * (math.Pow(float64(2), float64(length_str-i))))
	}
	return int(intv)
}

func main() {
	// var m map[string]int

	// Opening file for reading
	var fileName string
	if os.Getenv("DEBUG") == "1" {
		fmt.Println("In Debug Mode")
		fileName = "test_input.txt"
	} else {
		fileName = "input.txt"
	}

	file, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Errorf("Error, can't open file %s", err)
	}
	allfields := strings.Fields(string(file))
	// fmt.Printf("Type: %T\n", allfields)
	// fmt.Println(allfields)

	row_size := len(string(allfields[0]))

	//Oxygen Generator Rating
	sorted_rows := most_significants(row_size, allfields, "oxy")

	fmt.Println(sorted_rows)
	/*
		Steps:
		1. Get rows with same bit at that position
		2. Repopulate file variable to simulate file read
		3. Repeat until

	*/
	for i := 0; i < row_size; i++ {
		var new_wanted_fields []string

		for _, v := range sorted_rows {
			// fmt.Println(v)
			for _, r := range v.wantedRows {
				new_wanted_fields = append(new_wanted_fields, allfields[r])
				// sorted_rows = most_significants(row_size, allfields)
			}
			// fmt.Println(new_wanted_fields)]

		}
	}

	gamma_rate := make_to_string(sorted_rows)
	epsilon_rate := epsilon_rate_calc(gamma_rate)
	fmt.Printf("Value: %v Type: %T Decimal: %d\n", gamma_rate, gamma_rate, binary_to_decimal(gamma_rate))
	fmt.Printf("Value: %v Type: %T Decimal: %d\n", epsilon_rate, epsilon_rate, binary_to_decimal(epsilon_rate))
	final_val := binary_to_decimal(gamma_rate) * binary_to_decimal(epsilon_rate)
	fmt.Println(final_val)

}

func most_significants(row_size int, allfields []string, bitcrit string) []dt {

	rows := make(chan dt)
	for i := 0; i < row_size; i++ {
		go readRow(rows, allfields, i, bitcrit)
		// time.Sleep(100)
	}

	var sorted_rows []dt
	for i := 0; i < row_size; i++ {
		v := <-rows
		// fmt.Println(v)
		sorted_rows = append(sorted_rows, v)
	}
	sort.SliceStable(sorted_rows, func(i, j int) bool {
		if sorted_rows[i].column < sorted_rows[j].column {
			return true
		} else {
			return false
		}
	})

	return sorted_rows
}

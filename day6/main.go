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

	var fish_timers string

	for scanner.Scan() {
		fish_timers = scanner.Text()
	}

	fmt.Println(fish_timers)
	/*
		Each day, a 0 becomes a 6 and adds a new 8 to the end of the list, while each other number decreases by 1 if it was present at the start of the day.
	*/
	fish_life := strings.Split(fish_timers, ",")

	fmt.Println(fish_life)

	days := 3

	var fish_days []int

	for _, fl := range fish_life {
		v, _ := strconv.Atoi(fl)
		fish_days = append(fish_days, v)
	}

	for i := 1; i < days+1; i++ {
		for j := 0; j < len(fish_days); j++ {
			appendius := false
			if fish_days[j] == 0 {
				appendius = true
				fish_days[j] = 6
				fish_days = append(fish_days, 8)
			}
			if fish_days[j] == 8 && appendius == true {
				continue
			} else {
				fish_days[j]--
			}

		}
		fmt.Printf("Day %d: %v\n", i, fish_days)
	}
	fmt.Println(len(fish_days))
}

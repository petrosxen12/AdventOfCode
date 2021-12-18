package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
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
	days := flag.Int("days", 42, "reproduction days of fishies")
	flag.Parse()

	fish_life := strings.Split(fish_timers, ",")

	fmt.Println(fish_life)

	intdays := *days
	var fish_days []int

	for _, fl := range fish_life {
		v, _ := strconv.Atoi(fl)
		fish_days = append(fish_days, v)
	}

	// var fishies_spawn []int
	sort.Ints(fish_days)

	compressor := make(map[int]int)

	for _, f := range fish_days {
		compressor[f]++
	}
	fmt.Println(compressor)

	// Debugging return
	// return

	for i := 1; i < intdays+1; i++ {
		appender := false

		for k, v := range compressor {
			if v > 0 && k >= 0 && !(v < 0) {
				compressor[k] -= 1
				compressor[k-1] += 1
			} else {
				appender = true
			}
		}
		// fish_days = append(fish_days, fishies_spawn...)
		// fishies_spawn = nil
		if appender {
			compressor[6]++
			compressor[8]++
		}
		fmt.Printf("Day %d: %v\n", i, compressor)
	}
	fmt.Println(len(fish_days))
}

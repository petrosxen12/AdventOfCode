package main

import (
	"bufio"
	"flag"
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

	var fishies_spawn []int

	compressor := make(map[int]int)

	for _, f := range fish_days {
		compressor[f]++
	}
	fmt.Println(compressor)

	// Debugging return
	// return

	for i := 1; i < intdays+1; i++ {
		for j := 0; j < len(fish_days); j++ {
			fish_days[j]--
			if fish_days[j] == -1 {
				fish_days[j] = 6
				fishies_spawn = append(fishies_spawn, 8)
			}
		}
		fish_days = append(fish_days, fishies_spawn...)
		fishies_spawn = nil
		// fmt.Printf("Day %d: %v\n", i, fish_days)
	}
	fmt.Println(len(fish_days))
}

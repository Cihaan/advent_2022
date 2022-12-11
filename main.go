package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

const DAY_ONE string = "day_one.txt"

func main() {
	fmt.Println(getMostConsumer())
}

func readFile(path string) (*bufio.Scanner, *os.File) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	return bufio.NewScanner(file), file
}

// day one
func getMostConsumer() []int {
	var count int = 0
	var elves []int
	var scanner, file = readFile(DAY_ONE)

	for scanner.Scan() {
		var parsed, err = strconv.Atoi(scanner.Text())
		count += parsed

		if err != nil {
			elves = append(elves, count)
			count = 0
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sort.Ints(elves)

	defer file.Close()

	return elves
}

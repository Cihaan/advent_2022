package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

const FILE_NAME string = "input.txt"

func main() {
	fmt.Print(getMaxConsumer())

}

func readFile(path string) (*bufio.Scanner, *os.File) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	return bufio.NewScanner(file), file
}

func getMaxConsumer() int {
	var res int = 0
	var count int = 0
	var scanner, file = readFile(FILE_NAME)

	for scanner.Scan() {
		var parsed, err = strconv.Atoi(scanner.Text())
		count += parsed

		if err != nil {
			if res < count {
				res = count
				count = 0
			} else {
				count = 0
			}
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	return res
}

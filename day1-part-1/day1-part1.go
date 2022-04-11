package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("../data/day1.txt")

	if err != nil {
		log.Fatal("Failed to open file")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var previousDepth int = 0
	var depthCounter int = 0
	var index int = 0

	for scanner.Scan() {
		currentDepth, err := strconv.Atoi(scanner.Text())

		if err != nil {
			log.Fatal("Failed to convert string to number")
		}

		if currentDepth > previousDepth && index != 0 {
			depthCounter++
		}

		previousDepth = currentDepth
		index++
	}

	log.Printf("%v", depthCounter)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

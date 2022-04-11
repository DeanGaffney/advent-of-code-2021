package main

import (
	"bufio"
	"fmt"
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

	var depths []int

	var index int = 0

	for scanner.Scan() {
		currentDepth, err := strconv.Atoi(scanner.Text())

		if err != nil {
			log.Fatal("Failed to convert string to number")
		}

		depths = append(depths, currentDepth)

		index++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var increased int = 0

	for i := range depths {
		fmt.Println(depths[i])
		if i+3 < len(depths) {
			var firstWindow int = depths[i] + depths[i+1] + depths[i+2]
			fmt.Println(depths[i], depths[i+1], depths[i+2])
			var secondWindow int = depths[i+1] + depths[i+2] + depths[i+3]
			fmt.Println(depths[i+1], depths[i+2], depths[i+3])
			fmt.Println(firstWindow, secondWindow)

			if secondWindow > firstWindow {
				increased++
			}
		}
	}

	fmt.Println("Increased", increased)
}

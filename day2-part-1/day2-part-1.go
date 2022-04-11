package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("../data/day2.txt")

	if err != nil {
		log.Fatal("Failed to open file")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var horizontalPosition int = 0

	var depth int = 0

	for scanner.Scan() {
		instruction := scanner.Text()

		splitInstruction := strings.Split(instruction, " ")

		direction := splitInstruction[0]
		numberStr := splitInstruction[1]

		number, err := strconv.Atoi(numberStr)

		if err != nil {
			log.Fatal("Failed to convert string to number")
		}

		switch direction {
		case "forward":
			horizontalPosition += number
		case "down":
			depth += number
		case "up":
			depth -= number
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(horizontalPosition * depth)
}

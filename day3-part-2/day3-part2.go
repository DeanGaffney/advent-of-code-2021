package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type IGammaRate interface {
	addOne(int)
	addZero(int)
	getMostCommonBit() string
}

type GammaRate struct {
	ones  int
	zeros int
}

func (gammaRate *GammaRate) addOne(one int) {
	gammaRate.ones++
}

func (gammaRate *GammaRate) addZero(zero int) {
	gammaRate.zeros++
}

func (gammaRate *GammaRate) sameNumberOfBits() bool {
	return gammaRate.zeros == gammaRate.ones
}

func (gammaRate GammaRate) getMostCommonBit() string {
	if gammaRate.ones > gammaRate.zeros {
		return "1"
	} else {
		return "0"
	}
}

func (gammaRate GammaRate) getLeastCommonBit() string {
	if gammaRate.ones > gammaRate.zeros {
		return "0"
	} else {
		return "1"
	}
}

func readBinaryNumbersFromFile() [][]string {
	file, err := os.Open("../data/day3.txt")

	if err != nil {
		log.Fatal("Failed to open file")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var binaryNumbersMatrix [][]string

	for scanner.Scan() {
		binaryNumber := scanner.Text()

		var row []string

		for i := range binaryNumber {
			character := string(binaryNumber[i])
			row = append(row, character)
		}

		binaryNumbersMatrix = append(binaryNumbersMatrix, row)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return binaryNumbersMatrix
}

func calculateRating(columnIndex int, binaryNumbers [][]string, bitTypeOfInterest string, bitToUseInCaseOfEquality string) int {
	var numberOfRows = len(binaryNumbers)

	// we hit a point where there is only one number so return the value
	if numberOfRows == 1 {
		var bitString string
		for i := range binaryNumbers[0] {
			bitString += binaryNumbers[0][i]
		}
		rate, err := strconv.ParseInt(bitString, 2, 64)
		if err != nil {
			log.Fatal(err)
		}

		return int(rate)
	}

	var filteredBinaryNumbers [][]string

	gammaRate := GammaRate{}

	for j := 0; j < numberOfRows; j++ {
		var character string = binaryNumbers[j][columnIndex]
		if character == "1" {
			gammaRate.addOne(1)
		} else {
			gammaRate.addZero(1)
		}
	}

	var bitOfInterest string

	if bitTypeOfInterest == "mostCommon" {
		bitOfInterest = gammaRate.getMostCommonBit()
	} else if bitTypeOfInterest == "leastCommon" {
		bitOfInterest = gammaRate.getLeastCommonBit()
	}

	if gammaRate.sameNumberOfBits() {
		bitOfInterest = bitToUseInCaseOfEquality
	}

	for j := 0; j < numberOfRows; j++ {
		var character string = binaryNumbers[j][columnIndex]
		if character == bitOfInterest {
			filteredBinaryNumbers = append(filteredBinaryNumbers, binaryNumbers[j])
		}
	}

	return calculateRating(columnIndex+1, filteredBinaryNumbers, bitTypeOfInterest, bitToUseInCaseOfEquality)
}

func main() {

	var binaryNumbers [][]string = readBinaryNumbersFromFile()

	oxygenRating := calculateRating(0, binaryNumbers, "mostCommon", "1")
	scrubberRating := calculateRating(0, binaryNumbers, "leastCommon", "0")

	fmt.Println("Life Support Rating", oxygenRating*scrubberRating)
}

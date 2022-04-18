package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

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

func main() {

	var binaryNumbers [][]string = readBinaryNumbersFromFile()
	var gammaRates []GammaRate

	var numberOfRows = len(binaryNumbers)
	var numberOfCols = len(binaryNumbers[0])

	for i := 0; i < numberOfCols; i++ {

		gammaRate := GammaRate{}

		for j := 0; j < numberOfRows; j++ {
			var character string = binaryNumbers[j][i]
			if character == "1" {
				gammaRate.addOne(1)
			} else {
				gammaRate.addZero(1)
			}
		}

		gammaRates = append(gammaRates, gammaRate)
	}

	var mostCommonBits string
	var leastCommonBits string

	for i := range gammaRates {
		gammaRate := gammaRates[i]
		mostCommonBits += gammaRate.getMostCommonBit()
		leastCommonBits += gammaRate.getLeastCommonBit()
	}

	rate, gammaRateErr := strconv.ParseInt(mostCommonBits, 2, 64)
	epsilon, epsilonErr := strconv.ParseInt(leastCommonBits, 2, 64)

	if gammaRateErr != nil || epsilonErr != nil {
		log.Fatal("Failed to convert string to number")
	}

	fmt.Println(rate * epsilon)
}

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type Bingo struct {
	numbersToDraw []int
	gameBoards    [][][]int
}

func (bingo *Bingo) addNumberToDraw(number int) {
	bingo.numbersToDraw = append(bingo.numbersToDraw, number)
}

func (bingo *Bingo) addGameBoard(gameBoard [][]int) {
	bingo.gameBoards = append(bingo.gameBoards, gameBoard)
}

func (bingo *Bingo) markNumber(number int) {
	for gameBoard := range bingo.gameBoards {
		for row := range bingo.gameBoards[gameBoard] {
			for column := range bingo.gameBoards[gameBoard][row] {
				if bingo.gameBoards[gameBoard][row][column] == number {
					bingo.gameBoards[gameBoard][row][column] = -1
				}
			}
		}
	}
}

func (bingo Bingo) hasWinningGameBoard() (bool, int) {
	for gameBoard := range bingo.gameBoards {
		for row := range bingo.gameBoards[gameBoard] {

			var rowCount int
			for i := 0; i < len(bingo.gameBoards[gameBoard][row]); i++ {
				currentRowNumber := bingo.gameBoards[gameBoard][row][i]

				if currentRowNumber != -1 {
					break
				}

				if currentRowNumber == -1 {
					rowCount++
				}

			}

			if rowCount == len(bingo.gameBoards[gameBoard][row]) {
				return true, gameBoard
			}

			var colCount int
			for column := range bingo.gameBoards[gameBoard][row] {
				currentColumnNumber := bingo.gameBoards[gameBoard][column][row]

				if currentColumnNumber == -1 {
					colCount++
				}
			}

			if colCount == len(bingo.gameBoards[gameBoard][row]) {
				return true, gameBoard
			}
		}

	}

	return false, -1

}

func (bingo Bingo) sumUnmarkedNumbers(winningGameBoard int) int {
	var sum int
	for row := range bingo.gameBoards[winningGameBoard] {
		for column := range bingo.gameBoards[winningGameBoard][row] {
			if bingo.gameBoards[winningGameBoard][row][column] != -1 {
				sum += bingo.gameBoards[winningGameBoard][row][column]
			}
		}
	}

	return sum
}

func readInputFile() Bingo {
	file, err := os.Open("../data/day4.txt")

	if err != nil {
		log.Fatal("Failed to open file")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	bingo := Bingo{}
	var currentGameBoard [][]int

	var lineCount int

	for scanner.Scan() {
		currentLine := scanner.Text()

		if lineCount == 0 {
			numbersToDraw := scanner.Text()
			splitRegex := regexp.MustCompile(",")
			splitNumbers := splitRegex.Split(numbersToDraw, -1)
			for number := range splitNumbers {
				number, err := strconv.Atoi(splitNumbers[number])
				if err != nil {
					log.Fatal("Failed to convert string to int when converting numbers to draw")
				}
				bingo.addNumberToDraw(number)
			}
			lineCount++
			continue
		}

		if currentLine == "" && lineCount <= 1 {
			lineCount++
			continue
		}

		if currentLine == "" && lineCount > 1 {
			bingo.addGameBoard(currentGameBoard)
			currentGameBoard = [][]int{}
			lineCount++
			continue
		}

		regex := regexp.MustCompile(`\s+`)
		split := regex.Split(currentLine, -1)
		var numbersInRow []int

		for i := 0; i < len(split); i++ {
			if split[i] == "" {
				continue
			}

			number, err := strconv.Atoi(split[i])
			if err != nil {
				log.Fatal("Failed to convert string to int", err)
			}
			numbersInRow = append(numbersInRow, number)
		}

		currentGameBoard = append(currentGameBoard, numbersInRow)

		lineCount++
	}

	// handle the last game board
	if len(currentGameBoard) > 0 {
		bingo.addGameBoard(currentGameBoard)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return bingo
}

func main() {
	bingo := readInputFile()
	for i := 0; i < len(bingo.numbersToDraw); i++ {
		currentNumber := bingo.numbersToDraw[i]

		bingo.markNumber(currentNumber)

		hasWinningGameBoard, winningGameBoard := bingo.hasWinningGameBoard()

		if hasWinningGameBoard {
			fmt.Println(currentNumber * bingo.sumUnmarkedNumbers(winningGameBoard))
			break
		}
	}

}

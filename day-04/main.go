package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Cell struct {
	WasDrawn bool
	Value    int
}

type Board struct {
	Cells [5][5]Cell
}

func parseNumbersDrawn(allNumbersDrawn string) (numbersDrawn []int) {
	numbersDrawnStr := strings.Split(allNumbersDrawn, ",")

	numbersDrawn = make([]int, len(numbersDrawnStr))

	var err error

	for i, numStr := range numbersDrawnStr {
		numbersDrawn[i], err = strconv.Atoi(numStr)
		if err != nil {
			log.Fatal(err)
		}
	}

	return numbersDrawn
}

func readBoardLine(boardStr string) (board [5]Cell) {
	boardStrArray := strings.Fields(boardStr)

	var err error

	for i, valStr := range boardStrArray {
		board[i].Value, err = strconv.Atoi(valStr)
		if err != nil {
			log.Fatal(err)
		}
	}

	return board
}

func readInput() (numbersDrawn []int, boards []Board) {
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	// Read and parse the first line //////////////////////////////////////////
	scanner.Scan()
	numbersDrawn = parseNumbersDrawn(scanner.Text())

	newBoard := Board{}
	newBoardLine := 0

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			newBoardLine = 0
			continue
		}

		newBoard.Cells[newBoardLine] = readBoardLine(line)
		newBoardLine++

		if newBoardLine == 5 {
			boards = append(boards, newBoard)
		}
	}

	return numbersDrawn, boards
}

func checkIfBingo(board Board, line, col int) (bingo bool) {
	// Try to seep the columns in the given line
	for i := 0; i < 5; i++ {
		if !board.Cells[line][i].WasDrawn {
			break
		}
		if i == 4 {
			return true
		}
	}

	// Try to seep the lines in the given column
	for i := 0; i < 5; i++ {
		if !board.Cells[i][col].WasDrawn {
			break
		}
		if i == 4 {
			return true
		}
	}

	return false
}

func fillBoardNumber(currNumber int, board *Board) (bingo bool) {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if board.Cells[i][j].Value == currNumber {
				board.Cells[i][j].WasDrawn = true
				if checkIfBingo(*board, i, j) {
					return true
				}
			}
		}
	}
	return false
}

func sumUnusedNumbers(board Board) int {
	sum := 0
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if !board.Cells[i][j].WasDrawn {
				sum += board.Cells[i][j].Value
			}
		}
	}
	return sum
}

func printBoard(board Board) {
	for i := 0; i < 5; i++ {
		fmt.Println(board.Cells[i])
	}
}

func main() {
	numbersDrawn, boards := readInput()
	firstPart(numbersDrawn, boards)
	secondPart(numbersDrawn, boards)
}

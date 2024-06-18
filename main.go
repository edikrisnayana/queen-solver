package main

import (
	"fmt"
	"os"
	"strings"
)

var COLOR_DISTINCT int = 0

func main() {
	// put your sudoku board in here
	board := [][]int{
		{1, 1, 1, 3, 3, 3, 3, 3, 3},
		{1, 6, 6, 6, 6, 5, 5, 5, 3},
		{1, 7, 6, 8, 8, 8, 5, 5, 3},
		{1, 7, 8, 8, 8, 8, 8, 4, 3},
		{1, 7, 8, 8, 0, 8, 8, 4, 3},
		{1, 7, 8, 8, 8, 8, 8, 4, 3},
		{2, 7, 7, 8, 8, 8, 4, 4, 3},
		{2, 4, 4, 4, 4, 4, 4, 4, 3},
		{2, 2, 3, 3, 3, 3, 3, 3, 3},
	}

	for _, row := range board {
		for _, val := range row {
			if COLOR_DISTINCT < val {
				COLOR_DISTINCT = val
			}
		}
	}

	COLOR_DISTINCT++

	solution := make([][]rune, 9)
	for i := range solution {
		solution[i] = make([]rune, 9)
	}

	for i, row := range board {
		for j, val := range row {
			solution[i][j] = rune(48 + val)
		}
	}

	firstColorTiles := findColor(board, solution)

	currState := make([][]rune, 9)
	for i := range currState {
		currState[i] = make([]rune, 9)
	}
	moveState(&solution, &currState)
	for _, tile := range firstColorTiles {
		i := tile[0]
		j := tile[1]
		markSameColor(board, &solution, board[i][j])
		markDiagonal(&solution, i, j)
		markLine(&solution, i, j)
		solution[i][j] = 'q'

		proccess(board, &solution)
		if isValid(&solution) {
			print(solution)
			break
		}

		moveState(&currState, &solution)
	}
}

func proccess(board [][]int, solution *[][]rune) bool {
	currState := make([][]rune, 9)
	for i := range currState {
		currState[i] = make([]rune, 9)
	}
	moveState(solution, &currState)
	firstColorTiles := findColor(board, *solution)

	for _, tile := range firstColorTiles {
		i := tile[0]
		j := tile[1]
		if isEligible(solution, i, j) {
			markSameColor(board, solution, board[i][j])
			markDiagonal(solution, i, j)
			markLine(solution, i, j)
			(*solution)[i][j] = 'q'

			proccess(board, solution)
			if isValid(solution) {
				return true
			}
			moveState(&currState, solution)
		}
	}

	return isValid(solution)
}

func moveState(state1 *[][]rune, state2 *[][]rune) {
	for i, row := range *state1 {
		for j, val := range row {
			(*state2)[i][j] = val
		}
	}
}

func findColor(board [][]int, solution [][]rune) [][]int {
	colors := make([]int, COLOR_DISTINCT)
	for i, row := range board {
		for j, val := range row {
			if solution[i][j] != 'q' && solution[i][j] != 'x' {
				colors[val]++
			}
		}
	}

	min := 82
	minColor := -1
	for color, val := range colors {
		if val > 0 && min > val {
			min = val
			minColor = color
		}
	}

	firstColorTiles := make([][]int, 0)
	if minColor < 0 {
		return firstColorTiles
	}

	for i, row := range board {
		for j, val := range row {
			if val == minColor && solution[i][j] != 'q' && solution[i][j] != 'x' {
				firstColorTiles = append(firstColorTiles, []int{i, j})
			}
		}
	}

	return firstColorTiles
}

func toString(solution [][]rune) string {
	sb := strings.Builder{}
	for i := 1; i <= 2; i++ {
		for j := 0; j < 19; j++ {
			sb.WriteString("-")
		}
		sb.WriteString("\n")
	}

	for i := 0; i < 9; i++ {
		sb.WriteString("|")
		for j := 0; j < 9; j++ {
			sb.WriteString(fmt.Sprintf("%c|", solution[i][j]))
		}
		sb.WriteString("\n")
		for j := 0; j < 19; j++ {
			sb.WriteString("-")
		}
		sb.WriteString("\n")
	}

	for j := 0; j < 19; j++ {
		sb.WriteString("-")
	}

	return sb.String()
}

func print(solution [][]rune) {

	file, createErr := os.Create("output.txt")
	if createErr != nil {
		fmt.Println(createErr)
	}

	file.WriteString(toString(solution))
}

func clear(solution *[][]rune) {
	for i := range *solution {
		for j := range (*solution)[i] {
			(*solution)[i][j] = ' '
		}
	}
}

func isValid(solution *[][]rune) bool {
	count := 0
	for i := range *solution {
		for j := range (*solution)[i] {
			if (*solution)[i][j] == 'q' {
				count++
			}
		}
	}
	return count == COLOR_DISTINCT
}

func markSameColor(board [][]int, solution *[][]rune, color int) {
	for i, row := range board {
		for j := range row {
			if board[i][j] == color {
				(*solution)[i][j] = 'x'
			}
		}
	}
}

func markDiagonal(solution *[][]rune, i int, j int) {
	if i > 0 && j > 0 {
		(*solution)[i-1][j-1] = 'x'
	}
	if i < 8 && j < 8 {
		(*solution)[i+1][j+1] = 'x'
	}
	if i > 0 && j < 8 {
		(*solution)[i-1][j+1] = 'x'
	}
	if i < 8 && j > 0 {
		(*solution)[i+1][j-1] = 'x'
	}
}

func markLine(solution *[][]rune, i int, j int) {
	for k := 0; k < 9; k++ {
		if k != i {
			(*solution)[k][j] = 'x'
		}
	}
	for k := 0; k < 9; k++ {
		if k != j {
			(*solution)[i][k] = 'x'
		}
	}
}

func isEligible(solution *[][]rune, i int, j int) bool {
	if (*solution)[i][j] == 'q' || (*solution)[i][j] == 'x' {
		return false
	}

	for k := 0; k < 9; k++ {
		if k != i && (*solution)[k][j] == 'q' {
			return false
		}
	}
	for k := 0; k < 9; k++ {
		if k != j && (*solution)[i][k] == 'q' {
			return false
		}
	}

	if i > 0 && j > 0 && (*solution)[i-1][j-1] == 'q' {
		return false
	}
	if i < 8 && j < 8 && (*solution)[i+1][j+1] == 'q' {
		return false
	}
	if i > 0 && j < 8 && (*solution)[i-1][j+1] == 'q' {
		return false
	}
	if i < 8 && j > 0 && (*solution)[i+1][j-1] == 'q' {
		return false
	}

	return true
}

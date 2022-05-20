package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	emptyBoard := "-"
	strList := [][]string{{emptyBoard, emptyBoard, emptyBoard}, {emptyBoard, emptyBoard, emptyBoard}, {emptyBoard, emptyBoard, emptyBoard}}
	startPlayer := "O"
	for GameStat(strList, &startPlayer) {
	}
}

func GameStat(board [][]string, player *string) bool {
	ClearScreen()
	ShowBoard(board)

	fmt.Print("Enter Your Position: ")

	var input int
	fmt.Scanln(&input)
	if input < 1 || input > 9 {
		panic("Invalid input")
	}

	y := (input - 1) / 3
	x := (input - 1) % 3

	if board[y][x] == "O" || board[y][x] == "X" {
		panic("Value already replaced")
	}
	board[y][x] = *player
	if CheckWinner(board, *player) == true {
		ClearScreen()
		ShowBoard(board)
		fmt.Printf("Winner is %v\n", *player)
		return false
	}

	SwitchPlayer(player)
	return true
}

func ShowBoard(board [][]string) {
	for _, row := range board {
		fmt.Println(row)
	}
}

func ClearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func SwitchPlayer(input *string) {
	switch *input {
	case "O":
		*input = "X"
	default:
		*input = "O"
	}

}

func CheckWinner(board [][]string, checkValue string) bool {
	countDiagTopleftToBotright, countDiagToprightToBotleft := 0, 0
	for i, row := range board {
		// check row
		countRow, countColumn := 0, 0
		for j, block := range row {
			if block == checkValue {
				countRow++
			}
			if board[j][i] == checkValue {
				countColumn++
			}
		}
		if countRow == 3 || countColumn == 3 {
			return true
		}

		// check diag
		if board[i][i] == checkValue {
			countDiagTopleftToBotright++
		}
		if board[i][2-i] == checkValue {
			countDiagToprightToBotleft++
		}
	}
	if countDiagTopleftToBotright == 3 || countDiagToprightToBotleft == 3 {
		return true
	}
	return false
}

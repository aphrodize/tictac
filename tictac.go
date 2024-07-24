package main

import (
    "fmt"
)

const (
    EMPTY = iota
    PLAYER_X
    PLAYER_O
)

var board [3][3]int
var currentPlayer int

func printBoard() {
    symbols := map[int]string{
        EMPTY:   " ",
        PLAYER_X: "X",
        PLAYER_O: "O",
    }
    fmt.Println("  0 1 2")
    for i, row := range board {
        fmt.Printf("%d ", i)
        for _, cell := range row {
            fmt.Printf("%s ", symbols[cell])
        }
        fmt.Println()
    }
}

func checkWin(player int) bool {
    for i := 0; i < 3; i++ {
        if board[i][0] == player && board[i][1] == player && board[i][2] == player {
            return true
        }
        if board[0][i] == player && board[1][i] == player && board[2][i] == player {
            return true
        }
    }
    if board[0][0] == player && board[1][1] == player && board[2][2] == player {
        return true
    }
    if board[0][2] == player && board[1][1] == player && board[2][0] == player {
        return true
    }
    return false
}

func checkDraw() bool {
    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            if board[i][j] == EMPTY {
                return false
            }
        }
    }
    return true
}

func main() {
    currentPlayer = PLAYER_X
    for {
        printBoard()
        var row, col int
        fmt.Printf("Player %s, enter your move (row and column): ", map[int]string{PLAYER_X: "X", PLAYER_O: "O"}[currentPlayer])
        fmt.Scanf("%d %d", &row, &col)
        if row < 0 || row >= 3 || col < 0 || col >= 3 || board[row][col] != EMPTY {
            fmt.Println("Invalid move. Try again.")
            continue
        }
        board[row][col] = currentPlayer
        if checkWin(currentPlayer) {
            printBoard()
            fmt.Printf("Player %s wins!\n", map[int]string{PLAYER_X: "X", PLAYER_O: "O"}[currentPlayer])
            break
        }
        if checkDraw() {
            printBoard()
            fmt.Println("It's a draw!")
            break
        }
        if currentPlayer == PLAYER_X {
            currentPlayer = PLAYER_O
        } else {
            currentPlayer = PLAYER_X
        }
    }
}


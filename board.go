package main

import "fmt"

func CreateBoard(xSize int, ySize int) [][]rune {
  
  board := make([][]rune, ySize)
  for i := range board {
    board[i] = make([]rune, xSize)
  }

  for i := 0; i < ySize; i++ {
    for j := 0; j < xSize; j++ {
      board[i][j] = '☐'
    }
  }
  return board
}       

func DrawBoard(board [][]rune)  {
  for i := range board {
    line := ""
    for j := range board[i] {
      line += string(board[i][j])
    }
    fmt.Println(line)
  }
}

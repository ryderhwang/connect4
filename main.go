package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

var board [][]int
var wg sync.WaitGroup


func chanOwner() <-chan byte {
	defer  wg.Done()
	orders := make(chan byte, 10)

	go func() {
		defer close(orders)
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Split(bufio.ScanWords)
		for scanner.Scan() {
			b := scanner.Bytes()
			if len(b) > 1 || b[0] < '0' || b[0] > '6' {
				fmt.Println("Invalid input, please make a move between 0 and 6.")
			} else {
				orders <- b[0]
			}
		}
	}()
	return orders
}

func move(order <-chan byte) {
	defer wg.Done()
	player := 1
	cnt := 0
	for {
		select{
		case col := <-order:
			col -= '0'
			ind := 0
			for ;ind < 6; ind++ {
				if board[ind][col] == 0 {
					board[ind][col] = player
					cnt++
					isFinished := checkFinish()
					if isFinished != 0 {
						fmt.Printf("WINNER: Player %d\n", isFinished)
						return
					}
					if cnt == 42 {
						fmt.Println("DRAW")
						return
					}
					player = player ^ 3
					break
				}
			}
			if ind == 6 {
				fmt.Printf("Column %d is full, pick another column!\n", col)
			}
		}
	}
}

func checkFinish() int{
	var count int
	for i := 0; i < 6; i++ {
		count = 1
		player := board[i][0]
		for j := 1; j < 7; j++ {
			if board[i][j] == 0 {
				count = 0
				player = 0
				continue
			}
			if board[i][j] == player {
				count++
				if count >= 4 {
					return player
				}
			} else {
				player = board[i][j]
				count = 1
			}
		}
	}
	for i := 0; i < 7; i++ {
		count = 1
		player := board[0][i]
		for j := 1; j < 6; j++ {
			if board[j][i] == 0 {
				count = 0
				player = 0
				continue
			}
			if board[j][i] == player {
				count++
				if count >= 4 {
					return player
				}
			} else {
				player = board[j][i]
				count = 1
			}
		}
	}
	for rowStart := 0; rowStart <=  2; rowStart++{
		count = 0
		row := rowStart
		var col int
		player := 1
		for row < 6 && col < 7 {
			if board[row][col] == player{
				count++
				if count >= 4 {
					return player
				}
			} else {
				player = player ^ 3
				count = 1
			}
			row++
			col++
		}
	}
	for colStart := 0; colStart <=  3; colStart++{
		count = 0
		row := 0
		col := colStart
		player := 1
		for row < 6 && col < 7 {
			if board[row][col] == player{
				count++
				if count >= 4 {
					return player
				}
			} else {
				player = player ^ 3
				count = 1
			}
			row++
			col++
		}
	}
	return 0
}

func main() {
	for i := 0; i < 6; i++ {
		board = append(board, make([]int, 7))
	}
	wg.Add(2)
	orders := chanOwner()
	go move(orders)
	wg.Wait()
}

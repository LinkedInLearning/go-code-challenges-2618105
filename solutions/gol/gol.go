package gol

func Evolve(board [][]bool) [][]bool {
	newBoard := make([][]bool, len(board))
	for i := range newBoard {
		newBoard[i] = make([]bool, len((board)[i]))
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len((board)[0]); j++ {
			newBoard[i][j] = nextCellState(board, i, j)
		}
	}

	return newBoard
}

func nextCellState(board [][]bool, x int, y int) bool {
	neighbours := countNeighbours(board, x, y)

	isAlive := board[x][y]

	if isAlive && (neighbours == 2 || neighbours == 3) {
		return true
	} else if !isAlive && neighbours == 3 {
		return true
	} else {
		return false
	}
}

func countNeighbours(board [][]bool, x int, y int) int {
	neighbours := 0

	for i := x - 1; i <= x+1; i++ {
		for j := y - 1; j <= y+1; j++ {
			if i == x && j == y {
				// cell itself
				continue
			}

			if i < 0 || i >= len(board) {
				// out of bounds
				continue
			}

			if j < 0 || j >= len(board[0]) {
				// out of bounds
				continue
			}

			if board[i][j] {
				neighbours++
			}
		}
	}

	return neighbours
}

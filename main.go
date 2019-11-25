package main

import "fmt"

func toKey (i, j int) (key string) {
	return fmt.Sprintf("%d,%d", i, j)
}

func Dfs (word string, currentIndex int, board [][]byte, seen map[string]bool, i, j int) bool {
	if currentIndex >= len(word) {
		return true
	}

	if i < 0 || j < 0 || i >= len(board) || j >= len(board[0]) {
		return false
	}

	if board[i][j] != word[currentIndex] {
		return false
	}

	key := toKey(i, j)

	if _, saw := seen[key]; saw {
		return false
	}

	seen[key] = true

	switch {
	case Dfs(word, currentIndex+1, board, seen, i -1, j):
		return true
	case Dfs(word, currentIndex+1, board, seen, i+1, j):
		return true
	case Dfs(word, currentIndex+1, board, seen, i, j+1):
		return true
	case Dfs(word, currentIndex+1, board, seen, i, j-1):
		return true
	default:
		delete(seen, key)
		return false
	}
}

func exist(board [][]byte, word string) bool {
	if len(word) == 0 {
		return false
	}

	for i := 0; i < len(board); i++ {
		seen := make(map[string]bool)

		for j := 0; j < len(board[i]); j++ {
			if Dfs(word, 0, board, seen, i, j) {
				return true
			}
		}
	}

	return false
}

func main() {
	fmt.Printf("%v\n", exist([][]byte{{'A','B','C','E'}, {'S','F','C','S'}, {'A','D','E','E'}}, "ABCCE"))
}
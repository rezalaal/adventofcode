package main

import (
	"bufio"
	"fmt"
	"os"
)

// Directions for searching: (row increment, column increment)
var directions = [8][2]int{
	{0, 1},   // right
	{0, -1},  // left
	{1, 0},   // down
	{-1, 0},  // up
	{1, 1},   // down-right
	{-1, -1}, // up-left
	{1, -1},  // down-left
	{-1, 1},  // up-right
}

// Function to check if a word exists in a specific direction
func findWord(grid [][]rune, row, col int, word string, dRow, dCol int) bool {
	for i := 0; i < len(word); i++ {
		r := row + i*dRow
		c := col + i*dCol

		// Check boundaries
		if r < 0 || r >= len(grid) || c < 0 || c >= len(grid[0]) || grid[r][c] != rune(word[i]) {
			return false
		}
	}
	return true
}

// Function to count all occurrences of a word in the grid
func countWordOccurrences(grid [][]rune, word string) int {
	count := 0
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[0]); col++ {
			// Check all directions
			for _, dir := range directions {
				if findWord(grid, row, col, word, dir[0], dir[1]) {
					count++
				}
			}
		}
	}
	return count
}

// Main function
func main() {
	// Open input file
	filename := "input.txt"
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Read grid from file
	var grid [][]rune
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, []rune(line))
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Word to search
	word := "XMAS"

	// Count occurrences
	count := countWordOccurrences(grid, word)
	fmt.Printf("The word '%s' occurs %d times in the grid.\n", word, count)
}

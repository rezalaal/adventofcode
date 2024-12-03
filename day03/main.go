package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	filename := "./input3.txt"

	// Open the file
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// Read the file content using bufio
	var content strings.Builder
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		content.WriteString(scanner.Text())
	}

	// Check for scanner errors
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		return
	}

	text := strings.TrimSpace(content.String())

	// Regular expression to match mul(x, y)
	regex := `mul\((\d{1,3}),(\d{1,3})\)`

	// Compile the regex
	re := regexp.MustCompile(regex)

	// Find all matches
	matches := re.FindAllStringSubmatch(text, -1)

	// Create a map to store x and y values
	results := make(map[string][2]int)

	// Process each match
	for i, match := range matches {
		// Convert x and y to integers
		x, _ := strconv.Atoi(match[1])
		y, _ := strconv.Atoi(match[2])
		// Store in the map with a unique key
		results[fmt.Sprintf("match_%d", i+1)] = [2]int{x, y}
	}

	sum := 0
	// Print the map
	for key, value := range results {
		sum += value[0] * value[1]
		fmt.Printf("%s: x=%d, y=%d\n", key, value[0], value[1])
	}
	fmt.Println("Sum:", sum)
}

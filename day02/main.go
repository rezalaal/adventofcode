package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sum int

func checkLevels(fields []string) bool {
	fmt.Println(strings.Join(fields, "-"))
	increasing := true
	decreasing := true

	for i := 1; i < len(fields); i++ {
		// Convert strings to integers
		prev, err1 := strconv.Atoi(fields[i-1])
		curr, err2 := strconv.Atoi(fields[i])
		if err1 != nil || err2 != nil {
			fmt.Printf("Error: Non-numeric value found: %v or %v\n", fields[i-1], fields[i])
			return false
		}

		// Check step size
		step := curr - prev
		if step > 3 || step < -3 || step == 0 {
			return false
		}
		

		// Update flags for increasing or decreasing
		if curr > prev {
			decreasing = false
		} else if curr < prev {
			increasing = false
		}

		// If both flags are false, return false early
		if !increasing && !decreasing {
			return false
		}
	}
	return true
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()	

	
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		
		line := scanner.Text()
		fields := strings.Fields(line)
		result := checkLevels(fields)
		fmt.Println(result)
		if result {			
			sum++
		}
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Println("Sum: ", sum)
}
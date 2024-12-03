package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)



func main() {
	file, err := os.Open("./input1.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()	

	var column1,column2 []int
	var sum float64 = 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Split the line into two fields
		line := scanner.Text()
		fields := strings.Fields(line)

		if len(fields) != 2 {
			fmt.Println("Invalid line:", line)
			continue
		}

		// Convert the fields to integers
		num1, err1 := strconv.Atoi(fields[0])
		num2, err2 := strconv.Atoi(fields[1])

		if err1 != nil || err2 != nil {
			fmt.Println("Error parsing numbers in line:", line)
			continue
		}

		// Append the numbers to their respective columns
		column1 = append(column1, num1)
		column2 = append(column2, num2)

	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Sort the columns
	sort.Ints(column1)
	sort.Ints(column2)

	for i, v := range column1 {		
		result := math.Abs(float64(column2[i] - v))
		sum += result
		
	}
	fmt.Printf("%.4f\n", sum)
}
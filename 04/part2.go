//go:build ignore

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var keyword string = "MAS"
	var counter int = 0

	if len(keyword)%2 == 0 {
		fmt.Println("Keyword length must be odd")
		return
	}

	var middle_keyword int = len(keyword) / 2

	matrix := read_input_file("input")

	overlap_map := make([][]int, len(matrix))
	for i := range matrix {
		overlap_map[i] = make([]int, len(matrix[0]))
	}

	directions := [4][2]int{
		{1, 1},   // diagonal down-right
		{-1, 1},  // diagonal up-right
		{1, -1},  // diagonal down-left
		{-1, -1}, // diagonal up-left
	}

	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[0]); col++ {
			for dir_idx := 0; dir_idx < len(directions); dir_idx++ {
				if word_search_single_direction(matrix, row, col, keyword, 0, directions[dir_idx][0], directions[dir_idx][1]) {
					/*
					 * My rational here is that if you count the number of times the
					 * center of the keyword is part of a word search result you are able to find where
					 * a X shaped MAS exists (== two overlaps of diagonal results).
					 */
					a_y := row + (middle_keyword)*directions[dir_idx][0]
					a_x := col + (middle_keyword)*directions[dir_idx][1]

					if a_y < len(matrix) && a_x < len(matrix[0]) && a_y >= 0 && a_x >= 0 {
						overlap_map[a_y][a_x] += 1
					}
				}
			}
		}
	}

	for i := range overlap_map {
		for j := range overlap_map[i] {
			if overlap_map[i][j] == 2 {
				counter++
			}
		}
	}

	fmt.Println(counter)
}

func word_search_single_direction(matrix [][]byte, row int, column int, keyword string, keyword_idx int, direction_y int, direction_x int) bool {

	if row >= len(matrix) || column >= len(matrix[0]) || row < 0 || column < 0 {
		return false
	}

	if keyword_idx >= len(keyword) {
		return false
	}

	if matrix[row][column] != keyword[keyword_idx] {
		return false
	}

	if keyword_idx+1 == len(keyword) {
		return true
	}

	return word_search_single_direction(matrix, row+direction_y, column+direction_x, keyword, keyword_idx+1, direction_y, direction_x)
}

func read_input_file(file_path string) [][]byte {
	var matrix [][]byte

	file, err := os.Open(file_path)
	if err != nil {
		fmt.Println("Could not find input file")
	}

	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {
		/*
		 * Reading the file was tricky for two reasons:
		 *
		 * 1. **Slices and `Bytes()` behavior:**
		 *    The `Bytes()` method returns a slice containing the scanned data. In Go,
		 *    a slice is a data structure with a pointer to an underlying array, along
		 *    with metadata such as length and capacity. Each time `Scan()` is called,
		 *    the underlying memory of the slice may change. While `append()` copies the
		 *    slice contents, it does **not** copy the underlying memory, which can cause
		 *    issues when handling data across multiple scan calls.
		 *
		 * 2. **The `copy()` function's requirements:**
		 *    The `copy()` function requires the destination slice to have enough allocated
		 *    space to hold the data being copied. To address this, I used a temporary variable
		 *    to store the slice returned by `Bytes()`. This allowed me to create a new slice
		 *    with the appropriate length, ensuring that `copy()` could safely copy the data
		 *    without exceeding the available memory.
		 */
		scanned_bytes := scanner.Bytes()
		var line []byte = make([]byte, len(scanned_bytes))
		copy(line, scanner.Bytes())
		matrix = append(matrix, line)
		i++
	}

	return matrix
}

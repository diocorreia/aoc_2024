//go:build ignore

package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	var rules map[int][]int
	var list [][]int
	var checksum int = 0

	rules, list = read_file("input")

	for i := range list {
		if !check_updated_pages(list[i], rules) {
			reorder_update(list[i], rules)
			checksum += list[i][len(list[i])/2]
		}
	}

	fmt.Printf("The sum of the middle pages of the corrected updates is %v\n", checksum)
}

func read_file(file_path string) (map[int][]int, [][]int) {
	var rules map[int][]int = make(map[int][]int)
	var sequences [][]int

	file, err := os.Open(file_path)

	if err != nil {
		fmt.Println("Did not found file")
		return rules, sequences
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		values := strings.Split(line, "|")

		page, _ := strconv.Atoi(values[0])
		rule, _ := strconv.Atoi(values[1])

		rules[page] = append(rules[page], rule)
	}

	var sequence_idx = 0
	for scanner.Scan() {
		line := scanner.Text()

		values := strings.Split(line, ",")

		sequence := []int{}
		for idx := range values {
			tmp, _ := strconv.Atoi(values[idx])
			sequence = append(sequence, tmp)
		}

		sequences = append(sequences, sequence)

		sequence_idx++
	}

	return rules, sequences
}

func check_rule(list []int, list_idx int, value int) bool {
	if list_idx < 0 {
		return true
	}

	if list[list_idx] == value {
		return false
	}

	return check_rule(list, list_idx-1, value)
}

func check_page(list []int, list_idx int, rules map[int][]int) bool {
	value_in_list := list[list_idx]
	for i := range rules[value_in_list] {
		if !check_rule(list, list_idx, rules[value_in_list][i]) {
			return false
		}
	}

	return true
}

func check_updated_pages(list []int, rules map[int][]int) bool {
	for i := range list {
		if !check_page(list, i, rules) {
			return false
		}
	}
	return true
}

func reorder_update(list []int, rules map[int][]int) {
	slices.SortFunc(list, func(a int, b int) int {
		for _, before := range rules[a] {
			if b == before {
				return -1
			}
		}

		return 1
	})
}

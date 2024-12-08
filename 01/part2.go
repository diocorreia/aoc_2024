//go:build ignore

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var list1 []int
	var list2 []int
	var similarity_score int

	list1, list2 = read_two_lists_from_file("input")
	similarity_score = calculate_similarity_score(list1, list2)

	fmt.Printf("The similarity score between the two lists is: %v\n", similarity_score)
}

func read_two_lists_from_file(file_path string) ([]int, []int) {
	var list1 []int
	var list2 []int

	file, err := os.Open(file_path)

	if err != nil {
		fmt.Println("Could not find input file")
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		l1_num, _ := strconv.Atoi(line[0])
		l2_num, _ := strconv.Atoi(line[1])
		list1 = append(list1, l1_num)
		list2 = append(list2, l2_num)
	}

	return list1, list2
}

func find_occurences_in_list(list []int, number int) int {
	var occurences int = 0
	for _, val := range list {
		if val == number {
			occurences++
		}
	}
	return occurences
}

func calculate_similarity_score(l1 []int, l2 []int) int {
	var similarity_score int = 0

	for _, val := range l1 {
		occurences := find_occurences_in_list(l2, val)
		similarity_score += (occurences * val)
	}

	return similarity_score
}

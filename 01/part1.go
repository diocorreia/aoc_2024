//go:build ignore

package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var list1 []int
	var list2 []int
	var total_distance int

	list1, list2 = read_two_lists_from_file("input")
	total_distance, _ = calculate_total_distance(list1, list2)

	fmt.Printf("The total distance between lists is %v\n", total_distance)
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

func calculate_total_distance(l1 []int, l2 []int) (int, error) {

	var total_distance int = 0

	if len(l1) != len(l2) {
		return 0, errors.New("Lists have different sizes")
	}

	var l1_len int = len(l1)
	for i := 0; i < l1_len; i++ {
		var l1_min_idx int = get_min_idx(l1)
		var l2_min_idx int = get_min_idx(l2)

		total_distance += calculate_distance(l1[l1_min_idx], l2[l2_min_idx])

		remove_elemt(&l1, l1_min_idx)
		remove_elemt(&l2, l2_min_idx)
	}

	return total_distance, nil
}

func calculate_distance(val1 int, val2 int) int {
	if val1 > val2 {
		return val1 - val2
	} else {
		return val2 - val1
	}
}

func get_min_idx(list []int) int {
	var min_pos int = 0

	for i := 0; i < len(list); i++ {
		if list[min_pos] > list[i] {
			min_pos = i
		}
	}

	return min_pos
}

func remove_elemt(list *[]int, pos int) {
	var new_list = append((*list)[:pos], (*list)[pos+1:]...)
	*list = new_list
}

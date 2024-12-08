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
	var list_of_reports [][]int
	var safe_score int = 0

	list_of_reports = get_reports_from_file("input")

	for i := range list_of_reports {
		is_safe := is_report_safe(list_of_reports[i])
		if is_safe {
			safe_score++
		}
	}

	fmt.Printf("%v reports are safe\n", safe_score)

}

func get_reports_from_file(input_path string) [][]int {
	var list_of_reports [][]int
	var report_idx int = 0

	file, err := os.Open(input_path)
	if err != nil {
		fmt.Println("Could not find the file")
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		var report []int
		var level int

		line := strings.Fields(scanner.Text())

		for i := 0; i < len(line); i++ {
			level, _ = strconv.Atoi(line[i])
			report = append(report, level)
		}

		list_of_reports = append(list_of_reports, report)
		report_idx++
	}

	return list_of_reports
}

func is_report_safe(report []int) bool {
	var last_elem int = report[0]
	var gradient []int

	for i := 1; i < len(report); i++ {
		gradient = append(gradient, (report[i] - last_elem))
		last_elem = report[i]
	}

	last_elem = gradient[0]

	if check_gradient_validity(gradient[0]) == false {
		return false
	}

	for i := 1; i < len(gradient); i++ {

		if check_gradient_validity(gradient[i]) == false {
			return false
		}

		// Check if both are positive
		if last_elem > 0 && gradient[i] < 0 {
			return false
		}

		// Check if both are negative
		if last_elem < 0 && gradient[i] > 0 {
			return false
		}

		last_elem = gradient[i]
	}

	return true
}

func check_gradient_validity(gradient int) bool {
	// Check if there was no change
	if gradient == 0 {
		return false
	}

	// Check bounderies
	if gradient > 3 || gradient < -3 {
		return false
	}

	return true
}

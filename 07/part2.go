//go:build ignore

package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {

	var calibration_result int = 0

	a := read_file("input")

	for result, operands := range a {
		if is_operation_valid(result, operands, 0, 0, "+") {
			calibration_result += result
		}
	}

	fmt.Printf("The calibration result is %v\n", calibration_result)
}

func read_file(file_path string) map[int][]int {
	var equations map[int][]int = make(map[int][]int)

	file, err := os.Open(file_path)

	if err != nil {
		fmt.Println("Did not found file")
		return equations
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		re := regexp.MustCompile(`(\d+)`)
		numbers := re.FindAllString(line, -1)

		if len(numbers) <= 0 {
			return equations
		}

		result, _ := strconv.Atoi(numbers[0])

		for i := 1; i < len(numbers); i++ {
			number, _ := strconv.Atoi(numbers[i])
			equations[result] = append(equations[result], number)
		}
	}

	return equations
}

func is_operation_valid(result int, operands []int, operand_idx int, carry int, operator string) bool {
	var plus_is_valid, times_is_valid, concatenate_is_valid bool
	var next_carry int

	if len(operands) == operand_idx {
		if carry == result {
			return true
		} else {
			return false
		}
	}

	if operator == "+" {
		next_carry = carry + operands[operand_idx]
	}

	if operator == "*" {
		next_carry = carry * operands[operand_idx]
	}

	if operator == "||" {
		carry_str := strconv.Itoa(carry)
		operand_str := strconv.Itoa(operands[operand_idx])

		next_carry, _ = strconv.Atoi(carry_str + operand_str)
	}

	plus_is_valid = is_operation_valid(result, operands, operand_idx+1, next_carry, "+")
	times_is_valid = is_operation_valid(result, operands, operand_idx+1, next_carry, "*")
	concatenate_is_valid = is_operation_valid(result, operands, operand_idx+1, next_carry, "||")

	return plus_is_valid || times_is_valid || concatenate_is_valid
}

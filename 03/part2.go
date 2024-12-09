//go:build ignore

package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type command struct {
	operation string
	val1      int
	val2      int
}

func main() {
	var input string
	var cleaned_up string
	var list_of_commands []string
	var res int = 0

	input = read_input_file("input")

	cleaned_up = ignore_donts(input)

	list_of_commands = parse_commands_in_string(cleaned_up)

	for command_idx := range list_of_commands {
		command := parse_command(list_of_commands[command_idx])
		res += exec_command(command)
	}

	fmt.Printf("The result of this operation is %v\n", res)
}

func read_input_file(file_path string) string {
	var converted string

	file, err := os.Open(file_path)
	if err != nil {
		fmt.Println("Could not find input file")
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		converted += scanner.Text()
	}

	return converted
}

func parse_commands_in_string(input string) []string {
	re := regexp.MustCompile(`mul\(\d*,\d*\)`)
	return re.FindAllString(input, -1)
}

func parse_command(input string) command {
	var command command
	re := regexp.MustCompile(`mul\((\d*),(\d*)\)`)
	parsed := re.FindAllStringSubmatch(input, -1)
	command.operation = "mul"
	command.val1, _ = strconv.Atoi(parsed[0][1])
	command.val2, _ = strconv.Atoi(parsed[0][2])
	return command
}

func exec_command(command command) int {
	if command.operation == "mul" {
		return command.val1 * command.val2
	}
	return 0
}

func ignore_donts(input string) string {
	re := regexp.MustCompile(`don't\(\).*?do\(\)`)
	temp_string := re.ReplaceAllString(input, "")
	re = regexp.MustCompile(`don't\(\).*`)
	return re.ReplaceAllString(temp_string, "")
}

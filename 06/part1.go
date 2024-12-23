//go:build ignore

package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

type coordinates struct {
	y int
	x int
}

type guard struct {
	position  coordinates
	direction coordinates
	history   []coordinates
}

func main() {
	lab, guard, _ := read_map("test_input")
	move_guard(lab, &guard)
	number_of_distinct_positions := count_distinct_position(guard.history)
	fmt.Printf("The guard was in found in %v distinct positions", number_of_distinct_positions)
}

func read_map(filepath string) ([][]rune, guard, error) {
	var input_map [][]rune
	var guard guard
	var i int

	file, err := os.Open(filepath)
	if err != nil {
		return input_map, guard, errors.New("Could not find a file")
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		tmp_line := scanner.Text()
		input_map = append(input_map, []rune(tmp_line))
		i++
	}

	guard = find_guard(input_map)

	return input_map, guard, nil
}

func find_guard(input [][]rune) guard {
	for i := range input {
		for j := range input[i] {
			current := input[i][j]
			switch current {
			case 'v':
				return guard{position: coordinates{x: j, y: i}, direction: coordinates{x: 0, y: 1}}
			case '^':
				return guard{position: coordinates{x: j, y: i}, direction: coordinates{x: 0, y: -1}}
			case '>':
				return guard{position: coordinates{x: j, y: i}, direction: coordinates{x: 1, y: 0}}
			case '<':
				return guard{position: coordinates{x: j, y: i}, direction: coordinates{x: -1, y: 0}}
			}
		}
	}

	return guard{position: coordinates{x: -1, y: -1}, direction: coordinates{x: 0, y: 0}}
}

func move_guard(lab_map [][]rune, guard *guard) {
	if guard.position.y+guard.direction.y < 0 || guard.position.y+guard.direction.y >= len(lab_map) {
		guard.history = append(guard.history, coordinates{x: guard.position.x, y: guard.position.y})
		return
	}

	if guard.position.x+guard.direction.x < 0 || guard.position.x+guard.direction.x >= len(lab_map[guard.position.y]) {
		guard.history = append(guard.history, coordinates{x: guard.position.x, y: guard.position.y})
		return
	}

	if lab_map[guard.position.y+guard.direction.y][guard.position.x+guard.direction.x] == '#' {
		rotate_guard_90deg(guard)
		move_guard(lab_map, guard)
		return
	}

	guard.history = append(guard.history, coordinates{x: guard.position.x, y: guard.position.y})
	guard.position = coordinates{y: guard.position.y + guard.direction.y, x: guard.position.x + guard.direction.x}

	move_guard(lab_map, guard)
}

func rotate_guard_90deg(guard *guard) {
	directions_list := []coordinates{
		{x: 1, y: 0},
		{x: 0, y: 1},
		{x: -1, y: 0},
		{x: 0, y: -1},
	}

	for i := range directions_list {
		if guard.direction == directions_list[i] {
			if i+1 == len(directions_list) {
				guard.direction = directions_list[0]
				return
			}
			guard.direction = directions_list[i+1]
			return
		}
	}
}

func count_distinct_position(list_of_positions []coordinates) int {
	var frequency = make(map[coordinates]int)

	for i := range list_of_positions {
		frequency[list_of_positions[i]] += 1
	}

	return len(frequency)
}

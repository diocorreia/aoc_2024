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
	position        coordinates
	direction       coordinates
	history         []coordinates
	start_direction coordinates
	frequency       map[coordinates]int
}

func main() {
	lab, guard, _ := read_map("input")
	var loop_counter int = 0

	for y := range lab {
		for x := range lab[y] {
			reset_guard(&guard)

			obstacle := coordinates{y: y, x: x}
			if !is_position_free(lab, obstacle) {
				continue
			}

			move_guard(lab, &guard, obstacle)

			if check_loop(&guard) {
				loop_counter++
			}
		}
	}

	fmt.Printf("There were %v detected\n", loop_counter)
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
	guard.start_direction = guard.direction
	guard.frequency = make(map[coordinates]int)

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

func move_guard(lab_map [][]rune, guard *guard, obstacle coordinates) {
	next_position := coordinates{y: guard.position.y + guard.direction.y, x: guard.position.x + guard.direction.x}

	if next_position.y < 0 || next_position.y >= len(lab_map) {
		guard.history = append(guard.history, coordinates{x: guard.position.x, y: guard.position.y})
		return
	}

	if next_position.x < 0 || next_position.x >= len(lab_map[guard.position.y]) {
		guard.history = append(guard.history, coordinates{x: guard.position.x, y: guard.position.y})
		return
	}

	if check_for_obstacle(lab_map, next_position, obstacle) {
		rotate_guard_90deg(guard)
		move_guard(lab_map, guard, obstacle)
		return
	}

	guard.history = append(guard.history, coordinates{x: guard.position.x, y: guard.position.y})
	guard.position = next_position
	guard.frequency[guard.position] += 1

	if guard.frequency[guard.position] >= 5 {
		return
	}

	move_guard(lab_map, guard, obstacle)
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

func check_for_obstacle(lab_map [][]rune, position coordinates, additional_obstacle coordinates) bool {
	if lab_map[position.y][position.x] == '#' {
		return true
	}

	if position == additional_obstacle {
		return true
	}

	return false
}

func check_loop(guard *guard) bool {
	var loop_threshold int = 5

	if len(guard.history) < 1 {
		return false
	}

	for i := range guard.frequency {
		if guard.frequency[i] >= loop_threshold {
			return true
		}
	}

	return false
}

func reset_guard(guard *guard) {
	guard.direction = guard.start_direction
	if len(guard.history) > 0 {
		guard.position = guard.history[0]
	}
	guard.history = nil
	guard.frequency = make(map[coordinates]int)
}

func is_position_free(lab_map [][]rune, position coordinates) bool {
	if lab_map[position.y][position.x] == '#' {
		return false
	}
	if lab_map[position.y][position.x] == '>' {
		return false
	}
	if lab_map[position.y][position.x] == '^' {
		return false
	}
	if lab_map[position.y][position.x] == '<' {
		return false
	}
	if lab_map[position.y][position.x] == 'v' {
		return false
	}
	return true
}

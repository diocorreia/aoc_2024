//go:build ignore

package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

type coord struct {
	x int
	y int
}

func main() {
	var antenna_positions map[byte][]coord
	var antinodes []coord
	var map_width, map_length int

	input := read_file("input")
	map_width = len(input)
	map_length = len(input[0])

	antenna_positions = find_antenna_positions(input)

	for _, list_of_antennas := range antenna_positions {
		for _, antenna_a := range list_of_antennas {
			for _, antenna_b := range list_of_antennas {
				if antenna_a == antenna_b {
					continue
				}
				res := calculate_antinodes(antenna_a, antenna_b)
				antinodes = slices.Concat(antinodes, res)
			}
		}
	}

	antinodes = clean_positions_outside_map(antinodes, map_width, map_length)
	antinodes = remove_duplicated_positions(antinodes)

	fmt.Printf("A total of %v unique antinode positions were detected on the map.", len(antinodes))
}

func calculate_antinodes(antenna_a coord, antenna_b coord) []coord {
	var antinodes []coord = []coord{}

	x_distance := antenna_b.x - antenna_a.x
	y_distance := antenna_b.y - antenna_a.y

	antinodes = append(antinodes, coord{antenna_b.x + x_distance, antenna_b.y + y_distance})
	antinodes = append(antinodes, coord{antenna_a.x - x_distance, antenna_a.y - y_distance})

	return antinodes
}

func find_antenna_positions(input_map []string) map[byte][]coord {
	var antenna_positions map[byte][]coord = make(map[byte][]coord)
	for y := range input_map {
		for x := range input_map[y] {
			var symbol byte = input_map[y][x]
			if symbol != '.' {
				antenna_positions[symbol] = append(antenna_positions[symbol], coord{y: y, x: x})
			}
		}
	}
	return antenna_positions
}

func clean_positions_outside_map(positions []coord, y_max int, x_max int) []coord {
	return slices.DeleteFunc(positions, func(position coord) bool {
		if position.x < 0 || position.y < 0 {
			return true
		}
		if position.y >= y_max || position.x >= x_max {
			return true
		}
		return false
	})
}

func remove_antennas_out_of_list(positions []coord, antennas []coord) []coord {
	return slices.DeleteFunc(positions, func(position coord) bool {
		for _, antenna := range antennas {
			if antenna == position {
				return true
			}
		}
		return false
	})
}

func remove_duplicated_positions(positions []coord) []coord {
	var frequency_list map[coord]int = make(map[coord]int)
	var cleaned_positions_array []coord

	for _, position := range positions {
		frequency_list[position]++
	}

	for position := range frequency_list {
		cleaned_positions_array = append(cleaned_positions_array, position)
	}
	return cleaned_positions_array
}

func read_file(file_path string) []string {
	var lines []string

	file, err := os.Open(file_path)
	if err != nil {
		fmt.Println("Could not find input file")
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	return lines
}

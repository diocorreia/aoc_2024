# Advent of Code 2024

In this repository you will find my solutions to the puzzles in the 2024 edition of [Advent of Code](https://adventofcode.com).

Since I discovered this challange way after December 1st, I make no intentions on trying to catch up the advent calendar (if I end up catching it, the better).

I'm using this challange to learn [go](https://go.dev/), so don't expect my solutions to be good.

## How to run the solutions

First you will need to install Go in your system. For that, please follow the instructions in [go documentation](https://go.dev/doc/install).

All files in this repo start with `//go:build ignore`. This header allows us to have multiple declarations of the same functions (mostly the Main function) on the same directory. All files in this repo are meant to be compiled separatly.

To compile:
```
go build -o [output_name] [solution_file.go]
```

Or to run:
```
go run [solution_file.go]
```

Most puzzles will require an input file. To comply with licensing, these files are not included in this repository. Neither the puzzles. To have access to them, please visit [Advent of Code](https://adventofcode.com) official website.
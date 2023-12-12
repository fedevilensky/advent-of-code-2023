package main

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"
	"slices"
	"strconv"
)

var log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
	Level: slog.LevelDebug,
}))

func main() {
	file, err := os.Open("day-03/input.txt")
	if err != nil {
		log.Error("Error while opening file", "error", err)
	}
	defer file.Close()

	result := 0
	scanner := bufio.NewScanner(file)
	mat := make([]string, 0, 200)
	for scanner.Scan() {
		line := scanner.Text()
		mat = append(mat, line)
	}
	if err := scanner.Err(); err != nil {
		log.Error("Error while reading file", "erorr", err)
		panic(err)
	}

	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[i]); j++ {
			if mat[i][j] == '*' {
				adjacentNumbers := findAdjacentNumbers(mat, i, j)
				if len(adjacentNumbers) == 2 {
					result += adjacentNumbers[0] * adjacentNumbers[1]
				}
			}
		}
	}

	fmt.Println(result)
}

func isDigit(r byte) bool {
	return r >= '0' && r <= '9'
}

func findAdjacentNumbers(mat []string, i, j int) []int {
	adjacentNumbers := make([]int, 0, 8)
	for k := -1; k <= 1; k++ {
		coords := make([][2]int, 0, 8)
		if i+k < 0 || i+k >= len(mat) {
			continue
		}
		for l := -1; l <= 1; l++ {
			if j+l < 0 || j+l >= len(mat[i+k]) {
				continue
			}
			if isDigit(mat[i+k][j+l]) {
				number, coord := findNumber(mat, i+k, j+l)
				if !slices.ContainsFunc(coords, func(c [2]int) bool { return c[0] == coord[0] && c[1] == coord[1] }) {
					adjacentNumbers = append(adjacentNumbers, number)
					coords = append(coords, coord)
				}
			}
		}
	}
	return adjacentNumbers
}

func findNumber(mat []string, i, j int) (int, [2]int) {
	init, end := j, j
	for ; init >= 0 && isDigit(mat[i][init]); init-- {
	}
	init++
	for ; end < len(mat[i]) && isDigit(mat[i][end]); end++ {
	}
	number, err := strconv.Atoi(mat[i][init:end])
	if err != nil {
		log.Error("Error while converting string to int", "error", err)
	}
	return number, [2]int{init, end}
}

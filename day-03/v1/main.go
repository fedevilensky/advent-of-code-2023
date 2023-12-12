package main

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"
	"regexp"
	"strconv"
)

var log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
	Level: slog.LevelDebug,
}))

var (
	startsWithNumberRegex = regexp.MustCompile(`^[0-9]+`)
	symbolRegex           = regexp.MustCompile(`^[^0-9.]$`)
)

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
			if str := startsWithNumberRegex.FindString(mat[i][j:]); len(str) > 0 {
				if isTouchingSymbol(mat, i, j, str) {
					number, err := strconv.Atoi(str)
					if err != nil {
						log.Error("Error while converting string to int", "error", err)
					}
					result += number
				}
				j += len(str) - 1
			}
		}
	}

	fmt.Println(result)
}

func isTouchingSymbol(mat []string, i, j int, str string) bool {
	for k := -1; k <= 1; k++ {
		if i+k < 0 || i+k >= len(mat) {
			continue
		}
		for l := -1; l <= len(str); l++ {
			if j+l < 0 || j+l >= len(mat[i+k]) {
				continue
			}
			if symbolRegex.MatchString(string(mat[i+k][j+l : j+l+1])) {
				return true
			}
		}
	}

	return false
}

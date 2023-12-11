package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, err := os.Open("day-01/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	result := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		tokens := tokenize(line)
		toAdd := tokens[0] + tokens[len(tokens)-1]
		toAddInt, err := strconv.Atoi(string(toAdd))
		if err != nil {
			log.Fatal(err)
		}
		result += toAddInt
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)
}

var digit = regexp.MustCompile(`[0-9]`)

var (
	numberStrings  = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	stringToNumber = map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}
)

func prefixMatch(inputString string) (bool, bool) {
	for _, numberString := range numberStrings {
		if numberString == inputString {
			return true, true
		}
		if len(inputString) > len(numberString) {
			continue
		}
		if inputString == numberString[0:len(inputString)] {
			return true, false
		}
	}
	return false, false
}

func tokenize(input string) []string {
	var tokens []string

	for i := 0; i < len(input); i++ {
		if digit.MatchString(input[i : i+1]) {
			tokens = append(tokens, string(input[i:i+1]))
			continue
		}
		for j := i + 1; j <= len(input); j++ {
			if isPrefix, exactly := prefixMatch(input[i:j]); !isPrefix {
				break
			} else {
				if exactly {
					tokens = append(tokens, stringToNumber[input[i:j]])
					break
				}
			}
		}
	}

	return tokens
}

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

var notDigits = regexp.MustCompile(`[^0-9]`)

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
		onlyDigits := notDigits.ReplaceAll([]byte(line), []byte(""))
		toAdd := append(onlyDigits[0:1], onlyDigits[len(onlyDigits)-1])

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

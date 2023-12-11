package main

import (
	"bufio"
	"fmt"
	"log/slog"
	"math"
	"os"
	"strconv"
	"strings"
)

var log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
	Level: slog.LevelDebug,
}))

type Game struct {
	id    int
	plays []Play
}

type Play struct {
	Red   int
	Green int
	Blue  int
}

func main() {
	file, err := os.Open("day-02/input.txt")
	if err != nil {
		log.Error("Error while opening file", "error", err)
	}
	defer file.Close()

	result := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimPrefix(line, "Game ")
		game := getGame(line)
		minRed := math.MinInt
		minGreen := math.MinInt
		minBlue := math.MinInt
		for _, play := range game.plays {
			minRed = max(minRed, play.Red)
			minGreen = max(minGreen, play.Green)
			minBlue = max(minBlue, play.Blue)
		}

		result += minRed * minGreen * minBlue
	}
	if err := scanner.Err(); err != nil {
		log.Error("Error while reading file", "erorr", err)
		panic(err)
	}
	fmt.Println(result)
}

func getGame(line string) Game {
	var (
		err  error
		game Game
	)
	splitLine := strings.Split(line, ": ")
	if game.id, err = strconv.Atoi(splitLine[0]); err != nil {
		log.Error("Invalid input", "error", err, "line", line)
		panic(err)
	}
	if len(splitLine) != 2 {
		log.Error("Invalid splitLine length", "line", line, "splitLine", splitLine)
		panic("Invalid splitLine length")
	}
	playStrings := strings.Split(splitLine[1], "; ")
	for _, playString := range playStrings {
		play := getPlay(playString)
		game.plays = append(game.plays, play)
	}
	return game
}

func getPlay(playString string) Play {
	var (
		err  error
		play Play
	)
	playColors := strings.Split(playString, ", ")
	for _, playColor := range playColors {
		splitPlay := strings.Split(playColor, " ")
		if len(splitPlay) != 2 {
			log.Error("Invalid splitPlay length", "playColor", playColor, "splitPlay", splitPlay)
			panic("Invalid splitPlay length")
		}
		switch splitPlay[1] {
		case "red":
			if play.Red, err = strconv.Atoi(splitPlay[0]); err != nil {
				log.Error("Invalid red value", "error", err, "playColor", playColor, "splitPlay", splitPlay)
				panic(err)
			}
		case "green":
			if play.Green, err = strconv.Atoi(splitPlay[0]); err != nil {
				log.Error("Invalid green value", "error", err, "playColor", playColor, "splitPlay", splitPlay)
				panic(err)
			}
		case "blue":
			if play.Blue, err = strconv.Atoi(splitPlay[0]); err != nil {
				log.Error("Invalid blue value", "error", err, "playColor", playColor, "splitPlay", splitPlay)
				panic(err)
			}
		default:
			log.Error("Invalid color", "playColor", playColor, "splitPlay", splitPlay)
			panic("Invalid color")
		}
	}

	return play
}

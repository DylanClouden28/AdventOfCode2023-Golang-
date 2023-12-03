package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}

	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	str := string(data)

	lines := strings.Split(str, "\n")
	var GameCount int = 0
	powerTotal := 0
	for _, line := range lines {
		//fmt.Println(strings.ReplaceAll(strings.SplitAfter(line[5:], ":")[0], ":", ""))
		if line == "" {
			break
		}
		game, _ := strconv.Atoi(strings.ReplaceAll(strings.SplitAfter(line[5:], ":")[0], ":", ""))
		Validgame := true
		fmt.Println("Game: ", game)
		slice := strings.SplitAfter(line, ":")[1]
		//slice = strings.ReplaceAll(slice, " ", "")
		//fmt.Println(slice)
		sets := strings.Split(slice, ";")
		maxs := map[string]int{
			"red":   0,
			"blue":  0,
			"green": 0,
		}
		for _, set := range sets {
			//fmt.Println("  Set: ", set)
			totals := map[string]int{
				"red":   0,
				"blue":  0,
				"green": 0,
			}
			colors := strings.Split(set, ",")
			for _, color := range colors {
				if color[0] == ' ' {
					color = color[1:]
				}
				parts := strings.Split(color, " ")
				col := parts[1]
				col = strings.ReplaceAll(col, "\r", "")
				numberOfColor, _ := strconv.Atoi(parts[0])
				//fmt.Println("	Color:", col, "Number:", numberOfColor)
				totals[col] += numberOfColor
				if maxs[col] < numberOfColor {
					maxs[col] = numberOfColor
				}
			}

			if totals["red"] > 12 || totals["green"] > 13 || totals["blue"] > 14 {
				Validgame = false
			}
		}
		power := maxs["red"] * maxs["blue"] * maxs["green"]
		powerTotal += power
		println("	Power: ", power)
		//println("	Powertotal: ", powerTotal)
		println("	Red: ", maxs["red"], "Blue: ", maxs["blue"], "Green: ", maxs["green"])
		if Validgame {
			GameCount += game
		}
	}

	println("Game Count: ", GameCount)
	println("Power Total: ", powerTotal)
}

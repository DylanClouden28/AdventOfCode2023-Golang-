package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func trimNumber(s string, start int) string {
	length := 0
	for _, char := range s[start:] {
		if !unicode.IsDigit(char) {
			break
		}
		length++
	}
	return s[start : start+length]
}

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
	Numbers := make(map[int]map[int]int)
	validNumbers := make(map[int]int)

	lines := strings.Split(str, "\n")
	for i, line := range lines {
		for j := 0; j < len(line); j++ {
			char := rune(line[j])
			num := char - '0'
			if !(num >= 0 && num <= 9 || char == '.') {
				//fmt.Println("Char: ", string(char), i, j)
			}
			if unicode.IsDigit(char) {
				number := trimNumber(line, j)
				fmt.Println("Slice: ", number)
				if Numbers[i] == nil {
					Numbers[i] = make(map[int]int)
				}
				intNum, _ := strconv.Atoi(number)
				for iter := 0; iter < len(number); iter++ {
					Numbers[i][j+iter] = intNum

				}
				validNumbers[intNum] = 0
				j += len(number)
			}
		}
	}
	for i, line := range lines {
		for j := 0; j < len(line); j++ {
			char := rune(line[j])
			num := char - '0'
			if !(num >= 0 && num <= 9 || char == '.') {
				//fmt.Println("Char: ", string(char))
				for idx := i - 1; idx < (i - 1 + 3); idx++ {
					for idy := j - 1; idy < (j - 1 + 3); idy++ {
						val, valid := Numbers[idx][idy]
						//fmt.Println("\tChecking: ", idx, idy, " | ", val, "| ", valid)
						if valid {
							validNumbers[val] = 1
						}
					}
				}
			}
		}
	}
	total := 0
	for k, _ := range Numbers {
		if validNumbers[k] {
			total += k
		}
	}
	fmt.Println("Total: ", total)

	fmt.Println(validNumbers)
}

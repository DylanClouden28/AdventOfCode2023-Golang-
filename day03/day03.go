package main

import (
	"fmt"
	"io"
	"log"
	"os"
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
	for i, line := range lines {
		for j, char := range line {
			num := char - '0'
			if !(num >= 0 && num <= 9 || char == '.') {
				fmt.Println("Char: ", string(char), i, j)
			}
		}
	}
}

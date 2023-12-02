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
	file, err := os.Open("text.txt")
	if err != nil {
		log.Fatal(err)
	}

	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	str := string(data)

	lines := strings.Split(str, "\n")

	var total int
	for _, line := range lines {
		fmt.Println(line)
		var firstnum rune = -1
		var secondnum rune = -1
		for _, char := range line {
			num := int(char - '0')
			if num >= 0 && num < 10 {
				fmt.Println(num)
				if firstnum == -1 {
					firstnum = char
				}
				secondnum = char
			}
		}
		var combined string = string(firstnum) + string(secondnum)
		fmt.Println(combined)
		value, _ := strconv.Atoi(combined)
		total += value
	}
	fmt.Printf("Total is %d", total)
}

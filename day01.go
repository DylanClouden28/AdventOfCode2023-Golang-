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

	numbers := map[string]int{
		"one":   '1',
		"two":   '2',
		"three": '3',
		"four":  '4',
		"five":  '5',
		"six":   '6',
		"seven": '7',
		"eight": '8',
		"nine":  '9',
	}

	var total int
	for _, line := range lines {

		if line == "" {
			break
		}

		//fmt.Println(line)
		var firstnum rune = -1
		var secondnum rune = -1
		for i, char := range line {
			num := int(char - '0')
			if num >= 0 && num < 10 {
				//fmt.Println(num)
				if firstnum == -1 {
					firstnum = char
				}
				secondnum = char
			}
			end := i + 5
			if end > len(line) {
				end = len(line)
			}
			slice := line[i:end]
			//fmt.Println("Slice: ", slice)
			for word := range numbers {
				if strings.HasPrefix(slice, word) {
					if firstnum == -1 {
						firstnum = rune(numbers[word])
						//fmt.Println("First Number", firstnum)
					}
					secondnum = rune(numbers[word])
				}
			}
		}
		//fmt.Println("First Number:", string(firstnum), " | Second Number:", string(secondnum))
		var combined string = string(firstnum) + string(secondnum)
		//fmt.Println("	Combined", combined)
		value, _ := strconv.Atoi(combined)
		//resultingNumbers2 = append(resultingNumbers2, value)
		total += value

	}
	fmt.Printf("Total is %d\n", total)

	// var resultingNumbers1 []int
	// var resultingNumbers2 []int

	// for k, line := range lines {

	// 	if line == "" {
	// 		break
	// 	}

	// 	//fmt.Println(line)
	// 	var firstnum rune = -1
	// 	var secondnum rune = -1
	// 	for i, char := range line {
	// 		num := int(char - '0')
	// 		if num >= 0 && num < 10 {
	// 			//fmt.Println(num)
	// 			if firstnum == -1 {
	// 				firstnum = char
	// 			}
	// 			secondnum = char
	// 		}
	// 		end := i + 5
	// 		if end > len(line) {
	// 			end = len(line)
	// 		}

	// 		slice := line[i:end]
	// 		if k == 84 {
	// 			fmt.Println("Slice: ", slice)
	// 		}
	// 		var found bool = false
	// 		for key := range numbers {
	// 			if strings.Contains(slice, key) {
	// 				found = true
	// 				if k == 84 {
	// 					fmt.Printf("Found %s: %d\n", key, numbers[key])
	// 				}
	// 			}
	// 			if found {
	// 				if firstnum == -1 {
	// 					firstnum = rune(numbers[key])
	// 					if k == 84 {
	// 						fmt.Println("First Number", firstnum)
	// 					}
	// 				}
	// 				secondnum = rune(numbers[key])
	// 				if k == 84 {
	// 					fmt.Println("Second Number", secondnum)
	// 				}
	// 			}
	// 			found = false
	// 		}
	// 	}
	// 	//fmt.Println("First Number:", string(firstnum), " | Second Number:", string(secondnum))
	// 	var combined string = string(firstnum) + string(secondnum)
	// 	//fmt.Println("	Combined", combined)
	// 	value, _ := strconv.Atoi(combined)
	// 	resultingNumbers1 = append(resultingNumbers1, value)
	// 	total += value

	// }
	// fmt.Printf("Total is %d\n", total)
	// total = 0

	// for index, result := range resultingNumbers1 {
	// 	if result != resultingNumbers2[index] {
	// 		fmt.Println(lines[index])
	// 		fmt.Println("Result 1: ", result, "\ndoesn't match Result 2: ", resultingNumbers2[index], "\n")
	// 	}
	// }
}

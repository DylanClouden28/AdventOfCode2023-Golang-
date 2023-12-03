/* package main

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

		fmt.Println(line)
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
			var found bool = false
			for key := range numbers {
				if strings.Contains(slice, key) {
					found = true
					//fmt.Printf("Found %s: %d\n", key, numbers[key])
				}
				if found {
					if firstnum == -1 {
						firstnum = rune(numbers[key])
						//fmt.Println("First Number", firstnum)
					}
					secondnum = rune(numbers[key])
					//fmt.Println("Second Number", secondnum)
				}
				found = false
			}
		}
		//fmt.Println("First Number:", string(firstnum), " | Second Number:", string(secondnum))
		var combined string = string(firstnum) + string(secondnum)
		fmt.Println("	Combined", combined)
		value, _ := strconv.Atoi(combined)
		total += value
	}
	fmt.Printf("Total is %d", total)
}
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("text.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var sumPart1, sumPart2 int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		valuePart1 := getCalibrationValue(line, false)
		sumPart1 += valuePart1

		valuePart2 := getCalibrationValue(line, true)
		sumPart2 += valuePart2
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Printf("Sum of Part 1: %d\n", sumPart1)
	fmt.Printf("Sum of Part 2: %d\n", sumPart2)
}

func getCalibrationValue(s string, includeWords bool) int {
	first, last := getFirstAndLastDigits(s, includeWords)
	return first*10 + last
}

func getFirstAndLastDigits(s string, includeWords bool) (int, int) {
	digits := map[string]int{
		"one": 1, "two": 2, "three": 3, "four": 4, "five": 5,
		"six": 6, "seven": 7, "eight": 8, "nine": 9,
	}

	first, last := -1, -1
	words := strings.Fields(s)

	for _, word := range words {
		for k, v := range digits {
			if strings.Contains(word, k) {
				if first == -1 {
					first = v
				} else {
					last = v
				}
				break
			}
		}
		if !includeWords && first == -1 && last == -1 {
			for _, r := range word {
				if r >= '0' && r <= '9' {
					if first == -1 {
						first, _ = strconv.Atoi(string(r))
					} else {
						last, _ = strconv.Atoi(string(r))
					}
				}
			}
		}
	}

	return first, last
}

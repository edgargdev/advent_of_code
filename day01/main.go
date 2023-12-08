package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func CalibrateValue(input string) int {
	var foundIntegers []string
	stringSlice := strings.Split(input, "")

	integerStringMap := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
		"ten":   "10",
	}

	for i, v := range stringSlice {
		if _, err := strconv.Atoi(v); err == nil {
			log.Printf("%q looks like an integer.\n", v)
			foundIntegers = append(foundIntegers, v)
		} else {
			for j := i + 1; j <= len(stringSlice); j++ {
				testString := input[i:j]
				if val, ok := integerStringMap[testString]; ok {
					log.Printf("%q looks like an integer.\n", val)
					foundIntegers = append(foundIntegers, val)
				}
			}
		}
	}
	calibratedValue := foundIntegers[0] + foundIntegers[len(foundIntegers)-1]
	if i, err := strconv.Atoi(calibratedValue); err == nil {
		return i
	}
	return 0
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	runningSum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		runningSum += CalibrateValue(scanner.Text())
	}

	fmt.Println(runningSum)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

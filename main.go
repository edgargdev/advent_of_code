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
	for _, v := range stringSlice {
		if _, err := strconv.Atoi(v); err == nil {
			fmt.Printf("%q looks like an integer.\n", v)
			foundIntegers = append(foundIntegers, v)
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

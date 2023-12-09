package day04

import (
	"log"
	"strconv"
	"strings"
)

func parseLineForScratchcard(line string) Scratchcard {
	metadata, gameNumbersString := grabMetaAndNumbers(line)
	id := line[5:len(metadata)]
	winningNumbers, playNumbers := parseGameNumbers(gameNumbersString)
	value, err := strconv.Atoi(id)
	if err != nil {
		log.Printf("Error parsing game id")
		return Scratchcard{id: 0}
	}
	return Scratchcard{id: value, winningNumbers: winningNumbers, gameNumbers: playNumbers}
}

type Scratchcard struct {
	id int
	winningNumbers []int
	gameNumbers []int
}

func grabMetaAndNumbers(line string) (string, string) {
	stringSplit := strings.Split(line, ": ")
	return stringSplit[0], stringSplit[1]
}

func parseGameNumbers(gameNumbersString string) ([]int, []int) {
	numbersString := strings.Split(gameNumbersString, " | ")
	winningNumbers := numbersString[0]
	gameNumbers := numbersString[1]
	log.Printf("winning: %v\ngame: %v", winningNumbers, gameNumbers)
	return parseScratchCardNumbers(winningNumbers), parseScratchCardNumbers(gameNumbers)
}

func parseScratchCardNumbers(numbersString string) []int {
	var numbers []int
	for i := 0; i < len(numbersString)-1; i+=3 {
		stringNumber := strings.Trim(numbersString[i:i+2], " ")
		possibleNumber, err := strconv.Atoi(stringNumber)
		log.Printf("String Number: %v || Possible number: %v", stringNumber, possibleNumber)
		if err == nil {
			numbers = append(numbers, possibleNumber)
		}
	}
	return numbers
}

package day04

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

func GetPointsFromAllScratchcards(inputFileName string) {
	file, err := os.Open(inputFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var scratchcards []Scratchcard
	runningSum := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() != "" {
			scratchcard := parseLineForScratchcard(scanner.Text())
			scratchcards = append(scratchcards, scratchcard)
			winners := scratchcard.getWinners()
			if len(winners) != 0 {
				runningSum+= int(math.Pow(2,float64(len(winners)-1)))
			}
		}
	}

	allCopies := retrieveScratchcardCopies(scratchcards, scratchcards)

	fmt.Println("Part 1 answer")
	fmt.Println(runningSum)

	fmt.Println("Part 2 answer")
	fmt.Println(allCopies)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func retrieveScratchcardCopies(scratchcards []Scratchcard, originalList []Scratchcard) int {
	winnerMap := make([]int, len(scratchcards))
	sum := 0
	for i, scratchcard := range scratchcards {
		winnerMap[i]++
		winners := scratchcard.getWinners()
		for j, _ := range winners {
			winnerMap[i+j+1] += winnerMap[i]
		}
		sum += winnerMap[i]
	}
	return sum
}

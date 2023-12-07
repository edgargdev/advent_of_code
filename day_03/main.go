package day03

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func FindPartNumbersAdjacentToASymbol(inputFileName string) int {
	file, err := os.Open(inputFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	var gridNumbers []GridNumber
	var grid [][]string
	runningSum := 0
	lineNumber := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		grid = append(grid, strings.Split(scanner.Text(), ""))
		if scanner.Text() != "" {
			gridNumbers = append(gridNumbers, ParseLineForGridNumbers(scanner.Text(), lineNumber)...)
			lineNumber++
		}
	}

	log.Printf("Grid Numbers\n%v\n", gridNumbers)	

	validPartNumber := false
	for _,number := range gridNumbers {
		for j := number.initialPosition; j <= number.endPosition; j++ {
			i := number.lineNumber
			for x := i-1;  x <= i+1; x++ {
				for y := j-1; y <= j+1; y++ {
					if x!= i || y != j {
						if x >= 0  && x < len(grid) && y >= 0 && y < len(grid[x]) {
							_, err := strconv.Atoi(grid[x][y])
							if grid[x][y] !=  "." && err != nil {
								validPartNumber = true
							}
						}
					}
				}
			}
		}
		if validPartNumber {
			runningSum += number.value
		}
		validPartNumber = false
	}
	log.Println("Part 1 answer")
	log.Println(runningSum)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return runningSum
}

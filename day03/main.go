package day03

import (
	"bufio"
	"log"
	"os"
	"reflect"
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

	validPartNumber := false
	for _, number := range gridNumbers {
		for j := number.initialPosition; j <= number.endPosition; j++ {
			i := number.lineNumber
			for x := i - 1; x <= i+1; x++ {
				for y := j - 1; y <= j+1; y++ {
					if x != i || y != j {
						if x >= 0 && x < len(grid) && y >= 0 && y < len(grid[x]) {
							_, err := strconv.Atoi(grid[x][y])
							if grid[x][y] != "." && err != nil {
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

	gearRatioSum := 0
	for x, lineNumber := range grid {
		for y, character := range lineNumber {
			if character == "*" {
				gearRatioSum += getGearRatio(x, y, grid, gridNumbers)
			}
		}
	}

	log.Println("Part 1 answer")
	log.Println(runningSum)

	log.Println("Part 2 answer")
	log.Println(gearRatioSum)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return runningSum
}

func getGearRatio(i int, j int, grid [][]string, gridNumbers []GridNumber) int {
	var foundNumbers []GridNumber
	// log.Printf("GridNumbers: %v\n", gridNumbers)
	for xx := i - 1; xx <= i+1; xx++ {
		for yy := j - 1; yy <= j+1; yy++ {
			// log.Printf("CHECKING COORDS: %v,%v\n", xx,yy)
			if xx >= 0 && xx < len(grid) && yy >= 0 && yy < len(grid[xx]) {
				for _, gridNumber := range gridNumbers {
					// log.Printf("Checking GridNumber: %v\n", gridNumber)
					tagThisNumberAsFound := false
					for c := gridNumber.initialPosition; c <= gridNumber.endPosition; c++ {
						// log.Printf("Checking %v,%v against GNcoords: %v,%v\n", xx,yy,gridNumber.lineNumber, c)
						if c == yy && gridNumber.lineNumber == xx {
							// log.Printf("tagging this number as found %v\n", gridNumber)
							tagThisNumberAsFound = true
						}
					}
					if tagThisNumberAsFound && notDuplicateGridNumber(foundNumbers, gridNumber) {
						foundNumbers = append(foundNumbers, gridNumber)
					}
				}
			}
		}
	}
	// log.Printf("NUMBERS FOUND %v\n", foundNumbers)
	if len(foundNumbers) == 2 {
		return foundNumbers[0].value * foundNumbers[1].value
	}
	return 0
}

func notDuplicateGridNumber(foundNumbers []GridNumber, gridNumber GridNumber) bool {
	for _, foundNumber := range foundNumbers {
		if reflect.DeepEqual(foundNumber, gridNumber) {
			return false
		}
	}
	return true
}

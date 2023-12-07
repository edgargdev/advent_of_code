package day03

import (
	"bufio"
	"log"
	"os"
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
		// fmt.Println(number)
		for i := number.initialPosition; i <= number.endPosition; i++ {
			if i == number.initialPosition {
				//CHECK the following
				// * *
				// * i
				// * *
				for y := i-1; y <= i ; y++ {
					for x := number.lineNumber-1; x <= number.lineNumber+1; x++ {
						if y >= 0 && y < len(grid[number.lineNumber]) && x >= 0 && x < len(grid){
							if x != number.lineNumber || y != i {
								// log.Printf("I:Checking coords %v,%v for Number %v", x,y,number)
								if grid[x][y] != "." {
									validPartNumber = true
									// log.Printf("I FOUND SYMBOL i:%q for number %v on line %v", grid[x][y], number, number.lineNumber)
								}
							}
						}
					}
				}
			} else if i == number.endPosition {
				//CHECK the following
				//  * *
				//  i *
				//  * *
				for y := i; y <= i+1 ; y++ {
					for x := number.lineNumber-1; x <= number.lineNumber+1; x++ {
						// log.Printf("CORDS %v,%v", x, y)
						if y >= 0 && y < len(grid[number.lineNumber]) && x >= 0 && x < len(grid){
							if x != number.lineNumber || y != i {
								// log.Printf("E:Checking coords %v,%v for Number %v", k,j,number)
								if grid[x][y] != "." {
									validPartNumber = true
									// log.Printf("I FOUND SYMBOL e:%q for number %v on line %v", grid[x][y], number, number.lineNumber)
								}
							}
						}
					}
				}
			} else {
				//CHECK the following
				//  * 
				//  i
				//  *
				x := number.lineNumber
				if x-1 >= 0 {
					// log.Printf("MU:Checking coords %v,%v for Number %v", k-1,i,number)
					if grid[x-1][i] != "." {
						validPartNumber = true
						// log.Printf("I FOUND SYMBOL mu:%q for number %v on line %v", grid[x-1][i], number, number.lineNumber)
					}
				}
				if x+1 < len(grid) {
					// log.Printf("MB:Checking coords %v,%v for Number %v", k+1,i,number)
					if grid[x+1][i] != "." {
						validPartNumber = true
						// log.Printf("I FOUND SYMBOL mb:%q for number %v on line %v", grid[x+1][i], number, number.lineNumber)
					}
				}
			}
		}

		if validPartNumber {
			// log.Printf("Adding GridNumber %v\n", number)
			runningSum += number.value
		} else {
			log.Printf("Skipping GridNumber %v\n", number)
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

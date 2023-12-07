package day03

import (
	// "log"
	"strconv"
)

func ParseLineForGridNumbers(line string, lineNumber int) []GridNumber {
	// log.Printf("Line: %v", line)
	tempIntString := ""
	initialPosition := 69
	var gridNumbers []GridNumber
	for i, char := range line {
		currentStringChararcter := string(char)
		if _, err := strconv.Atoi(currentStringChararcter); err == nil {
			// log.Printf("Looking at character %q and tempInt is %v", currentStringChararcter, tempIntString)
			if tempIntString == "" {
				initialPosition = i
			}
			tempIntString += currentStringChararcter
			
		}
		
		if i >= len(line)-1 && tempIntString != "" {
			gridNumberInt, err := strconv.Atoi(tempIntString)
			if err == nil {
				gridNumber := GridNumber{value: gridNumberInt, initialPosition: initialPosition, endPosition: i, lineNumber: lineNumber}
				// log.Printf("E:Creating Grid Number %v", gridNumber)
				gridNumbers = append(gridNumbers, gridNumber)
			}
		}
		if i < len(line)-1 {
			if _, err := strconv.Atoi(string(line[i+1])); err != nil {
				if tempIntString != "" {
					gridNumberInt, err := strconv.Atoi(tempIntString)
					if err == nil {
						gridNumber := GridNumber{value: gridNumberInt, initialPosition: initialPosition, endPosition: i, lineNumber: lineNumber}
						// log.Printf("A:Creating Grid Number %v", gridNumber)
						gridNumbers = append(gridNumbers, gridNumber)
					}
				}
				tempIntString = ""
			}
		}
	}
	// log.Printf("GridNumbers %v", gridNumbers)
	// log.Printf("count: %v", len(gridNumbers))
	return gridNumbers
}

type GridNumber struct {
	value           int
	initialPosition int
	endPosition int
	lineNumber int
}

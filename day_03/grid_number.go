package day03

import (
	// "log"
	"strconv"
)

func ParseLineForGridNumbers(line string, lineNumber int) []GridNumber {
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
		} else {
			if tempIntString != "" {
				gridNumberInt, err := strconv.Atoi(tempIntString)
				if err == nil {
					gridNumbers = append(gridNumbers, GridNumber{value: gridNumberInt, initialPosition: initialPosition, endPosition: i-1, lineNumber: lineNumber})
				}
			}
			tempIntString = ""
		}
	}
	return gridNumbers
}

type GridNumber struct {
	value           int
	initialPosition int
	endPosition int
	lineNumber int
}

package day02

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	// runningSum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		// runningSum += CalibrateValue(scanner.Text())
	}

	// fmt.Println(runningSum)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

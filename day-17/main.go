package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type targetArea struct {
	initialX int
	finalX   int
	initialY int
	finalY   int
}

func readInput() (target targetArea) {
	// f, err := os.Open("example-input")
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var xStr string
	var yStr string

	scanner.Scan()
	_, err = fmt.Sscanf(scanner.Text(), "target area: %s %s", &xStr, &yStr)
	if err != nil {
		log.Fatal(err)
	}

	// Remove the "x=" and "," parts
	xFields := xStr[2:]
	xFields = xFields[:len(xFields)-1]

	// Remove the "y=" part
	yFields := yStr[2:]

	xValuesStr := strings.Split(xFields, "..")
	target.initialX, err = strconv.Atoi(xValuesStr[0])
	if err != nil {
		log.Fatal(err)
	}

	target.finalX, err = strconv.Atoi(xValuesStr[1])
	if err != nil {
		log.Fatal(err)
	}

	yValuesStr := strings.Split(yFields, "..")

	target.initialY, err = strconv.Atoi(yValuesStr[0])
	if err != nil {
		log.Fatal(err)
	}

	target.finalY, err = strconv.Atoi(yValuesStr[1])
	if err != nil {
		log.Fatal(err)
	}

	return target
}

func main() {
	// Add location (line and file) of log calls
	log.SetFlags(log.Lshortfile)

	target := readInput()
	fmt.Println("target", target)

	firstPart(target)
}

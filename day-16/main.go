package main

import (
	"bufio"
	"log"
	"os"
)

// Some "hand parsed inputs":
//
// example-input-05:
//
// 011 000 1 00000000010
//    000 000 0 000000000010110
//        000 100 0 1010
//        101 100 0 1011
//    001 000 1 00000000010
//        000 100 0 1100
//        011 100 0 1101
//        00

func readInput() (message string) {
	// f, err := os.Open("example-input-15-part-2")
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		message = scanner.Text()
	}

	return message
}

func main() {
	// Add location (line and file) of log calls
	log.SetFlags(log.Lshortfile)

	message := readInput()
	firstPart(message)
	secondPart(message)
}

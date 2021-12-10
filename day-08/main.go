package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type signalNote struct {
	signals [10]string
	digits  [4]string
}

func readInput() (notes []signalNote) {
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		newNote := signalNote{}

		_, err := fmt.Sscanf(scanner.Text(), "%s %s %s %s %s %s %s %s %s %s | %s %s %s %s",
			&newNote.signals[0], &newNote.signals[1], &newNote.signals[2],
			&newNote.signals[3], &newNote.signals[4], &newNote.signals[5],
			&newNote.signals[6], &newNote.signals[7], &newNote.signals[8],
			&newNote.signals[9],
			&newNote.digits[0], &newNote.digits[1], &newNote.digits[2],
			&newNote.digits[3],
		)

		notes = append(notes, newNote)

		if err != nil {
			log.Fatal(err)
		}
	}

	return notes
}

func main() {
	entries := readInput()
	firstPart(entries)
	secondPart(entries)
}

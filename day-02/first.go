package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func firstPart() {
	fmt.Println("Starting part 1")
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var direction string
	var intensity int

	var horizontal_pos, depth int

	for scanner.Scan() {
		_, err := fmt.Sscanf(scanner.Text(), "%s %d", &direction, &intensity)
		if err != nil {
			log.Fatal(err)
		}

		horizontal_pos = horizontal_pos + intensity*joystick[direction][1]
		depth = depth + intensity*joystick[direction][0]
	}

	fmt.Println("horizontal_pos", horizontal_pos)
	fmt.Println("depth", depth)
	fmt.Println("horizontal_pos * depth", horizontal_pos*depth)
}

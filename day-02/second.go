package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func second_part() {
	fmt.Println("Starting part 2")

	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var direction string
	var intensity int

	var horizontal_pos, depth, aim int

	for scanner.Scan() {
		_, err := fmt.Sscanf(scanner.Text(), "%s %d", &direction, &intensity)
		if err != nil {
			log.Fatal(err)
		}

		horizontal_pos += intensity * joystick[direction][1]

		aim += intensity * joystick[direction][0]
		depth += aim * intensity * joystick[direction][1]
	}

	fmt.Println("horizontal_pos", horizontal_pos)
	fmt.Println("depth", depth)
	fmt.Println("horizontal_pos * depth", horizontal_pos*depth)
}

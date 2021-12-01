package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("../input")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	// read the first separately
	scanner.Scan()
	curr, err := strconv.Atoi(scanner.Text())
	if err != nil {
		log.Fatal(err)
	}

	prev := curr

	inc := 0

	for scanner.Scan() {
		curr, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}

		if curr > prev {
			inc++
		}

		prev = curr
	}

	fmt.Println(inc)
}

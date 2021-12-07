package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func firstPart() {
	fmt.Println("===================")
	fmt.Println("Running first part")
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	// var gammaRateBits, epsilonRateBits []int16

	// Okay, so hear me out, I'm really hopping that there's never
	// going to be an equal amount of bits ones and zeros, cause what
	// I'm going to do is to increase the value in the bit position if
	// there is a one and decrease it if there is a zero.
	//
	// If the resulting position is positive, the bit 1 is the most
	// common in that position. If the result is negative, then it's a
	// 0. If the result is zero, then the problem is wrong ¯\_(ツ)_/¯
	// and I am just sit in the corner thinking about my life
	// decisions.
	//
	// Thanks for coming to my TED talk.
	var mostCommonBits []int

	for scanner.Scan() {
		diagnosticReport := scanner.Text()

		for i, diagnosticBit := range diagnosticReport {
			var counterInc int

			switch diagnosticBit {
			case '0':
				counterInc = -1
			case '1':
				counterInc = 1
			}

			if i > len(mostCommonBits)-1 {
				mostCommonBits = append(mostCommonBits, counterInc)
			} else {
				mostCommonBits[i] += counterInc
			}
		}
	}

	// Transform the counters into binary
	for i, bit := range mostCommonBits {
		if bit > 0 {
			mostCommonBits[i] = 1
		} else if bit < 0 {
			mostCommonBits[i] = 0
		} else {
			log.Fatal("Undefined behavior, just cry in the corner")
		}
	}

	gamma := convertBitArrayToDec(mostCommonBits)
	epsilon := convertBitArrayToDec(invertBitArray(mostCommonBits))

	fmt.Println("gamma", gamma)
	fmt.Println("epsilon", epsilon)
	fmt.Println("gamma * epsilon", gamma*epsilon)
}

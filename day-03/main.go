package main

import (
	"math"
)

func bitStringToDec(bitString string) int {
	var tmpArray []int
	for _, bitStr := range bitString {
		tmpArray = append(tmpArray, int(bitStr-'0'))
	}
	return convertBitArrayToDec(tmpArray)
}

// We could use some package to do this, but where's the fun in that?
func convertBitArrayToDec(bitArray []int) int {
	result := 0

	// powers of two
	j := 0

	for i := len(bitArray) - 1; i >= 0; i-- {
		if bitArray[i] == 1 {
			result += int(math.Pow(2, float64(j)))
		}
		j++
	}

	return result
}

func invertBitArray(bitArray []int) []int {
	for i, bit := range bitArray {
		bitArray[i] = bit ^ 1
	}
	return bitArray
}

func main() {
	firstPart()
	secondPart()
}

package main

import (
	"encoding/hex"
	"fmt"
	"log"
)

type generalPacket struct {
	// 3 bits
	version byte
	// 3 bits
	typeID byte
	// literal value or operator
	content interface{}
}

// type ID == 4
//
// Padded with zeros, grouped into blocks of 4 with a leading
// indicator of 1 if it not the last block or 0 if it is
type contentLiteralValue int

// Any type ID != 4
//
// Two types of len indicator:
//
// - 0: 15 bits indicating the total bitlen of sub-packet
// - 1: 11 bits indicating the number of sub-packet
type contentOperator struct {
	// 1 bit
	lenID lenType
	// 11 or 15 bits
	length  int
	content []generalPacket
}

type lenType uint16

const (
	LenTypeTotalBitLen     lenType = 0x0
	LenTypeNumberOfPackets lenType = 0x1
)

var LenTypeToBitLenMap = map[lenType]int{
	LenTypeTotalBitLen:     15,
	LenTypeNumberOfPackets: 11,
}

// If the received packet is not, we can pass a bitOffset to start
// reading from the correct position in the byte. (OBS: the bit offset
// is from the left / most significant bit)
func decodePacket(rawValue []byte, bitOffset int) (generalPacket, int, int) {
	var decodedMessage generalPacket

	currentBit := bitOffset
	currentByte := 0

	// get the version
	for i := 2; i >= 0; i-- {
		if isBitSet(rawValue[currentByte], currentBit) {
			decodedMessage.version |= 1 << i
		}

		currentBit, currentByte = nextBit(currentBit, currentByte)
	}

	// get the type ID
	for i := 2; i >= 0; i-- {
		if isBitSet(rawValue[currentByte], currentBit) {
			decodedMessage.typeID |= 1 << i
		}

		currentBit, currentByte = nextBit(currentBit, currentByte)
	}

	var newByteOffset int

	if decodedMessage.typeID == 4 {
		decodedMessage.content, currentBit, newByteOffset = decodeLiteralValue(rawValue[currentByte:], currentBit)
	} else {
		decodedMessage.content, currentBit, newByteOffset = decodeOperator(rawValue[currentByte:], currentBit)
	}

	currentByte += newByteOffset

	return decodedMessage, currentBit, currentByte
}

func decodeLiteralValue(rawValue []byte, bitOffset int) (contentLiteralValue, int, int) {
	var decodedValue contentLiteralValue

	currentBit := bitOffset
	currentByte := 0

	isLast := false
	for !isLast {
		// check the first bit
		if !isBitSet(rawValue[currentByte], currentBit) {
			isLast = true
		}

		currentBit, currentByte = nextBit(currentBit, currentByte)

		// read the block
		for j := 3; j >= 0; j-- {
			if isBitSet(rawValue[currentByte], currentBit) {
				decodedValue |= (1 << j)
			}
			currentBit, currentByte = nextBit(currentBit, currentByte)
		}

		if !isLast {
			decodedValue = decodedValue << 4
		}
	}

	return decodedValue, currentBit, currentByte
}

func decodeOperator(rawValue []byte, bitOffset int) (contentOperator, int, int) {
	var decodedValue contentOperator

	currentBit := bitOffset
	currentByte := 0

	// check the length type ID
	if !isBitSet(rawValue[currentByte], currentBit) {
		decodedValue.lenID = LenTypeTotalBitLen
	} else {
		decodedValue.lenID = LenTypeNumberOfPackets
	}

	currentBit, currentByte = nextBit(currentBit, currentByte)

	// read the length of the operator
	for i := LenTypeToBitLenMap[decodedValue.lenID] - 1; i >= 0; i-- {
		if isBitSet(rawValue[currentByte], currentBit) {
			decodedValue.length |= 1 << i
		}
		currentBit, currentByte = nextBit(currentBit, currentByte)
	}

	switch decodedValue.lenID {
	case LenTypeTotalBitLen:
		// get the position of the last bit / byte for this operator
		finalBit := currentBit
		finalByte := currentByte
		for i := 0; i < decodedValue.length; i++ {
			finalBit, finalByte = nextBit(finalBit, finalByte)
		}

		for currentByte <= finalByte {
			var newPacket generalPacket
			var newByteOffset int

			if currentByte == finalByte && currentBit == finalBit {
				break
			}

			newPacket, currentBit, newByteOffset = decodePacket(rawValue[currentByte:], currentBit)

			decodedValue.content = append(decodedValue.content, newPacket)

			currentByte += newByteOffset
		}
	case LenTypeNumberOfPackets:
		for i := 0; i < decodedValue.length; i++ {
			var newPacket generalPacket
			var newByteOffset int

			newPacket, currentBit, newByteOffset = decodePacket(rawValue[currentByte:], currentBit)

			decodedValue.content = append(decodedValue.content, newPacket)

			currentByte += newByteOffset
		}
	}

	return decodedValue, currentBit, currentByte
}

func nextBit(currentBit, currentByte int) (int, int) {
	currentBit++
	if currentBit > 7 {
		currentByte++
		currentBit = 0
	}
	return currentBit, currentByte
}

func isBitSet(originalByte byte, bitLeftOffset int) bool {
	return (originalByte & (1 << (7 - bitLeftOffset))) != 0
}

func sumVersions(decodedPacket generalPacket) int {
	totalSum := 0

	totalSum += int(decodedPacket.version)

	if operator, ok := decodedPacket.content.(contentOperator); ok {
		for _, subPacket := range operator.content {
			totalSum += sumVersions(subPacket)
		}
	}

	return totalSum
}

func firstPart(message string) {
	fmt.Println("===================")
	fmt.Println("Starting first part")
	fmt.Println("===================")

	rawValue, err := hex.DecodeString(message)
	if err != nil {
		log.Panic(err)
	}

	packet, _, _ := decodePacket(rawValue, 0)

	fmt.Println("Version sum:", sumVersions(packet))
}

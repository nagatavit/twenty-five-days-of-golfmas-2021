package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"math"
)

const (
	OperationSum         byte = 0x0
	OperationProduct     byte = 0x1
	OperationMinimum     byte = 0x2
	OperationMaximum     byte = 0x3
	LiteralValue         byte = 0x4
	OperationGreaterThan byte = 0x5
	OperationLessThan    byte = 0x6
	OperationEqualTo     byte = 0x7
)

func executeOperation(packet generalPacket) (result int) {
	if packet.typeID == LiteralValue {
		literalVal, ok := packet.content.(contentLiteralValue)
		if !ok {
			log.Fatal("erm... idk how did this happened")
		}
		return int(literalVal)
	}

	operator, ok := packet.content.(contentOperator)
	if !ok {
		log.Fatal("erm... idk how did this happened either")
	}

	switch packet.typeID {
	case OperationSum:
		sum := 0
		for _, packet := range operator.content {
			sum += executeOperation(packet)
		}
		return sum

	case OperationProduct:
		prod := 1
		for _, packet := range operator.content {
			prod *= executeOperation(packet)
		}
		return prod

	case OperationMinimum:
		min := math.MaxInt
		for _, packet := range operator.content {
			newCandidate := executeOperation(packet)
			if newCandidate < min {
				min = newCandidate
			}
		}
		return min

	case OperationMaximum:
		max := 0
		for _, packet := range operator.content {
			newCandidate := executeOperation(packet)
			if newCandidate > max {
				max = newCandidate
			}
		}
		return max

	case OperationGreaterThan:
		firstOperand := executeOperation(operator.content[0])
		secondOperand := executeOperation(operator.content[1])
		if firstOperand > secondOperand {
			return 1
		} else {
			return 0
		}

	case OperationLessThan:
		firstOperand := executeOperation(operator.content[0])
		secondOperand := executeOperation(operator.content[1])
		if firstOperand < secondOperand {
			return 1
		} else {
			return 0
		}

	case OperationEqualTo:
		firstOperand := executeOperation(operator.content[0])
		secondOperand := executeOperation(operator.content[1])
		if firstOperand == secondOperand {
			return 1
		} else {
			return 0
		}

	default:
		log.Fatal("unknown operation")
	}

	return result
}

func secondPart(message string) {
	fmt.Println("====================")
	fmt.Println("Starting second part")
	fmt.Println("====================")

	rawValue, err := hex.DecodeString(message)
	if err != nil {
		log.Panic(err)
	}

	packet, _, _ := decodePacket(rawValue, 0)

	fmt.Println("Resulting operation", executeOperation(packet))
}

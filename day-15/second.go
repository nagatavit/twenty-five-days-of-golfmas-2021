package main

import (
	"fmt"
)

func secondPart(firstMap [][]int) {
	// pre-allocate the real map size /////////////////////////////////////////
	colLen := 5 * len(firstMap)
	lineLen := 5 * len(firstMap[0])

	realSizeMap := make([][]int, colLen)

	for i, _ := range realSizeMap {
		realSizeMap[i] = make([]int, lineLen)
	}

	// actually fill it //////////////////////////////////////////////////////
	for i := 0; i < lineLen; i++ {
		for j := 0; j < colLen; j++ {
			newValue := firstMap[i%len(firstMap)][j%len(firstMap[0])] + (i/len(firstMap) + j/len(firstMap[0]))
			if newValue > 9 {
				newValue = 1
			}
			realSizeMap[i][j] = newValue
		}

		for _, a := range realSizeMap[i] {
			fmt.Printf("%d", a)
		}

		break
	}

	// for i := 0; i < 5; i++ {
	//	for j := 0; j < 5; j++ {
	//		for firstMaplineIdx, line := range firstMap {
	//			for firstMapColumnIdx, val := range line {
	//				realSizeMap[i*firstMaplineIdx][j*firstMapColumnIdx] = val
	//			}
	//		}
	//	}
	// }

}

package main

func abs(x int) int {
	if x >= 0 {
		return x
	} else {
		return -x
	}
}

func max(x1, x2 int) int {
	if x1 >= x2 {
		return x1
	} else {
		return x2
	}
}

func min(x1, x2 int) int {
	if x1 <= x2 {
		return x1
	} else {
		return x2
	}
}

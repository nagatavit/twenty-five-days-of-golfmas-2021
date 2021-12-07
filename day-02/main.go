package main

// First part:
// int slice representing [depth, horizontal pos]
//
// Second part:
// int slice representing [aim, horizontal pos / depth increment]
var joystick = map[string][]int{
	"forward": {0, 1},
	"down":    {1, 0},
	"up":      {-1, 0},
}

func main() {
	firstPart()

	secondPart()
}

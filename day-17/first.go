package main

import (
	"fmt"
	"sort"
)

type Probe struct {
	x    int
	y    int
	xVel int
	yVel int
}

func (probe *Probe) Step(target targetArea) (isInTarget bool, passedTarget bool) {
	probe.x = probe.x + probe.xVel
	probe.y = probe.y + probe.yVel

	if probe.xVel > 0 {
		probe.xVel = probe.xVel - 1
	} else if probe.xVel < 0 {
		probe.xVel = probe.xVel - 1
	}

	probe.yVel = probe.yVel - 1

	if probe.x >= min(target.finalX, target.initialX) && probe.x <= max(target.finalX, target.initialX) &&
		probe.y >= min(target.finalY, target.initialY) && probe.y <= max(target.finalY, target.initialY) {
		return true, false
	}

	if probe.y < min(target.initialY, target.finalY) {
		return false, true
	}

	return false, false
}

func firstPart(target targetArea) {
	fmt.Println("===================")
	fmt.Println("Starting first part")
	fmt.Println("===================")

	var probe Probe

	var highestPeakCandidates []int

	// As we are aiming for the highest y value, we are going to
	// search the initial y velocities from 0 to the y position
	// further from the axis.
	//
	// The reason for this is, with a velocity greater than the
	// target, the probe will always skip it.
	for velY := -max(abs(target.initialY), abs(target.finalY)); velY <= max(abs(target.initialY), abs(target.finalY)); velY++ {
	L:
		for velX := -1000; velX < 1000; velX++ {
			probe.x = 0
			probe.y = 0
			probe.xVel = velX
			probe.yVel = velY

			highestPeak := 0

			for {
				isInTarget, passedTarget := probe.Step(target)

				if probe.y > highestPeak {
					highestPeak = probe.y
				}

				if isInTarget {
					highestPeakCandidates = append(highestPeakCandidates, highestPeak)
					continue L
				} else if passedTarget {
					continue L
				}
			}
		}
	}

	sort.Ints(highestPeakCandidates)

	fmt.Println(highestPeakCandidates)

	fmt.Println("====================")
	fmt.Println("Starting second part")
	fmt.Println("====================")

	fmt.Println(len(highestPeakCandidates))
}

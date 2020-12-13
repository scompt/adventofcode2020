package main

import "fmt"

func main() {

	buses := map[int]int{
		23:  0,
		41:  13,
		37:  17,
		479: 23,
		13:  36,
		17:  40,
		29:  52,
		373: 54,
		19:  73,
	}

	orderedBuses := []int{479, 373, 41, 37, 29, 23, 19, 17, 13}

	busIndex := 0
	bus := orderedBuses[busIndex]
	busOffset := buses[bus]
	longestSequence := 0

	for i := 0; ; {
		anchor := i - busOffset
		fmt.Printf("%d\n", anchor)
		yay, sequenceLength := woot(orderedBuses, buses, anchor)
		if yay {
			fmt.Printf("%d\n", anchor)
			return
		} else if sequenceLength > longestSequence {
			longestSequence = sequenceLength
			bus = 1
			for j := 0; j < longestSequence; j++ {
				bus *= orderedBuses[j]
			}
		}
		i += bus
	}
}

func woot(orderedBuses []int, buses map[int]int, anchor int) (bool, int) {
	for i, bus2 := range orderedBuses {
		if (anchor+buses[bus2])%bus2 != 0 {

			//fmt.Printf("%d\n", bus2)
			return false, i
		}
	}
	return true, 0
}

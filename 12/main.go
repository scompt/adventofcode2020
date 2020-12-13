package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type location struct {
	x int
	y int
}

func east(loc location, orient int, value int) (location, int) {
	loc.x -= value
	return loc, orient
}
func west(loc location, orient int, value int) (location, int) {
	loc.x += value
	return loc, orient
}
func south(loc location, orient int, value int) (location, int) {
	loc.y -= value
	return loc, orient
}
func north(loc location, orient int, value int) (location, int) {
	loc.y += value
	return loc, orient
}
func left(loc location, orient int, value int) (location, int) {
	orient = (orient + 360 - value) % 360
	return loc, orient
}
func right(loc location, orient int, value int) (location, int) {
	orient = (orient + 360 + value) % 360
	return loc, orient
}
func forward(loc location, orient int, value int) (location, int) {
	if orient == 0 {
		return east(loc, orient, value)
	} else if orient == 180 {
		return west(loc, orient, value)
	} else if orient == 270 {
		return north(loc, orient, value)
	} else if orient == 90 {
		return south(loc, orient, value)
	}
	panic(orient)
}

func abs(in int) int {
	if in < 0 {
		return -in
	}
	return in
}

func sail(reader *bufio.Reader) (location, int) {
	scanner := bufio.NewScanner(reader)
	loc := location{}
	orientation := 0

	for {
		if !scanner.Scan() {
			if scanner.Err() == nil {
				break
			}
			panic(scanner.Err())
		}
		token := scanner.Text()
		command := string(token[0])
		value, err := strconv.Atoi(token[1:])
		if err != nil {
			panic(err)
		}

		switch command {
		case "N":
			loc, orientation = north(loc, orientation, value)

		case "S":
			loc, orientation = south(loc, orientation, value)

		case "E":
			loc, orientation = east(loc, orientation, value)

		case "W":
			loc, orientation = west(loc, orientation, value)

		case "L":
			loc, orientation = left(loc, orientation, value)

		case "R":
			loc, orientation = right(loc, orientation, value)

		case "F":
			loc, orientation = forward(loc, orientation, value)

		default:
			panic(command)
		}

	}

	return loc, abs(loc.x) + abs(loc.y)
}
func main() {
	reader := bufio.NewReader(os.Stdin)
	_, distance := sail(reader)
	fmt.Printf("%d\n", distance)
}

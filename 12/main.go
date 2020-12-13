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

func east(shipLoc location, waypointLoc location, orient int, value int) (location, location, int) {
	waypointLoc.x += value
	return shipLoc, waypointLoc, orient
}
func west(shipLoc location, waypointLoc location, orient int, value int) (location, location, int) {
	waypointLoc.x -= value
	return shipLoc, waypointLoc, orient
}
func south(shipLoc location, waypointLoc location, orient int, value int) (location, location, int) {
	waypointLoc.y -= value
	return shipLoc, waypointLoc, orient
}
func north(shipLoc location, waypointLoc location, orient int, value int) (location, location, int) {
	waypointLoc.y += value
	return shipLoc, waypointLoc, orient
}
func left(shipLoc location, waypointLoc location, orient int, value int) (location, location, int) {
	if value == 90 {
		waypointLoc.x, waypointLoc.y = -waypointLoc.y, waypointLoc.x

	} else if value == 180 {
		waypointLoc.x, waypointLoc.y = -waypointLoc.x, -waypointLoc.y

	} else if value == 270 {
		waypointLoc.x, waypointLoc.y = waypointLoc.y, -waypointLoc.x
	} else {
		panic("asdf")
	}
	return shipLoc, waypointLoc, orient
}
func right(shipLoc location, waypointLoc location, orient int, value int) (location, location, int) {
	if value == 90 {
		waypointLoc.x, waypointLoc.y = waypointLoc.y, -waypointLoc.x

	} else if value == 180 {
		waypointLoc.x, waypointLoc.y = -waypointLoc.x, -waypointLoc.y

	} else if value == 270 {
		waypointLoc.x, waypointLoc.y = -waypointLoc.y, waypointLoc.x
	} else {
		panic("asdf")
	}
	return shipLoc, waypointLoc, orient
}
func forward(shipLoc location, waypointLoc location, orient int, value int) (location, location, int) {
	shipLoc.x += value * waypointLoc.x
	shipLoc.y += value * waypointLoc.y
	return shipLoc, waypointLoc, orient
}

func abs(in int) int {
	if in < 0 {
		return -in
	}
	return in
}

func sail(reader *bufio.Reader) (location, int) {
	scanner := bufio.NewScanner(reader)
	shipLoc := location{}
	waypointLoc := location{
		x: 10,
		y: 1,
	}
	orientation := 0

	for {
		fmt.Printf("%+v %+v\n", shipLoc, waypointLoc)
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
			shipLoc, waypointLoc, orientation = north(shipLoc, waypointLoc, orientation, value)

		case "S":
			shipLoc, waypointLoc, orientation = south(shipLoc, waypointLoc, orientation, value)

		case "E":
			shipLoc, waypointLoc, orientation = east(shipLoc, waypointLoc, orientation, value)

		case "W":
			shipLoc, waypointLoc, orientation = west(shipLoc, waypointLoc, orientation, value)

		case "L":
			shipLoc, waypointLoc, orientation = left(shipLoc, waypointLoc, orientation, value)

		case "R":
			shipLoc, waypointLoc, orientation = right(shipLoc, waypointLoc, orientation, value)

		case "F":
			shipLoc, waypointLoc, orientation = forward(shipLoc, waypointLoc, orientation, value)

		default:
			panic(command)
		}

	}

	return shipLoc, abs(shipLoc.x) + abs(shipLoc.y)
}
func main() {
	reader := bufio.NewReader(os.Stdin)
	_, distance := sail(reader)
	fmt.Printf("%d\n", distance)
}

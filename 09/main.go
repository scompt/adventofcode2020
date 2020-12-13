package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strconv"
)

func readInt(scanner *bufio.Scanner) (int, error) {
	if !scanner.Scan() {
		return 0, fmt.Errorf("scanner.Err(): %+v", scanner.Err())
	}
	strVal := scanner.Text()
	intVal, err := strconv.Atoi(strVal)
	if err != nil {
		return 0, err
	}
	return intVal, nil
}

func readPreamble(scanner *bufio.Scanner, windowSize int) (*list.List, *map[int]struct{}, error) {
	windowList := list.New()
	windowSet := make(map[int]struct{})

	for i := 0; i < windowSize; i++ {
		intVal, err := readInt(scanner)
		if err != nil {
			return nil, nil, err
		}
		windowList.PushBack(intVal)
		windowSet[intVal] = struct{}{}
	}

	return windowList, &windowSet, nil
}

func findFirstFailure(scanner *bufio.Scanner, windowList *list.List, windowSetPtr *map[int]struct{}) (int, error) {
	windowSet := *windowSetPtr
	for {
	top:
		for e := windowList.Front(); e != nil; e = e.Next() {
			fmt.Printf("  % 5d, ", e.Value.(int))
		}
		fmt.Printf("\n")

		for k := range windowSet {
			fmt.Printf("  % 5d, ", k)
		}
		fmt.Printf("\n")

		anInt, err := readInt(scanner)
		fmt.Printf("Read %d, ", anInt)
		if err != nil {
			return 0, err
		}

		for e := windowList.Front(); e != nil; e = e.Next() {
			windowInt := e.Value.(int)
			seekingInt := anInt - windowInt

			_, found := windowSet[seekingInt]
			fmt.Printf("Saw %d, looking for %d, found %t\n", windowInt, seekingInt, found)

			if seekingInt == windowInt {
				continue
			}

			if found {
				//fmt.Printf("L1: %d\n", windowList.Len())
				frontInt := windowList.Remove(windowList.Front()).(int)
				//fmt.Printf("L2: %d\n", windowList.Len())
				windowList.PushBack(anInt)
				//fmt.Printf("L3: %d\n", windowList.Len())

				delete(windowSet, frontInt)
				windowSet[anInt] = struct{}{}

				goto top
			}

		}
		return anInt, nil

	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	windowList, windowSet, err := readPreamble(scanner, 25)
	if err != nil {
		panic(err)
	}
	firstFailure, err := findFirstFailure(scanner, windowList, windowSet)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%d\n", firstFailure)
}

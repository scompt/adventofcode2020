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

func readPreamble(scanner *bufio.Scanner, windowSize int) (*list.List, *map[int]struct{}, []int, error) {
	var allInts []int
	windowList := list.New()
	windowSet := make(map[int]struct{})

	for i := 0; i < windowSize; i++ {
		intVal, err := readInt(scanner)
		if err != nil {
			return nil, nil, nil, err
		}
		allInts = append(allInts, intVal)
		windowList.PushBack(intVal)
		windowSet[intVal] = struct{}{}
	}

	return windowList, &windowSet, allInts, nil
}

func findFirstFailure(scanner *bufio.Scanner, windowList *list.List, windowSetPtr *map[int]struct{}, allInts []int) (int, []int, error) {
	windowSet := *windowSetPtr
	for {
	top:
		//for e := windowList.Front(); e != nil; e = e.Next() {
		//fmt.Printf("  % 5d, ", e.Value.(int))
		//}
		//fmt.Printf("\n")

		//for k := range windowSet {
		//fmt.Printf("  % 5d, ", k)
		//}
		//fmt.Printf("\n")

		anInt, err := readInt(scanner)
		//fmt.Printf("Read %d, ", anInt)
		if err != nil {
			return 0, nil, err
		}

		for e := windowList.Front(); e != nil; e = e.Next() {
			windowInt := e.Value.(int)
			seekingInt := anInt - windowInt

			_, found := windowSet[seekingInt]
			//fmt.Printf("Saw %d, looking for %d, found %t\n", windowInt, seekingInt, found)

			if seekingInt == windowInt {
				continue
			}

			if found {
				frontInt := windowList.Remove(windowList.Front()).(int)
				windowList.PushBack(anInt)
				allInts = append(allInts, anInt)

				delete(windowSet, frontInt)
				windowSet[anInt] = struct{}{}

				goto top
			}

		}
		return anInt, allInts, nil

	}
}

func findContiguousRange(allInts []int, soughtInt int) ([]int, error) {
	//fmt.Printf("%d\n%+v\n", soughtInt, allInts)
	for i := 0; i < len(allInts); i++ {
		sum := 0

		for j := i; j < len(allInts); j++ {
			sum += allInts[j]

			if sum == soughtInt {
				return allInts[i : j+1], nil
			} else if sum > soughtInt {
				break
			}
		}
	}

	return nil, fmt.Errorf("Not found")
}

func findMinMax(ints []int) (int, int) {
	min, max := ints[0], ints[0]

	for _, anInt := range ints {
		//fmt.Printf("%d\n", anInt)
		if anInt > max {
			max = anInt
		}
		if anInt < min {
			min = anInt
		}
	}
	return min, max
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	windowList, windowSet, allInts, err := readPreamble(scanner, 25)
	if err != nil {
		panic(err)
	}
	firstFailure, allInts, err := findFirstFailure(scanner, windowList, windowSet, allInts)
	if err != nil {
		panic(err)
	}
	contiguousRange, err := findContiguousRange(allInts, firstFailure)
	if err != nil {
		panic(err)
	}
	min, max := findMinMax(contiguousRange)
	fmt.Printf("%d\n", min+max)
}

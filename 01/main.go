package main

import "bufio"
import "os"
import "io"
import "strings"
import "strconv"
import "sort"
import "fmt"

func main() {
	seekingSum, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(err)
	}

	ints := readInts(os.Stdin)
	ints.Sort()
	i, left, right, err := findBlah3(ints, seekingSum)
	if err != nil {
		panic(err)
	}

	product := ints[left] * ints[right] * ints[i]
	fmt.Printf("%d x %d x %d = %d\n", ints[left], ints[right], ints[i], product)
}

func findBlah3(inSlice sort.IntSlice, seekingSum int) (int, int, int, error) {
	for i := 0; i < len(inSlice); i++ {
		currentVal := inSlice[i]
		currentTarget := seekingSum - currentVal
		inSlice[i] = 0

		left, right, err := findBlah(inSlice, currentTarget)
		inSlice[i] = currentVal

		if err != nil {
			continue
		} else if left != i && right != i {
			return i, left, right, nil
		}
	}
	return -1, -1, -1, fmt.Errorf("Nope")
}

func findBlah(inSlice sort.IntSlice, seekingSum int) (int, int, error) {
	leftPointer := 0
	rightPointer := inSlice.Len() - 1

	if len(inSlice) == 0 {
		return leftPointer, rightPointer, fmt.Errorf("Empty slice")
	}

	for {
		if leftPointer == rightPointer {
			return leftPointer, rightPointer, fmt.Errorf("%d == %d", leftPointer, rightPointer)
		}

		currentSum := inSlice[leftPointer] + inSlice[rightPointer]

		if currentSum > seekingSum {
			rightPointer--
		} else if currentSum < seekingSum {
			leftPointer++
		} else {
			return leftPointer, rightPointer, nil
		}
	}
}

func readInts(reader io.Reader) sort.IntSlice {
	var ret []int
	r := bufio.NewReader(reader)
	for {
		line, err := r.ReadString('\n')
		if len(line) == 0 && err != nil {
			if err == io.EOF {
				break
			}
		}
		line = strings.TrimSuffix(line, "\n")
		num, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		ret = append(ret, num)
	}
	return ret
}

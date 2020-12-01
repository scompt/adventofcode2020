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
	left, right, err := findBlah(ints, seekingSum)
	if err != nil {
		panic(err)
	}

	product := ints[left] * ints[right]
	fmt.Printf("%d x %d = %d\n", left, right, product)
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

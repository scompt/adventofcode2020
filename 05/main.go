package main

import "bufio"
import "os"
import "io"
import "fmt"
import "strings"
import "strconv"

func bspToRow(input string) int {
	input = strings.ReplaceAll(input, "F", "0")
	input = strings.ReplaceAll(input, "B", "1")
	output, _ := strconv.ParseInt(input, 2, 0)
	return int(output)
}

func bspToCol(input string) int {
	input = strings.ReplaceAll(input, "L", "0")
	input = strings.ReplaceAll(input, "R", "1")
	output, _ := strconv.ParseInt(input, 2, 0)
	return int(output)
}

func bspToSeatID(input string) (int, int, int) {
	row := bspToRow(input[0:7])
	col := bspToCol(input[7:10])
	return row*8 + col, row, col
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	plane := [128][8]bool{}

	for {
		line, err := reader.ReadString('\n')
		if len(line) < 7 {
			break
		}

		_, row, col := bspToSeatID(line)
		plane[row][col] = true

		if err == io.EOF {
			break
		}
	}

	prev := false
	for i := 0; i < 128; i++ {
		for j := 0; j < 8; j++ {
			if plane[i][j] {
				fmt.Printf("*")
				prev = true
			} else if prev {

				fmt.Printf("%d", i*8+j)
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Printf("\n")
	}
}

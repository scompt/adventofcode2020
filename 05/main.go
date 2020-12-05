package main

import "math"
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

func bspToSeatID(input string) int {
	row := bspToRow(input[0:7])
	col := bspToCol(input[7:10])
	return row*8 + col
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	maxSeatID := 0.0
	for {
		line, err := reader.ReadString('\n')
		if len(line) < 7 {
			break
		}

		seatID := float64(bspToSeatID(line))
		maxSeatID = math.Max(maxSeatID, seatID)

		if err == io.EOF {
			break
		}
	}

	fmt.Printf("%d\n", int(maxSeatID))
}

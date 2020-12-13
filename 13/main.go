package main

import "fmt"

func main() {

	start := 1000417
	nums := []int{23, 41, 37, 479, 13, 17, 29, 373, 19}
	for _, num := range nums {

		for i := 0; ; i += num {
			if i > start {
				fmt.Printf("%d %d %d %d\n", num, i, i-start, (i-start)*num)
				break
			}

		}
	}
}

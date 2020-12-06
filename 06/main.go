package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type group struct {
	yesQuestions map[byte]struct{}
}

func parseGroup(reader *bufio.Reader) (*group, error) {
	yesQuestions := make(map[byte]struct{})
	for {
		for {
			b, err := reader.ReadByte()
			if b == '\n' || b == 0 {
				break
			}

			yesQuestions[b] = struct{}{}
			if err == io.EOF {
				break
			}
		}

		b, err := reader.ReadByte()
		if b == '\n' {
			return &group{yesQuestions}, nil
		}
		if err != nil {
			if len(yesQuestions) == 0 {
				return nil, err
			}
			return &group{yesQuestions}, nil
		}

		err = reader.UnreadByte()
		if err != nil {
			return nil, err
		}
	}
}

func countYesQuestions(reader *bufio.Reader) int {
	groups := []group{}

	for {
		g, err := parseGroup(reader)
		if err == nil {
			groups = append(groups, *g)
		} else {
			break
		}
	}

	count := 0
	for _, g := range groups {
		count += len(g.yesQuestions)
	}
	return count
}

func main() {
	count := countYesQuestions(bufio.NewReader(os.Stdin))
	fmt.Printf("%d\n", count)
}

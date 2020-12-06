package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type group struct {
	yesQuestions map[byte]struct{}
	personCount  int
	answerCounts map[byte][]int
}

func (g *group) countUnanimousQuestions() int {
	count := 0
	for _, answers := range g.answerCounts {
		if len(answers) == g.personCount {
			count++
		}
	}
	return count
}

func parseGroup(reader *bufio.Reader) (*group, error) {
	yesQuestions := make(map[byte]struct{})
	personCount := 0
	answerCounts := make(map[byte][]int)
	for {
		for {
			b, err := reader.ReadByte()
			if b == '\n' || b == 0 {
				personCount++
				break
			}

			yesQuestions[b] = struct{}{}
			answerCounts[b] = append(answerCounts[b], 1)
			if err == io.EOF {
				personCount++
				break
			}
		}

		b, err := reader.ReadByte()
		if b == '\n' {
			return &group{yesQuestions, personCount, answerCounts}, nil
		}
		if err != nil {
			if len(yesQuestions) == 0 {
				return nil, err
			}
			return &group{yesQuestions, personCount, answerCounts}, nil
		}

		err = reader.UnreadByte()
		if err != nil {
			return nil, err
		}
	}
}

func countUnanimousQuestions(reader *bufio.Reader) int {
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
		count += g.countUnanimousQuestions()
	}
	return count
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
	count := countUnanimousQuestions(bufio.NewReader(os.Stdin))
	fmt.Printf("%d\n", count)
}

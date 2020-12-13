package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"container/list"
)

type bag struct {
	adjective string
	color     string
}

type target struct {
	bagCount  int
	targetBag bag
}

type rule struct {
	source  bag
	targets []target
}

type edge struct {
	other  bag
	weight int
}

type graph struct {
	nodes        map[bag]struct{}
	forwardEdges map[bag]*list.List
	reverseEdges map[bag]*list.List
}

func newGraph() *graph {
	var g graph
	g.nodes = make(map[bag]struct{})
	g.forwardEdges = make(map[bag]*list.List)
	g.reverseEdges = make(map[bag]*list.List)
	return &g
}

func parseBag(scanner *bufio.Scanner, theGraph *graph) (bag, bool, error) {
	if !scanner.Scan() {
		return bag{}, false, fmt.Errorf("")
	}
	adjective := scanner.Text()
	if !scanner.Scan() {
		return bag{}, false, fmt.Errorf("")
	}
	color := scanner.Text()
	if !scanner.Scan() { // Bag(s)
		return bag{}, false, fmt.Errorf("")
	}

	bagsString := scanner.Text()
	final := bagsString[len(bagsString)-1:] == "."
	parsedBag := bag{adjective, color}
	theGraph.nodes[parsedBag] = struct{}{}
	return parsedBag, final, nil
}

func parseTargets(scanner *bufio.Scanner, source *bag, theGraph *graph) ([]target, error) {
	if !scanner.Scan() {
		return nil, fmt.Errorf("")
	}
	if scanner.Text() == "no" {
		if !scanner.Scan() { // other
			return nil, fmt.Errorf("")
		}
		if !scanner.Scan() { // bags.
			return nil, fmt.Errorf("")
		}
		return []target{}, nil
	}

	var targets []target
	for {
		bagCount, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}
		bag, final, err := parseBag(scanner, theGraph)
		targetBag := target{bagCount, bag}
		targets = append(targets, targetBag)
		if _, ok := theGraph.forwardEdges[*source]; !ok {
			theGraph.forwardEdges[*source] = list.New()
		}
		theGraph.forwardEdges[*source].PushBack(edge{bag, bagCount})
		if _, ok := theGraph.reverseEdges[bag]; !ok {
			theGraph.reverseEdges[bag] = list.New()
		}
		theGraph.reverseEdges[bag].PushBack(*source)

		if final {
			break
		} else {
			if !scanner.Scan() { // pre-read the next number
				return nil, fmt.Errorf("")
			}
		}

	}
	return targets, nil
}

func parseRule(scanner *bufio.Scanner, theGraph *graph) (rule, error) {
	bag, _, err := parseBag(scanner, theGraph)
	if err != nil {
		return rule{}, err
	}
	if !scanner.Scan() { // Contains
		return rule{}, fmt.Errorf("")
	}
	targets, err := parseTargets(scanner, &bag, theGraph)
	if err != nil {
		return rule{}, err
	}

	return rule{bag, targets}, nil
}

func parseRules(scanner *bufio.Scanner) (map[bag][]target, *graph, error) {
	theGraph := newGraph()
	rules := make(map[bag][]target)
	for {
		rule, err := parseRule(scanner, theGraph)
		if err == nil {
			rules[rule.source] = rule.targets
		} else {
			break
		}
	}
	return rules, theGraph, nil
}

func count(rules *map[bag][]target, sought *bag) int {
	targetBags := make([]target, 0)
	for _, ts := range *rules {
		for _, t := range ts {
			if t.targetBag == *sought {
				targetBags = append(targetBags, t)
			}
		}
	}
	return 0
}

// https://stackoverflow.com/a/15323988/111777
func bagInSlice(a bag, bagList *list.List) bool {
	for e := bagList.Front(); e != nil; e = e.Next() {
		if a == e.Value {
			return true
		}
	}
	return false
}

func woot(theGraph *graph, sought *bag) int {
	seen := list.New()
	open := list.New()

	open.PushFront(*sought)

	for {
		if open.Len() == 0 {
			break
		}

		current := open.Remove(open.Front()).(bag)
		fmt.Printf("current: %+v\n", current)
		if !bagInSlice(current, seen) {
			seen.PushBack(current)
			neighbors := theGraph.reverseEdges[current]
			if neighbors != nil {
				for e := neighbors.Front(); e != nil; e = e.Next() {
					b := e.Value.(bag)
					if !bagInSlice(b, open) && !bagInSlice(b, seen) {
						open.PushBack(b)
					}
				}
			}
		}
	}
	return seen.Len() - 1
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)
	_, theGraph, err := parseRules(scanner)
	fmt.Printf("%+v\n", theGraph.nodes)
	fmt.Printf("%+v\n", theGraph.forwardEdges)
	fmt.Printf("%+v\n", theGraph.reverseEdges)
	if err != nil {
		panic(err)
	}
	count := woot(theGraph, &bag{"shiny", "gold"})
	fmt.Printf("%d\n", count)
}

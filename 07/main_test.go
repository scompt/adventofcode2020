package main

import (
	"bufio"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseBag(t *testing.T) {
	tests := []struct {
		input  string
		output bag
	}{
		{"light red bags", bag{"light", "red"}},
		{"bright white bag", bag{"bright", "white"}},
		{"faded blue bags", bag{"faded", "blue"}},
	}

	for _, tt := range tests {
		theGraph := newGraph()
		reader := bufio.NewReader(strings.NewReader(tt.input))
		scanner := bufio.NewScanner(reader)
		scanner.Split(bufio.ScanWords)
		outputRule, final, err := parseBag(scanner, theGraph)
		assert.NoError(t, err, tt.input)
		assert.False(t, final)
		assert.Equal(t, tt.output, outputRule, tt.input)
	}
}
func TestParseRule(t *testing.T) {
	tests := []struct {
		input  string
		output rule
	}{
		{"light red bags contain 1 bright white bag, 2 muted yellow bags.", rule{bag{"light", "red"}, []target{target{1, bag{"bright", "white"}}, target{2, bag{"muted", "yellow"}}}}},
		//{"dark orange bags contain 3 bright white bags, 4 muted yellow bags."},
		{"bright white bags contain 1 shiny gold bag.", rule{bag{"bright", "white"}, []target{target{1, bag{"shiny", "gold"}}}}},
		//{"muted yellow bags contain 2 shiny gold bags, 9 faded blue bags."},
		//{"shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags."},
		//{"dark olive bags contain 3 faded blue bags, 4 dotted black bags."},
		//{"vibrant plum bags contain 5 faded blue bags, 6 dotted black bags."},
		{"faded blue bags contain no other bags.", rule{bag{"faded", "blue"}, []target{}}},
		//{"dotted black bags contain no other bags."},
	}

	for _, tt := range tests {
		theGraph := graph{}
		reader := bufio.NewReader(strings.NewReader(tt.input))
		scanner := bufio.NewScanner(reader)
		scanner.Split(bufio.ScanWords)
		outputRule, err := parseRule(scanner, &theGraph)
		assert.NoError(t, err, tt.input)
		assert.Equal(t, tt.output, outputRule, tt.input)
	}
}

func TestParseRules(t *testing.T) {
	input := `light red bags contain 1 bright white bag, 2 muted yellow bags.
dark orange bags contain 3 bright white bags, 4 muted yellow bags.
bright white bags contain 1 shiny gold bag.
muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.
shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.
dark olive bags contain 3 faded blue bags, 4 dotted black bags.
vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
faded blue bags contain no other bags.
dotted black bags contain no other bags.`

	reader := bufio.NewReader(strings.NewReader(input))
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)
	_, rules, err := parseRules(scanner)

	assert.NoError(t, err)
	assert.NotNil(t, rules)
	//assert.Equal(t, 9, len(rules))
}

func TestCount(t *testing.T) {
	input := `light red bags contain 1 bright white bag, 2 muted yellow bags.
dark orange bags contain 3 bright white bags, 4 muted yellow bags.
bright white bags contain 1 shiny gold bag.
muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.
shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.
dark olive bags contain 3 faded blue bags, 4 dotted black bags.
vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
faded blue bags contain no other bags.
dotted black bags contain no other bags.`

	reader := bufio.NewReader(strings.NewReader(input))
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)
	_, rules, err := parseRules(scanner)

	assert.NoError(t, err)
	assert.NotNil(t, rules)
	//assert.Equal(t, 9, len(rules))

	//c := count(rules, bag{"shiny", "gold"})
	//assert.Equal(t, 4, c)
}

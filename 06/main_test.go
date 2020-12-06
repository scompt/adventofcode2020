package main

import (
	"bufio"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestParseGroup(t *testing.T) {
	tests := []struct {
		input            string
		yesQuestionCount int
	}{
		{"abc", 3},
		{"a\nb\nc", 3},
		{"ab\nac", 3},
		{"a\na\na\na", 1},
		{"b", 1},
	}

	for _, tt := range tests {
		reader := bufio.NewReader(strings.NewReader(tt.input))
		group, err := parseGroup(reader)
		assert.NoError(t, err, tt.input)
		assert.NotNil(t, group, tt.input)
		assert.Equal(t, tt.yesQuestionCount, len(group.yesQuestions), tt.input)
	}
}

func TestCountUnanimousQuestions(t *testing.T) {
	tests := []struct {
		input            string
		personCount      int
		yesQuestionCount int
	}{
		{"abc", 1, 3},
		{"a\nb\nc", 3, 0},
		{"ab\nac", 2, 1},
		{"a\na\na\na", 4, 1},
		{"b", 1, 1},
	}

	for _, tt := range tests {
		reader := bufio.NewReader(strings.NewReader(tt.input))
		group, err := parseGroup(reader)
		assert.NoError(t, err, tt.input)
		assert.NotNil(t, group, tt.input)
		assert.Equal(t, tt.personCount, group.personCount, tt.input)
		assert.Equal(t, tt.yesQuestionCount, group.countUnanimousQuestions(), tt.input)
	}
}

func TestCountYesQuestions(t *testing.T) {
	tests := []struct {
		input            string
		yesQuestionCount int
	}{
		{"abc\n\na\nb\nc\n\nab\nac\n\na\na\na\na\n\nb", 11},
	}

	for _, tt := range tests {
		reader := bufio.NewReader(strings.NewReader(tt.input))
		count := countYesQuestions(reader)
		assert.Equal(t, tt.yesQuestionCount, count)
	}
}

func TestCountUnanimousQuestions2(t *testing.T) {
	tests := []struct {
		input                  string
		unanimousQuestionCount int
	}{
		{"abc\n\na\nb\nc\n\nab\nac\n\na\na\na\na\n\nb", 6},
	}

	for _, tt := range tests {
		reader := bufio.NewReader(strings.NewReader(tt.input))
		count := countUnanimousQuestions(reader)
		assert.Equal(t, tt.unanimousQuestionCount, count)
	}
}

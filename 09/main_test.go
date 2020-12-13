package main

import (
	"bufio"
	"container/list"
	"github.com/stretchr/testify/assert"
	"strconv"
	"strings"
	"testing"
)

func TestReadPreamble(t *testing.T) {
	tests := []struct {
		input        string
		preambleSize int
		errorCase    bool
		nextInt      int
	}{
		{"35\n20\n15\n25\n47\n40\n62\n55\n65\n95\n102\n117\n150\n182\n127\n219\n299\n277\n309\n576", 3, false, 25},
		{"35\n20\n15\n25", 7, true, 0},
		{"35\n20\n15\n25", 2, false, 15},
		{"35\nwoot", 2, true, 0},
	}

	for _, tt := range tests {
		reader := bufio.NewReader(strings.NewReader(tt.input))
		scanner := bufio.NewScanner(reader)
		windowList, windowSet, allInts, err := readPreamble(scanner, tt.preambleSize)
		if tt.errorCase {
			assert.Error(t, err, tt.input)
			assert.Nil(t, windowList, tt.input)
			assert.Nil(t, windowSet, tt.input)
			assert.Nil(t, allInts, tt.input)
		} else {
			assert.NoError(t, err, tt.input)
			assert.NotNil(t, windowList, tt.input)
			assert.Equal(t, tt.preambleSize, windowList.Len(), tt.input)
			assert.Equal(t, tt.preambleSize, len(allInts), tt.input)

			assert.True(t, scanner.Scan(), tt.input)
			nextIntStr := scanner.Text()
			nextInt, err := strconv.Atoi(nextIntStr)
			assert.NoError(t, err, tt.input)
			assert.Equal(t, tt.nextInt, nextInt)

			assert.Equal(t, 35, windowList.Front().Value, tt.input)
			assert.NotNil(t, (*windowSet)[35], tt.input)
		}
	}
}

func TestLinkedList(t *testing.T) {
	l := list.List{}
	l.Init()
	assert.Equal(t, 0, l.Len())

	l.PushBack(1)
	assert.Equal(t, 1, l.Len())

	l.Remove(l.Front())
	assert.Equal(t, 0, l.Len())

}
func TestFindFirstFailure(t *testing.T) {
	tests := []struct {
		input        string
		preambleSize int
		firstFailure int
	}{
		{"35\n20\n15\n25\n47\n40\n62\n55\n65\n95\n102\n117\n150\n182\n127\n219\n299\n277\n309\n576", 5, 127},
	}

	for _, tt := range tests {
		reader := bufio.NewReader(strings.NewReader(tt.input))
		scanner := bufio.NewScanner(reader)
		windowList, windowSet, allInts, err := readPreamble(scanner, tt.preambleSize)
		assert.NoError(t, err, tt.input)
		assert.NotNil(t, windowList, tt.input)
		assert.Equal(t, tt.preambleSize, windowList.Len(), tt.input)

		firstFailure, allInts, err := findFirstFailure(scanner, windowList, windowSet, allInts)
		assert.NoError(t, err, tt.input)
		assert.Equal(t, tt.firstFailure, firstFailure)
	}
}

func TestFindContiguousRange(t *testing.T) {
	allInts := []int{35, 20, 15, 25, 47, 40, 62, 55, 65, 95, 102, 117, 150, 182, 127, 219, 299, 277, 309, 576}
	contiguousRange, err := findContiguousRange(allInts, 127)

	assert.NoError(t, err)
	assert.NotNil(t, contiguousRange)
	assert.Equal(t, 15, contiguousRange[0])
	assert.Equal(t, 40, contiguousRange[len(contiguousRange)-1])

	min, max := findMinMax(contiguousRange)
	assert.Equal(t, 15, min, contiguousRange)
	assert.Equal(t, 47, max, contiguousRange)
}

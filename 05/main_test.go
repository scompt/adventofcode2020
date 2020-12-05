package main

import "testing"
import "github.com/stretchr/testify/assert"

func TestBspToRow(t *testing.T) {
	tests := []struct {
		input  string
		output int
	}{
		{"BFFFBBF", 70},
		{"FFFBBBF", 14},
		{"BBFFBBF", 102},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.output, bspToRow(tt.input))
	}
}

func TestBspToCol(t *testing.T) {
	tests := []struct {
		input  string
		output int
	}{
		{"RRR", 7},
		{"RLL", 4},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.output, bspToCol(tt.input))
	}
}

func TestBspToSeatID(t *testing.T) {
	tests := []struct {
		input  string
		output int
	}{
		{"BFFFBBFRRR", 567},
		{"FFFBBBFRRR", 119},
		{"BBFFBBFRLL", 820},
	}

	for _, tt := range tests {
		seatID, _, _ := bspToSeatID(tt.input)
		assert.Equal(t, tt.output, seatID)
	}
}

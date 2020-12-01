package main

import "testing"
import "strings"
import "github.com/stretchr/testify/assert"
import "sort"

func TestFindBlah(t *testing.T) {
	_, _, err := findBlah([]int{}, 10)
	assert.Error(t, err)

	sli := sort.IntSlice{1, 2, 3, 4, 5}
	tests := []struct {
		seekingSum    int
		expectedLeft  int
		expectedRight int
		expectedErr   error
	}{
		{6, 0, 4, nil},
		{5, 0, 3, nil},
		{7, 1, 4, nil},
	}
	for _, tt := range tests {
		left, right, err := findBlah(sli, tt.seekingSum)
		if tt.expectedErr != nil {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
			assert.Equal(t, tt.expectedLeft, left)
			assert.Equal(t, tt.expectedRight, right)
		}
	}
}

func TestReadLines(t *testing.T) {
	assert.Panics(t, func() { readInts(strings.NewReader("woot")) })

	assert.Equal(t, sort.IntSlice{123}, readInts(strings.NewReader("123")))
	assert.Equal(t, sort.IntSlice{123, 456}, readInts(strings.NewReader("123\n456")))
}

package main

import (
	"bufio"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestSail(t *testing.T) {
	reader := bufio.NewReader(strings.NewReader("N10\nE10"))
	_, distance := sail(reader)
	assert.Equal(t, 20, distance)
}

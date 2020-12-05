package main

import "testing"
import "github.com/stretchr/testify/assert"

func TestParseBadTreeMap(t *testing.T) {
	treeMap, err := parseTreeMap("asdf")

	assert.Nil(t, treeMap)
	assert.Error(t, err)
}

func TestParseGoodTreeMap(t *testing.T) {
	input := `..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#`
	treeMap, err := parseTreeMap(input)

	assert.NotNil(t, treeMap)
	assert.NoError(t, err)
	assert.Equal(t, 11, treeMap.width)
	assert.Equal(t, 11, treeMap.height)
	assert.NotNil(t, treeMap.trees)
	assert.Equal(t, 37, len(treeMap.trees))
}

func TestCountTrees(t *testing.T) {
	input := `..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#`
	treeMap, _ := parseTreeMap(input)
	count := treeMap.countTrees(3, 1)
	assert.Equal(t, 7, count)

	count = treeMap.countTrees(1, 1)
	assert.Equal(t, 2, count)

	count = treeMap.countTrees(5, 1)
	assert.Equal(t, 3, count)

	count = treeMap.countTrees(7, 1)
	assert.Equal(t, 4, count)

	count = treeMap.countTrees(1, 2)
	assert.Equal(t, 2, count)
}

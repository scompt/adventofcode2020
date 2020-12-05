package main

import "fmt"
import "strings"
import "bufio"
import "os"
import "io"

type location struct {
	x int
	y int
}

type treeMap struct {
	width  int
	height int
	trees  map[location]struct{}
}

func (m *treeMap) String() string {
	builder := strings.Builder{}
	for y := 0; y < m.height; y++ {
		for x := 0; x < m.width; x++ {
			if _, ok := m.trees[location{x, y}]; ok {
				builder.WriteByte('#')
			} else {
				builder.WriteByte('.')
			}
		}
		builder.WriteByte('\n')
	}
	return builder.String()
}

func (m *treeMap) countTrees(right int, down int) int {
	x := 0
	y := 0
	count := 0

	for {
		if y > m.height {
			break
		}

		if _, ok := m.trees[location{x % m.width, y}]; ok {
			count++
		}

		x += right
		y += down
	}

	return count
}

func parseTreeMap(input string) (*treeMap, error) {
	width := -1
	height := -1
	trees := make(map[location]struct{})

	row := 0
	col := -1
	for i, c := range input {
		col++
		if c == '\n' {
			col = -1
			row++
			if width == -1 {
				width = i
				height = 1
			} else {
				height++
			}
		} else if c == '#' {
			trees[location{col, row}] = struct{}{}
		}
	}

	if width == -1 || height == -1 {
		return nil, fmt.Errorf("%d %d", width, height)
	}
	m := treeMap{width, height + 1, trees}

	return &m, nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	buf := new(strings.Builder)
	_, err := io.Copy(buf, reader)
	if err != nil {
		panic(err)
	}

	treeMap, err := parseTreeMap(buf.String())

	product := 1
	product *= treeMap.countTrees(1, 1)
	product *= treeMap.countTrees(3, 1)
	product *= treeMap.countTrees(5, 1)
	product *= treeMap.countTrees(7, 1)
	product *= treeMap.countTrees(1, 2)

	fmt.Printf("%d\n", product)

}

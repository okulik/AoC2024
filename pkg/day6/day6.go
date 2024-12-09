package day6

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

const (
	guardVisitedSymbol = 'X'
	obstacleSymbol     = '#'
)

type location struct{ x, y int }
type path map[string]bool
type move func(location) location

var (
	guardSymbols = "^>v<"
	moves        = []move{
		func(loc location) location { return location{loc.x, loc.y - 1} },
		func(loc location) location { return location{loc.x + 1, loc.y} },
		func(loc location) location { return location{loc.x, loc.y + 1} },
		func(loc location) location { return location{loc.x - 1, loc.y} },
	}
)

func Run() {
	file, err := os.Open("pkg/day6/input")
	if err != nil {
		panic("Can't open input file")
	}
	defer func() { _ = file.Close() }()

	gd := NewGuardDetector(file)

	count := gd.CountDistinctGuardLocations()
	fmt.Printf("Guard visited locations: %d\n", count)
}

type GuardDetector struct {
	grid  [][]byte
	guard location
	move  int
}

func NewGuardDetector(input io.Reader) *GuardDetector {
	var guard location
	var move int = -1
	grid := make([][]byte, 0)

	row := 0
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if len(line) == 0 {
			continue
		}

		gridline := make([]byte, len(line))
		for col := range line {
			gridline[col] = line[col]
			if index := strings.IndexAny(guardSymbols, string(line[col])); index != -1 {
				guard = location{col, row}
				move = index
			}
		}

		grid = append(grid, gridline)
		row++
	}

	return &GuardDetector{
		grid:  grid,
		guard: guard,
		move:  move,
	}
}

func (gd *GuardDetector) CountDistinctGuardLocations() int {
	grid := copyGrid(gd.grid)
	gd.moveGuard(grid, gd.guard, gd.move)
	return gd.countVisitedPositions(grid)
}

func (gd *GuardDetector) moveGuard(grid [][]byte, guardLocation location, moveIndex int) {
	grid[guardLocation.y][guardLocation.x] = guardVisitedSymbol

	tentativeGuardLocation := moves[moveIndex](guardLocation)
	if tentativeGuardLocation.x < 0 ||
		tentativeGuardLocation.x >= len(gd.grid[0]) ||
		tentativeGuardLocation.y < 0 ||
		tentativeGuardLocation.y >= len(gd.grid) {
		return
	}

	if grid[tentativeGuardLocation.y][tentativeGuardLocation.x] == obstacleSymbol {
		gd.moveGuard(grid, guardLocation, nextMoveIndex(moveIndex))
		return
	}

	gd.moveGuard(grid, tentativeGuardLocation, moveIndex)
}

func (gd *GuardDetector) countVisitedPositions(grid [][]byte) int {
	count := 0
	for row := range grid {
		for col := range grid[row] {
			if grid[row][col] == guardVisitedSymbol {
				count++
			}
		}
	}

	return count
}

func copyGrid(src [][]byte) [][]byte {
	dst := make([][]byte, 0, len(src))
	for _, row := range src {
		gr := make([]byte, len(row))
		copy(gr, row)
		dst = append(dst, gr)
	}

	return dst
}

func nextMoveIndex(moveIndex int) int {
	mi := moveIndex + 1
	if mi >= len(moves) {
		mi = 0
	}
	return mi
}

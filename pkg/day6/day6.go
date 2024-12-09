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

	count = gd.CountNumberOfInfiniteLoops()
	fmt.Printf("Guard infinite loops: %d\n", count)
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
	path := make(map[string]bool)
	_ = gd.moveGuard(path, gd.guard, gd.move)
	return gd.countVisitedPositions(path)
}

func (gd *GuardDetector) moveGuard(path map[string]bool, guardLocation location, moveIndex int) bool {
	tag := tag(guardLocation, moveIndex)
	if _, ok := path[tag]; ok {
		return true
	}

	path[tag] = true

	tentativeGuardLocation := moves[moveIndex](guardLocation)
	if tentativeGuardLocation.x < 0 ||
		tentativeGuardLocation.x >= len(gd.grid[0]) ||
		tentativeGuardLocation.y < 0 ||
		tentativeGuardLocation.y >= len(gd.grid) {
		return false
	}

	if gd.grid[tentativeGuardLocation.y][tentativeGuardLocation.x] == obstacleSymbol {
		return gd.moveGuard(path, guardLocation, nextMoveIndex(moveIndex))
	}

	return gd.moveGuard(path, tentativeGuardLocation, moveIndex)
}

func (gd *GuardDetector) countVisitedPositions(path map[string]bool) int {
	tags := make(map[string]bool)
	for tag := range path {
		tags[strings.Join(strings.Split(tag, "|")[:2], "|")] = true
	}
	return len(tags)
}

func (gd *GuardDetector) CountNumberOfInfiniteLoops() int {
	var tmpSymbol byte
	path := make(map[string]bool, len(gd.grid)*len(gd.grid))
	count := 0

	for y := range gd.grid {
		for x := range gd.grid[y] {
			if (gd.guard.x == x && gd.guard.y == y) || gd.grid[y][x] == obstacleSymbol {
				continue
			}
			tmpSymbol = gd.grid[y][x]
			gd.grid[y][x] = obstacleSymbol
			if gd.moveGuard(path, gd.guard, gd.move) {
				count++
			}
			clear(path)
			gd.grid[y][x] = tmpSymbol
		}
	}

	return count
}

func nextMoveIndex(moveIndex int) int {
	mi := moveIndex + 1
	if mi >= len(moves) {
		mi = 0
	}
	return mi
}

func tag(loc location, mv int) string {
	return fmt.Sprintf("%d|%d|%d", loc.x, loc.y, mv)
}

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
	poolSize           = 50
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
	_ = gd.moveGuard(path, nil)
	return gd.countVisitedPositions(path)
}

func (gd *GuardDetector) moveGuard(path map[string]bool, extraObstacle *location) bool {
	loc := gd.guard
	mv := gd.move

	for {
		tag := tag(loc, mv)
		if _, ok := path[tag]; ok {
			return true
		}

		path[tag] = true

		tentativeGuardLocation := moves[mv](loc)
		if tentativeGuardLocation.x < 0 ||
			tentativeGuardLocation.x >= len(gd.grid[0]) ||
			tentativeGuardLocation.y < 0 ||
			tentativeGuardLocation.y >= len(gd.grid) {
			return false
		}

		if gd.grid[tentativeGuardLocation.y][tentativeGuardLocation.x] == obstacleSymbol ||
			extraObstacle != nil && extraObstacle.x == tentativeGuardLocation.x && extraObstacle.y == tentativeGuardLocation.y {
			mv = nextMoveIndex(mv)
			continue
		}

		loc = tentativeGuardLocation
	}
}

func (gd *GuardDetector) countVisitedPositions(path map[string]bool) int {
	tags := make(map[string]bool)
	for tag := range path {
		tags[strings.Join(strings.Split(tag, "|")[:2], "|")] = true
	}
	return len(tags)
}

func (gd *GuardDetector) CountNumberOfInfiniteLoops() int {
	numJobs := len(gd.grid) * len(gd.grid[0])
	jobs := make(chan location, numJobs)
	results := make(chan bool, numJobs)

	for i := 0; i < poolSize; i++ {
		go func(jobs <-chan location, results chan<- bool) {
			path := make(map[string]bool, 10000)
			for loc := range jobs {
				results <- gd.moveGuard(path, &loc)
				clear(path)
			}
		}(jobs, results)
	}

	gd.gridIterator(func(x, y int) {
		jobs <- location{x, y}
	})
	close(jobs)

	count := 0
	gd.gridIterator(func(_, _ int) {
		if <-results {
			count++
		}
	})

	return count
}

func (gd *GuardDetector) gridIterator(mapFn func(x, y int)) {
	for y := range gd.grid {
		for x := range gd.grid[y] {
			if (gd.guard.x == x && gd.guard.y == y) || gd.grid[y][x] == obstacleSymbol {
				continue
			}
			mapFn(x, y)
		}
	}
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

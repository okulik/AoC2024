package day8

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func Run() {
	file, err := os.Open("pkg/day8/input")
	if err != nil {
		panic("Can't open input file")
	}
	defer func() { _ = file.Close() }()

	af := NewAntennaFinder(file)

	count := af.CountUniqueAntinodes()
	fmt.Printf("Total count of antinodes is %d\n", count)
}

type location struct {
	x, y int
}

type AntennaFinder struct {
	grid     [][]byte
	antennas map[byte][]location
}

func NewAntennaFinder(input io.Reader) *AntennaFinder {
	grid := make([][]byte, 0)
	antennas := make(map[byte][]location)

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
			chr := byte(line[col])
			if isAntenna(chr) {
				if antennas[chr] == nil {
					antennas[chr] = []location{}
				}
				antennas[chr] = append(antennas[chr], location{x: col, y: row})
			}
		}

		grid = append(grid, gridline)
		row++
	}

	return &AntennaFinder{
		grid:     grid,
		antennas: antennas,
	}
}

func (af *AntennaFinder) CountUniqueAntinodes() int {
	for _, locs := range af.antennas {
		antennaPairs := generateAntennaPairs(locs)
		for _, pair := range antennaPairs {
			af.generateAndSaveAntinodes(pair...)
		}
	}

	count := 0
	for row := range af.grid {
		for col := range af.grid[row] {
			if af.grid[row][col] == byte('#') {
				count++
			}
		}
	}

	return count
}

func (af *AntennaFinder) generateAndSaveAntinodes(loc ...location) (location, location) {
	an1, an2 := generateAntinodes(loc[0], loc[1])
	af.storeAntinode(an1)
	af.storeAntinode(an2)
	return location{}, location{}
}

func (af *AntennaFinder) storeAntinode(an location) {
	if af.outOfGrid(an) {
		return
	}

	af.grid[an.y][an.x] = byte('#')
}

func (af *AntennaFinder) outOfGrid(loc location) bool {
	return loc.y < 0 || loc.y >= len(af.grid) || loc.x < 0 || loc.x >= len(af.grid[0])
}

func generateAntennaPairs(locs []location) [][]location {
	pairs := [][]location{}
	for i := 0; i < len(locs)-1; i++ {
		for j := i + 1; j < len(locs); j++ {
			pairs = append(pairs, []location{locs[i], locs[j]})
		}
	}
	return pairs
}

func generateAntinodes(ant1, ant2 location) (location, location) {
	diffx := ant1.x - ant2.x
	diffy := ant1.y - ant2.y

	an1 := location{ant1.x + diffx, ant1.y + diffy}
	an2 := location{ant2.x - diffx, ant2.y - diffy}

	return an1, an2
}

func isAntenna(chr byte) bool {
	return chr >= 'a' && chr <= 'z' || chr >= 'A' && chr <= 'Z' || chr >= '0' && chr <= '9'
}

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

	count = af.CountUniqueAntinodesWithHarmonics()
	fmt.Printf("Total count of antinodes including resonant harmonics is %d\n", count)
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
	af.clearGrid()

	for _, locs := range af.antennas {
		if len(locs) <= 1 {
			continue
		}
		antennaPairs := generateAntennaPairs(locs)
		for _, pair := range antennaPairs {
			af.generateAndSaveAntinodes(pair[0], pair[1], 1)
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

func (af *AntennaFinder) CountUniqueAntinodesWithHarmonics() int {
	af.clearGrid()

	for _, locs := range af.antennas {
		if len(locs) <= 1 {
			continue
		}
		antennaPairs := generateAntennaPairs(locs)
		for _, pair := range antennaPairs {
			af.generateAndSaveAntinodes(pair[0], pair[1], 0)
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

func (af *AntennaFinder) generateAndSaveAntinodes(ant1, ant2 location, maxNodes int) {
	dx := ant1.x - ant2.x
	dy := ant1.y - ant2.y

	antinodes := []location{}
	anDiag := af.generateUnidirAntinodes(ant1, dx, dy, maxNodes, func(loc *location, dx, dy int) {
		loc.x += dx
		loc.y += dy
	})
	antinodes = append(antinodes, anDiag...)
	anDiag = af.generateUnidirAntinodes(ant2, dx, dy, maxNodes, func(loc *location, dx, dy int) {
		loc.x -= dx
		loc.y -= dy
	})
	antinodes = append(antinodes, anDiag...)

	for _, an := range antinodes {
		af.grid[an.y][an.x] = byte('#')
	}

	if maxNodes == 0 {
		af.grid[ant1.y][ant1.x] = byte('#')
		af.grid[ant2.y][ant2.x] = byte('#')
	}
}

func (af *AntennaFinder) generateUnidirAntinodes(antenna location, dx, dy, maxNodes int, step func(*location, int, int)) []location {
	antinodes := []location{}
	ant := location{antenna.x, antenna.y}
	for {
		step(&ant, dx, dy)
		newAntinode := location{ant.x, ant.y}
		if af.outOfGrid(newAntinode) {
			break
		}
		antinodes = append(antinodes, newAntinode)
		if maxNodes != 0 && len(antinodes) >= maxNodes {
			break
		}
	}
	return antinodes
}

func (af *AntennaFinder) outOfGrid(loc location) bool {
	return loc.y < 0 || loc.y >= len(af.grid) || loc.x < 0 || loc.x >= len(af.grid[0])
}

func (af *AntennaFinder) clearGrid() {
	for _, row := range af.grid {
		clear(row)
	}
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

func isAntenna(chr byte) bool {
	return chr >= 'a' && chr <= 'z' || chr >= 'A' && chr <= 'Z' || chr >= '0' && chr <= '9'
}

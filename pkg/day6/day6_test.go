package day6_test

import (
	"fmt"
	"strings"

	"github.com/okulik/AoC2024/pkg/day6"
)

var testInput string = `
....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...
`

func ExampleGuardDetector_CountDistinctGuardLocations() {
	gd := day6.NewGuardDetector(strings.NewReader(testInput))
	fmt.Println(gd.CountDistinctGuardLocations())
	// Output: 41
}

func ExampleGuardDetector_CountNumberOfInfiniteLoops() {
	gd := day6.NewGuardDetector(strings.NewReader(testInput))
	fmt.Println(gd.CountNumberOfInfiniteLoops())
	// Output: 6
}

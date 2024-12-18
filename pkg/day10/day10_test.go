package day10_test

import (
	"fmt"
	"strings"

	"github.com/okulik/AoC2024/pkg/day10"
)

var testInput string = `
89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732
`

func ExampleHikingGuide_CalculateTrailheadScore() {
	hg := day10.NewHikingGuide(strings.NewReader(testInput))
	fmt.Println(hg.CalculateTrailheadScore())
	// Output: 36
}

func ExampleHikingGuide_CalculateTrailheadRate() {
	hg := day10.NewHikingGuide(strings.NewReader(testInput))
	fmt.Println(hg.CalculateTrailheadRate())
	// Output: 81
}

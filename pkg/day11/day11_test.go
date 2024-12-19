package day11_test

import (
	"fmt"
	"strings"

	"github.com/okulik/AoC2024/pkg/day11"
)

var testInput string = `
125 17
`

func ExamplePlutonianPebbles_TotalStoneCount() {
	pp := day11.NewPlutonianPebbles(strings.NewReader(testInput))
	fmt.Println(pp.TotalStoneCount(25))
	// Output: 55312
}

package day9_test

import (
	"fmt"
	"strings"

	"github.com/okulik/AoC2024/pkg/day9"
)

var testInput string = `
2333133121414131402
`

func ExampleDiskDefragmenter_CalculateChecksum() {
	dd := day9.NewDiskDefragmenter(strings.NewReader(testInput))
	fmt.Println(dd.CalculateChecksum())
	// Output: 1928
}

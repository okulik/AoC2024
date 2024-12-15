package day9_test

import (
	"fmt"
	"strings"

	"github.com/okulik/AoC2024/pkg/day9"
)

var testInput string = `
2333133121414131402
`

func ExampleDiskDefragmenter_DefragmentAndCalculateChecksum() {
	dd := day9.NewDiskDefragmenter(strings.NewReader(testInput))
	fmt.Println(dd.DefragmentAndCalculateChecksum())
	// Output: 1928
}

func ExampleDiskDefragmenter_BetterDefragmentAndCalculateChecksum() {
	dd := day9.NewDiskDefragmenter(strings.NewReader(testInput))
	fmt.Println(dd.BetterDefragmentAndCalculateChecksum())
	// Output: 2858
}

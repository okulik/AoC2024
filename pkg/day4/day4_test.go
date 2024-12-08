package day4_test

import (
	"fmt"
	"strings"

	"github.com/okulik/AoC2024/pkg/day4"
)

func ExampleWordCounter_CountSequences() {
	testInput := `
	XMASS.SAMX
	.M..A...M.
	..A.M..A..
	...SX.S...
	`

	wc := day4.NewWordCounter(strings.NewReader(testInput))
	count, _ := wc.CountSequences([]byte("XMAS"))
	fmt.Println(count)
	// Output: 5
}

func ExampleWordCounter_CountCrossSequences() {
	testInput := `
	.M.S......
	..A..MSMS.
	.M.S.MAA..
	..A.ASMSM.
	.M.S.M....
	..........
	S.S.S.S.S.
	.A.A.A.A..
	M.M.M.M.M.
	`

	wc := day4.NewWordCounter(strings.NewReader(testInput))
	count, _ := wc.CountCrossSequences([]byte("MAS"))
	fmt.Println(count)
	// Output: 9
}

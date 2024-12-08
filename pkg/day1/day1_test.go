package day1_test

import (
	"fmt"
	"strings"

	"github.com/okulik/AoC2024/pkg/day1"
)

var testInput = `
52630   25877
45307   76591
95628   97627
72163   78355
98522   95603
81222   61384
27778   94076
78355   89436
91829   43098
64666   25877
68679   78355
20680   89105
41869   90570
17863   78355
`

func ExampleCalculateDistances() {
	distances := day1.CalculateDistances(strings.NewReader(testInput))
	fmt.Println(distances)
	// Output:
	// 168958
}

func ExampleCalculateSimilarityScore() {
	similarity := day1.CalculateSimilarityScore(strings.NewReader(testInput))
	fmt.Println(similarity)
	// Output:
	// 235065
}

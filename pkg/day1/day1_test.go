package day1_test

import (
	"fmt"
	"strings"

	"github.com/okulik/AoC2024/pkg/day1"
)

func ExampleCalculateDistances() {
	testInput := `
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
	distances := day1.CalculateDistances(strings.NewReader(testInput))
	fmt.Printf("Total distances: %d\n", distances)

	similarity := day1.CalculateSimilarityScore(strings.NewReader(testInput))
	fmt.Printf("Similarity score: %d\n", similarity)

	// Output:
	// Total distances: 168958
	// Similarity score: 235065
}

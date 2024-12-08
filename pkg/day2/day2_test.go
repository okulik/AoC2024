package day2_test

import (
	"fmt"
	"strings"

	"github.com/okulik/AoC2024/pkg/day2"
)

func ExampleCountSafeReports() {
	testInput := `
25 27 28 28 30 32
61 63 66 68 68 66
51 54 54 57 60 60
50 52 52 53 56 60
73 75 76 76 83
19 20 24 26 28
36 38 41 42 45 49 47
56 59 63 64 64
26 29 32 36 40
70 72 74 75 77 80 84 89
81 83 88 89 92 95 96
79 80 85 87 89 92 93 90
1 3 6 8 9 12
`
	safe, unsafe := day2.CountSafeReports(strings.NewReader(testInput), day2.ReportIsSafe)
	fmt.Printf("%d, %d", safe, unsafe)
	// Output: 1, 12
}

func ExampleReportIsSafe() {
	levels := [][]int{
		{79, 80, 85, 87, 89, 92, 93, 90},
		{1, 3, 6, 8, 9, 12},
	}
	safe, i := day2.ReportIsSafe(levels[0])
	fmt.Printf("%t, %d\n", safe, i)

	safe, i = day2.ReportIsSafe(levels[1])
	fmt.Printf("%t, %d\n", safe, i)

	// Output:
	// false, 1
	// true, 0
}

func ExampleReportIsSafeWithDampener() {
	for _, level := range [][]int{
		{79, 80, 83, 83, 85, 86, 89, 90},
		{1, 3, 7, 8, 9, 12},
		{8, 6, 7, 8, 9, 12},
	} {
		safe, _ := day2.ReportIsSafeWithDampener(level)
		fmt.Println(safe)
	}

	// Output:
	// true
	// false
	// true
}

package day7_test

import (
	"fmt"
	"strings"

	"github.com/okulik/AoC2024/pkg/day7"
)

var testInput string = `
190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20
`

func ExampleFormulaCalibrator_SumFixFormulas() {
	fc := day7.NewFormulaCalibrator(strings.NewReader(testInput))
	fmt.Println(fc.SumFixFormulas())
	// Output: 3749
}

func ExampleFormulaCalibrator_SumFixFormulasWithConcatenation() {
	fc := day7.NewFormulaCalibrator(strings.NewReader(testInput))
	fmt.Println(fc.SumFixFormulasWithConcatenation())
	// Output: 11387
}

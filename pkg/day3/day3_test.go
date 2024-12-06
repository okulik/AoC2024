package day3_test

import (
	"fmt"
	"strings"

	"github.com/okulik/aoc2024/pkg/day3"
)

func ExampleSumAllMultiplications() {
	testInput := `
	from()when()*^$@:do()mul(1,2)'{^:-*what()mul(2,1) >- don't()mul(5,5)!do()mul(1,2)from(),where()mul(3,4)
	do()mul(2,3)don't()mul(3,4)do()mul(3333,-1)mul(5,6)
	`
	sum := day3.SumAllMultiplications(strings.NewReader(testInput))
	fmt.Printf("%d", sum)
	// Output: 54
}

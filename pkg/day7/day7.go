package day7

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type funcOp func(int, int) int

var opsMap = map[string]funcOp{
	"+": func(m, n int) int { return m + n },
	"*": func(m, n int) int { return m * n },
	"|": func(m, n int) int {
		num, _ := strconv.Atoi(strconv.Itoa(m) + strconv.Itoa(n))
		return num
	},
}

func Run() {
	file, err := os.Open("pkg/day7/input")
	if err != nil {
		panic("Can't open input file")
	}
	defer func() { _ = file.Close() }()

	fc := NewFormulaCalibrator(file)

	sum := fc.SumFixFormulas()
	fmt.Printf("Fixed formulas sum is: %d\n", sum)

	sum = fc.SumFixFormulasWithConcatenation()
	fmt.Printf("Fixed formulas with concatenations sum is: %d\n", sum)
}

type formula struct {
	sum  int
	nums []int
}

type FormulaCalibrator struct {
	formulas []formula
}

func NewFormulaCalibrator(input io.Reader) *FormulaCalibrator {
	formulas := make([]formula, 0)

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if len(line) == 0 {
			continue
		}

		parts := strings.Split(line, ": ")
		sum, err := strconv.Atoi(parts[0])
		if err != nil {
			panic("Invalid sum, line " + line)
		}

		strNums := strings.Split(parts[1], " ")
		nums := make([]int, 0, len(strNums))
		for _, num := range strNums {
			n, err := strconv.Atoi(num)
			if err != nil {
				panic("Invalid number, line " + line)
			}
			nums = append(nums, n)
		}

		formulas = append(formulas, formula{sum: sum, nums: nums})
	}

	return &FormulaCalibrator{
		formulas: formulas,
	}
}

func (fc *FormulaCalibrator) SumFixFormulas() int {
	totalSum := 0

	for _, formula := range fc.formulas {
		permOps := permutateAlphabet("+*", len(formula.nums)-1)

		for _, ops := range permOps {
			sum := formula.nums[0]
			for i, op := range ops {
				sum = opsMap[string(op)](sum, formula.nums[i+1])
			}
			if sum == formula.sum {
				totalSum += sum
				break
			}
		}
	}

	return totalSum
}

func (fc *FormulaCalibrator) SumFixFormulasWithConcatenation() int {
	totalSum := 0

	for _, formula := range fc.formulas {
		permOps := permutateAlphabet("+*|", len(formula.nums)-1)

		for _, ops := range permOps {
			sum := formula.nums[0]
			for i, op := range ops {
				sum = opsMap[string(op)](sum, formula.nums[i+1])
			}
			if sum == formula.sum {
				totalSum += sum
				break
			}
		}
	}

	return totalSum
}

func permutateAlphabet(alphabet string, length int) []string {
	if length == 0 {
		return []string{""}
	}

	var permutations []string
	rowPermutations := permutateAlphabet(alphabet, length-1)
	for _, letter := range alphabet {
		for _, word := range rowPermutations {
			permutations = append(permutations, string(letter)+word)
		}
	}

	return permutations
}

package day2

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func Run() {
	str, err := os.ReadFile("pkg/two/input")
	if err != nil {
		panic("Can't read input file")
	}

	safe, unsafe := CountSafeReports(strings.NewReader(string(str)), ReportIsSafe)
	fmt.Printf("Safe reports %d, unsafe reports %d\n", safe, unsafe)

	safe, unsafe = CountSafeReports(strings.NewReader(string(str)), ReportIsSafeWithDampener)
	fmt.Printf("Dampener safe reports %d, unsafe reports %d\n", safe, unsafe)
}

func CountSafeReports(reader *strings.Reader, safeFn func([]int) (bool, int)) (int, int) {
	safeReports, unsafeReports := 0, 0

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}
		strLevels := strings.Split(line, " ")
		if len(strLevels) < 2 {
			panic("Invalid line")
		}

		numLevels, err := ConvertStr2IntSlice(strLevels)
		if err != nil {
			continue
		}

		safe, _ := safeFn(numLevels)
		if safe {
			safeReports += 1
		} else {
			unsafeReports += 1
		}
	}

	return safeReports, unsafeReports
}

func ConvertStr2IntSlice(levels []string) ([]int, error) {
	numLevels := make([]int, 0, len(levels))

	for _, level := range levels {
		lvl, err := strconv.Atoi(level)
		if err != nil {
			return nil, fmt.Errorf("invalid input")
		}
		numLevels = append(numLevels, lvl)
	}

	return numLevels, nil
}

func ReportIsSafe(levels []int) (bool, int) {
	var dir *bool
	for i := 0; i < len(levels)-1; i++ {
		if !isAdjacent(levels[i], levels[i+1]) {
			return false, i
		}

		d := isAscending(levels[i], levels[i+1])
		if dir == nil {
			dir = &d
		}

		if *dir != d {
			return false, i
		}
	}

	return true, 0
}

func ReportIsSafeWithDampener(levels []int) (bool, int) {
	safe, i := ReportIsSafe(levels)
	if !safe {
		offsets := []int{-1, 0, 1}
		for _, offset := range offsets {
			index := i + offset
			if index < 0 || index >= len(levels) {
				continue
			}
			newLevels := newSliceDropAt(levels, index)
			safe, _ = ReportIsSafe(newLevels)
			if safe {
				break
			}
		}
	}

	return safe, i
}

func newSliceDropAt(levels []int, i int) []int {
	newLevels := make([]int, len(levels)-1)

	copy(newLevels, levels[:i])
	copy(newLevels[i:], levels[i+1:])

	return newLevels
}

func isAscending(a, b int) bool {
	return a < b
}

func isAdjacent(a, b int) bool {
	diff := int(math.Abs(float64(a - b)))
	return diff > 0 && diff < 4
}

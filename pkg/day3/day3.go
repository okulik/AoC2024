package day3

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var re *regexp.Regexp = regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)|do\(\)|don't\(\)`)

func Run() {
	file, err := os.Open("pkg/day3/input")
	if err != nil {
		panic("Can't open input file")
	}
	defer func() { _ = file.Close() }()

	sum := SumAllMultiplications(file)
	fmt.Printf("All instructions multiplication add up to %d\n", sum)
}

func SumAllMultiplications(input io.Reader) int {
	sum := 0

	do := true

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		matches := re.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			if strings.HasPrefix(match[0], "mul") && do {
				x, _ := strconv.Atoi(match[1])
				y, _ := strconv.Atoi(match[2])
				sum += x * y
				continue
			}

			if strings.HasPrefix(match[0], "do(") {
				do = true
				continue
			}

			if strings.HasPrefix(match[0], "don't(") {
				do = false
				continue
			}
		}
	}

	return sum
}

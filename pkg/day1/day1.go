package day1

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Run() {
	file, err := os.Open("pkg/day1/input")
	if err != nil {
		panic("Can't open input file")
	}
	defer func() { _ = file.Close() }()

	distances := CalculateDistances(file)
	fmt.Printf("Total distances: %d\n", distances)

	if _, err := file.Seek(0, 0); err != nil {
		panic("Can't seek file start")
	}

	similarity := CalculateSimilarityScore(file)
	fmt.Printf("Similarity score: %d\n", similarity)
}

func CalculateDistances(input io.Reader) int {
	lists, err := convertInputToLists(input)
	if err != nil {
		panic(fmt.Sprintln(err))
	}

	for _, list := range lists {
		sort.Ints(list)
	}

	sum := 0
	for i := 0; i < len(lists[0]); i++ {
		sum += int(math.Abs(float64(lists[0][i] - lists[1][i])))
	}

	return sum
}

func CalculateSimilarityScore(input io.Reader) int {
	lists, err := convertInputToLists(input)
	if err != nil {
		panic(fmt.Sprintln(err))
	}

	for _, list := range lists {
		sort.Ints(list)
	}

	m := make(map[int]int)
	for _, el := range lists[1] {
		m[el] += 1
	}

	sim := 0
	for _, el := range lists[0] {
		if cnt, ok := m[el]; ok {
			sim += el * cnt
		}
	}

	return sim
}

func convertInputToLists(input io.Reader) ([][]int, error) {
	lists := make([][]int, 2)
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			continue
		}
		elements := strings.Split(line, "   ")
		if err := appendLineElementsToLists(elements, lists); err != nil {
			return nil, fmt.Errorf("error occured: %v", err)
		}
	}

	return lists, nil
}

func appendLineElementsToLists(elements []string, lists [][]int) error {
	for ind, el := range elements {
		num, err := strconv.Atoi(el)
		if err != nil {
			return fmt.Errorf("not a number %s", el)
		}
		if lists[ind] == nil {
			lists[ind] = make([]int, 0)
		}
		lists[ind] = append(lists[ind], num)
	}

	return nil
}

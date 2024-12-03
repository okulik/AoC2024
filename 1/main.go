package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	lists, err := convertInputToLists()
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

	fmt.Printf("Total distances: %d\n", sum)

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

	fmt.Printf("Similarity score: %d\n", sim)
}

func convertInputToLists() ([][]int, error) {
	file, err := os.Open("input")
	if err != nil {
		return nil, fmt.Errorf("Can't open input file")
	}
	defer file.Close()

	lists := make([][]int, 2)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		elements := strings.Split(scanner.Text(), "   ")
		if err := appendLineElementsToLists(elements, lists); err != nil {
			return nil, fmt.Errorf("Error occured: %v", err)
		}
	}

	return lists, nil
}

func appendLineElementsToLists(elements []string, lists [][]int) error {
	for ind, el := range elements {
		num, err := strconv.Atoi(el)
		if err != nil {
			return fmt.Errorf("No a number %s", el)
		}
		if lists[ind] == nil {
			lists[ind] = make([]int, 0)
		}
		lists[ind] = append(lists[ind], num)
	}

	return nil
}

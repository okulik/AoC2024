package day11

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func Run() {
	file, err := os.Open("pkg/day11/input")
	if err != nil {
		panic("Can't open input file")
	}
	defer func() { _ = file.Close() }()

	pp := NewPlutonianPebbles(file)
	fmt.Printf("Total number of stones after 25 blinks is %d\n", pp.TotalStoneCount(25))
}

type PlutonianPebbles struct {
	stones []string
}

func NewPlutonianPebbles(input io.Reader) *PlutonianPebbles {
	stones := make([]string, 0)

	var row string
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		row = strings.TrimSpace(scanner.Text())
		if len(row) == 0 {
			continue
		}
		stones = strings.Split(row, " ")
		break
	}

	return &PlutonianPebbles{
		stones: stones,
	}
}

func (pp *PlutonianPebbles) TotalStoneCount(blinks int) int {
	stones := pp.stones
	for i := 0; i < 25; i++ {
		stones = pp.rearrangeStones(stones)
	}

	return len(stones)
}

func (pp *PlutonianPebbles) rearrangeStones(stones []string) []string {
	newStones := make([]string, 0, len(stones))

	for _, stone := range stones {
		if stone == "0" {
			newStones = append(newStones, "1")
			continue
		}
		if len(stone)%2 == 0 {
			left, err := strconv.Atoi(stone[:len(stone)/2])
			if err != nil {
				panic("Can't read stone")
			}
			right, err := strconv.Atoi(stone[len(stone)/2:])
			if err != nil {
				panic("Can't read stone")
			}
			newStones = append(newStones, fmt.Sprintf("%d", left))
			newStones = append(newStones, fmt.Sprintf("%d", right))
			continue
		}
		stoneNum, err := strconv.Atoi(stone)
		if err != nil {
			panic("Can't read stone")
		}
		newStones = append(newStones, fmt.Sprintf("%d", stoneNum*2024))
	}

	return newStones
}

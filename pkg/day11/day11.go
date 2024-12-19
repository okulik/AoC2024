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
	fmt.Printf("Total number of stones after 75 blinks is %d\n", pp.TotalStoneCount(75))
}

type PlutonianPebbles struct {
	stones map[int]int
}

func NewPlutonianPebbles(input io.Reader) *PlutonianPebbles {
	reader := bufio.NewReader(input)
	line, _ := reader.ReadString('\n')
	stones := make(map[int]int, 0)
	for _, numStr := range strings.Split(strings.TrimSpace(line), " ") {
		if len(numStr) == 0 {
			continue
		}
		num, _ := strconv.Atoi(numStr)
		stones[num] = 1
	}

	return &PlutonianPebbles{
		stones: stones,
	}
}

func (pp *PlutonianPebbles) TotalStoneCount(blinks int) int {
	stones := pp.stones

	for i := 0; i < blinks; i++ {
		newStones := make(map[int]int)
		for num, cnt := range stones {
			numStr := strconv.Itoa(num)
			numStrLen := len(numStr)
			even := numStrLen%2 == 0
			if num == 0 {
				newStones[1] += cnt
				continue
			} else if even {
				for _, stn := range []string{numStr[:numStrLen/2], numStr[numStrLen/2:]} {
					num, _ := strconv.Atoi(stn)
					newStones[num] += cnt
				}
			} else {
				newStones[num*2024] += cnt
			}
		}
		stones = newStones
	}

	sum := 0
	for _, values := range stones {
		sum += values
	}

	return sum
}

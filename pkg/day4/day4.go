package day4

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type move func(r, c int) (int, int)

var moves = map[string]move{
	"e":  func(r, c int) (int, int) { return r, c + 1 },
	"ne": func(r, c int) (int, int) { return r - 1, c + 1 },
	"n":  func(r, c int) (int, int) { return r - 1, c },
	"nw": func(r, c int) (int, int) { return r - 1, c - 1 },
	"w":  func(r, c int) (int, int) { return r, c - 1 },
	"sw": func(r, c int) (int, int) { return r + 1, c - 1 },
	"s":  func(r, c int) (int, int) { return r + 1, c },
	"se": func(r, c int) (int, int) { return r + 1, c + 1 },
}

func Run() {
	file, err := os.Open("pkg/day4/input")
	if err != nil {
		panic("Can't open input file")
	}
	defer func() { _ = file.Close() }()

	wordCounter := NewWordCounter(file)
	count := wordCounter.CountSequences([]byte("XMAS"))
	fmt.Printf("Total count of XMAS is %d\n", count)

	count = wordCounter.CountCrossSequences([]byte("MAS"))
	fmt.Printf("Total count of X-MAS is %d\n", count)
}

type WordCounter struct {
	letterMatrix [][]byte
}

func NewWordCounter(input io.Reader) *WordCounter {
	matrix := make([][]byte, 0)
	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			continue
		}
		row := make([]byte, 0, len(line))
		for _, val := range line {
			row = append(row, byte(val))
		}
		matrix = append(matrix, row)
	}

	return &WordCounter{
		letterMatrix: matrix,
	}
}

func (wc *WordCounter) CountSequences(seq []byte) int {
	if len(seq) < 2 {
		panic("Sequence should be at least 2 chars long")
	}

	count := 0
	for row := range wc.letterMatrix {
		for col := range wc.letterMatrix[row] {
			count += wc.countSequence(seq, row, col)
		}
	}

	return count
}

func (wc *WordCounter) countSequence(seq []byte, row, col int) int {
	count := 0
	for _, foo := range moves {
		if wc.findSubsequence(seq, row, col, foo) {
			count++
		}
	}

	return count
}

func (wc *WordCounter) CountCrossSequences(seq []byte) int {
	sum := 0

	for row := range wc.letterMatrix {
		for col := range wc.letterMatrix[row] {
			sum += wc.countCrossSequence(seq, row, col)
		}
	}

	return sum / 2
}

func (wc *WordCounter) countCrossSequence(seq []byte, row, col int) int {
	count := 0

	found :=
		wc.findSubsequence(seq, row, col, moves["ne"]) &&
			(wc.findSubsequence(seq, row, col+len(seq)-1, moves["nw"]) ||
				wc.findSubsequence(seq, row-len(seq)+1, col, moves["se"]))
	if found {
		count++
	}

	found =
		wc.findSubsequence(seq, row, col, moves["nw"]) &&
			(wc.findSubsequence(seq, row, col-len(seq)+1, moves["ne"]) ||
				wc.findSubsequence(seq, row-len(seq)+1, col, moves["sw"]))
	if found {
		count++
	}

	found =
		wc.findSubsequence(seq, row, col, moves["se"]) &&
			(wc.findSubsequence(seq, row+len(seq)-1, col, moves["ne"]) ||
				wc.findSubsequence(seq, row, col+len(seq)-1, moves["sw"]))
	if found {
		count++
	}

	found =
		wc.findSubsequence(seq, row, col, moves["sw"]) &&
			(wc.findSubsequence(seq, row, col-len(seq)+1, moves["se"]) ||
				wc.findSubsequence(seq, row+len(seq)-1, col, moves["nw"]))
	if found {
		count++
	}

	return count
}

func (wc *WordCounter) findSubsequence(seq []byte, row, col int, foo func(int, int) (int, int)) bool {
	if row < 0 || col < 0 || row >= len(wc.letterMatrix) || col >= len(wc.letterMatrix[0]) {
		return false
	}

	if wc.letterMatrix[row][col] != seq[0] {
		return false
	}

	if len(seq) == 1 {
		return true
	}

	row, col = foo(row, col)

	return wc.findSubsequence(seq[1:], row, col, foo)
}

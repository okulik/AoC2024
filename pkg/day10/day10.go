package day10

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func Run() {
	file, err := os.Open("pkg/day10/input")
	if err != nil {
		panic("Can't open input file")
	}
	defer func() { _ = file.Close() }()

	hg := NewHikingGuide(file)
	fmt.Printf("Total trailhead score is %d\n", hg.CalculateTrailheadScore())
}

const (
	trailHead = 0
	trailTop  = 9
)

type HikingGuide struct {
	hikingMap [][]byte
}

func NewHikingGuide(input io.Reader) *HikingGuide {
	hikingMap := make([][]byte, 0)

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		row := strings.TrimSpace(scanner.Text())
		if len(row) == 0 {
			continue
		}

		hikingMapRow := make([]byte, len(row))
		for i := range row {
			val := row[i] - '0'
			if val >= trailHead && val <= trailTop {
				hikingMapRow[i] = val
			} else {
				hikingMapRow[i] = 255
			}
		}
		hikingMap = append(hikingMap, hikingMapRow)
	}

	return &HikingGuide{
		hikingMap: hikingMap,
	}
}

func (hg *HikingGuide) CalculateTrailheadScore() int {
	sum := 0
	for i := range hg.hikingMap {
		for j := range hg.hikingMap[i] {
			cache := make(map[string]struct{})
			if hg.hikingMap[i][j] == trailHead {
				sum += hg.pathCount(i, j, trailHead, &cache)
			}
			// if len(cache) > 0 {
			// 	keys := make([]string, 0, len(cache))
			// 	for k := range cache {
			// 		keys = append(keys, k)
			// 	}
			// 	fmt.Println(keys)
			// }
		}
	}

	return sum
}

func (hg *HikingGuide) pathCount(i, j int, expect byte, cache *map[string]struct{}) int {
	if i < 0 || i >= len(hg.hikingMap) || j < 0 || j >= len(hg.hikingMap[i]) {
		return 0
	}

	if hg.hikingMap[i][j] != expect {
		return 0
	}

	if hg.hikingMap[i][j] == trailTop {
		cacheKey := fmt.Sprintf("%d%d", i, j)
		if _, ok := (*cache)[cacheKey]; !ok {
			(*cache)[cacheKey] = struct{}{}
			return 1
		}
		return 0
	}

	return hg.pathCount(i-1, j, expect+1, cache) +
		hg.pathCount(i+1, j, expect+1, cache) +
		hg.pathCount(i, j-1, expect+1, cache) +
		hg.pathCount(i, j+1, expect+1, cache)
}

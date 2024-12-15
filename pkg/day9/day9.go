package day9

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func Run() {
	file, err := os.Open("pkg/day9/input")
	if err != nil {
		panic("Can't open input file")
	}
	defer func() { _ = file.Close() }()

	dd := NewDiskDefragmenter(file)
	fmt.Printf("Check sum is %d\n", dd.CalculateChecksum())
}

type DiskDefragmenter struct {
	indexBitmap []int
	diskImage   []*int
}

func NewDiskDefragmenter(input io.Reader) *DiskDefragmenter {
	indexBitmap := make([]int, 100000)
	for i := 0; i < 100000; i++ {
		indexBitmap[i] = i
	}

	diskImage := make([]*int, 0)
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if len(line) == 0 {
			continue
		}

		fileIndex := 0
		for i := range line {
			for j := 0; j < int(line[i]-'0'); j++ {
				var indexPtr *int
				if i%2 == 0 {
					indexPtr = &indexBitmap[fileIndex]
				}
				diskImage = append(diskImage, indexPtr)
			}
			if i%2 == 0 {
				fileIndex++
			}
		}
	}

	innerIndex := 0
	for i := len(diskImage) - 1; i >= 0; i-- {
		if diskImage[i] == nil {
			continue
		}
		for j := innerIndex; j < i; j++ {
			if diskImage[j] == nil {
				diskImage[i], diskImage[j] = nil, diskImage[i]
				innerIndex = j + 1
				break
			}
		}
	}

	return &DiskDefragmenter{
		indexBitmap: indexBitmap,
		diskImage:   diskImage,
	}
}

func (dd *DiskDefragmenter) CalculateChecksum() int {
	sum := 0
	for i := range dd.diskImage {
		if dd.diskImage[i] == nil {
			continue
		}
		sum += i * *dd.diskImage[i]
	}

	return sum
}

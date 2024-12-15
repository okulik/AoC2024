package day9

import (
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
	fmt.Printf("Defragmented check sum is %d\n", dd.DefragmentAndCalculateChecksum())
	fmt.Printf("Better defragmented check sum is %d\n", dd.BetterDefragmentAndCalculateChecksum())
}

type DiskDefragmenter struct {
	indexBitmap []int
	diskImage   []*int
}

func NewDiskDefragmenter(input io.Reader) *DiskDefragmenter {
	content, err := io.ReadAll(input)
	if err != nil {
		panic("Can't read input file")
	}

	contentStr := strings.TrimSpace(string(content))
	diskImage := make([]*int, 0)
	indexBitmap := make([]int, len(contentStr))

	for i := 0; i < len(contentStr); i++ {
		indexBitmap[i] = i
	}

	fileIndex := 0
	for i := range contentStr {
		for j := 0; j < int(contentStr[i]-'0'); j++ {
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

	return &DiskDefragmenter{
		indexBitmap: indexBitmap,
		diskImage:   diskImage,
	}
}

func (dd *DiskDefragmenter) DefragmentAndCalculateChecksum() int {
	diskImage := make([]*int, len(dd.diskImage))
	copy(diskImage, dd.diskImage)

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

	return calculateCheckSum(diskImage)
}

func (dd *DiskDefragmenter) BetterDefragmentAndCalculateChecksum() int {
	diskImage := make([]*int, len(dd.diskImage))
	copy(diskImage, dd.diskImage)

	for i := len(diskImage) - 1; i >= 0; {
		if diskImage[i] == nil {
			i--
			continue
		}
		fileStart, fileStop := dd.findFile(diskImage, i)
		size := fileStop - fileStart + 1
		if freeStart := dd.findFreeSpace(diskImage, fileStart, size); freeStart != -1 {
			for j := 0; j < size; j++ {
				diskImage[freeStart+j], diskImage[fileStart+j] = diskImage[fileStart+j], nil
			}
		}
		i -= size
	}

	return calculateCheckSum(diskImage)
}

func (dd *DiskDefragmenter) findFile(diskImage []*int, stop int) (int, int) {
	fileIndex := *diskImage[stop]
	start := stop
	for i := start - 1; i >= 0; i-- {
		if diskImage[i] == nil || *diskImage[i] != fileIndex {
			break
		}
		start = i
	}

	return start, stop
}

func (dd *DiskDefragmenter) findFreeSpace(diskImage []*int, stop, size int) int {
	var freeStart int = -1
	var freeStop int

	for i := 0; i < stop; i++ {
		if diskImage[i] != nil {
			freeStart = -1
			continue
		}
		if freeStart == -1 {
			freeStart = i
		}
		freeStop = i
		if size <= freeStop-freeStart+1 {
			return freeStart
		}
	}

	return -1
}

func calculateCheckSum(diskImage []*int) int {
	sum := 0
	for i := range diskImage {
		if diskImage[i] == nil {
			continue
		}
		sum += i * *diskImage[i]
	}

	return sum
}

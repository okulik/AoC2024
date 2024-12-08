package day5

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func Run() {
	file, err := os.Open("pkg/day5/input")
	if err != nil {
		panic("Can't open input file")
	}
	defer func() { _ = file.Close() }()

	updatesValidator := NewUpdatesValidator(file)

	sum := updatesValidator.SumUpdatesWithCorrectOrder()
	fmt.Printf("Valid updates middle page sum is: %d\n", sum)

	sum = updatesValidator.SumUpdatesWithIncorrectOrder()
	fmt.Printf("Invalid/fixed updates middle page sum is: %d\n", sum)
}

type UpdatesValidator struct {
	rules   map[string]bool
	updates [][]int
}

func NewUpdatesValidator(input io.Reader) *UpdatesValidator {
	rules := make(map[string]bool)
	updates := make([][]int, 0)

	scanner := bufio.NewScanner(input)
	orderingRules := true
	anyRules := false

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if len(line) == 0 && anyRules {
			orderingRules = false
			continue
		}

		if orderingRules {
			if len(line) == 0 {
				continue
			}

			anyRules = true

			rules[line] = true
		} else {
			if len(line) == 0 {
				continue
			}

			pagesStrArr := strings.Split(line, ",")
			pagesNumArr := make([]int, 0, len(pagesStrArr))

			for _, el := range pagesStrArr {
				pageNum, err := strconv.Atoi(el)
				if err != nil {
					panic("invalid page in line: " + line)
				}
				pagesNumArr = append(pagesNumArr, pageNum)
			}
			updates = append(updates, pagesNumArr)
		}
	}

	return &UpdatesValidator{
		rules:   rules,
		updates: updates,
	}
}

func (uv *UpdatesValidator) SumUpdatesWithCorrectOrder() int {
	sum := 0
	for _, update := range uv.updates {
		if len(update) < 2 {
			continue
		}
		if uv.isCorrectOrder(update) {
			sum += uv.middlePage(update)
		}
	}

	return sum
}

func (uv *UpdatesValidator) SumUpdatesWithIncorrectOrder() int {
	sum := 0
	for _, update := range uv.updates {
		if !uv.isCorrectOrder(update) {
			for {
				uv.correctUpdate(update)
				if uv.isCorrectOrder(update) {
					break
				}
			}
			sum += uv.middlePage(update)
		}
	}

	return sum
}

func (uv *UpdatesValidator) isCorrectOrder(update []int) bool {
	for i := 0; i < len(update)-1; i++ {
		for j := i + 1; j < len(update); j++ {
			badOrder := fmt.Sprintf("%d|%d", update[j], update[i])
			if _, ok := uv.rules[badOrder]; ok {
				return false
			}
		}
	}

	return true
}

func (uv *UpdatesValidator) middlePage(update []int) int {
	return update[len(update)/2]
}

func (uv *UpdatesValidator) correctUpdate(update []int) {
	for i := 0; i < len(update)-1; i++ {
		for j := i + 1; j < len(update); j++ {
			badOrder := fmt.Sprintf("%d|%d", update[j], update[i])
			if _, ok := uv.rules[badOrder]; ok {
				update[i], update[j] = update[j], update[i]
				return
			}
		}
	}
}

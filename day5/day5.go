package day5

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

var incorrectArr [][]string

func ComputeAOCDay5_1(name string) {
	nextPageNumbers, input := readFile(name)

	var res [][]string
	for _, inp := range input {
		data := strings.Split(inp, ",")
		var r []string
		for i := 1; i < len(data); i++ {
			currPage := string(data[i])
			prevPage := string(data[i-1])
			v, ok := nextPageNumbers[prevPage]
			if ok {
				for _, n := range v {
					if n == currPage {
						r = append(r, prevPage)
						break
					}
				}
			} else {
				r = []string{}
				break
			}
		}
		if len(r) == len(data)-1 {
			lastPage := string(data[len(data)-1])
			r = append(r, lastPage)
			res = append(res, r)
		} else {
			incorrectArr = append(incorrectArr, data)
		}
	}

	s := computeSumOfMid(res)

	fmt.Printf("Result day 5 part 1: %d\n", s)
}

func ComputeAOCDay5_2(name string) {

	nextPageNumbers, _ := readFile(name)

	var correctArr [][]string
	for _, arr := range incorrectArr {
		isValidArray := false
		for !isValidArray {
			for i := 1; i < len(arr); i++ {
				currIdx := i
				prevIdx := i - 1

				currPage := arr[currIdx]
				prevPage := arr[prevIdx]

				isCurrPagePresent := false
				isPrevPagePresent := false

				vPrevPage, okPrev := nextPageNumbers[prevPage]
				if okPrev {
					for _, p := range vPrevPage {
						if p == currPage {
							isCurrPagePresent = true
							break
						}
					}
				}

				if !isCurrPagePresent || !okPrev {
					vCurrPage, okCurr := nextPageNumbers[currPage]
					if okCurr {
						for _, p := range vCurrPage {
							if p == prevPage {
								isPrevPagePresent = true
								break
							}
						}
					}
				}

				if isPrevPagePresent {
					// swap
					arr[currIdx] = prevPage
					arr[prevIdx] = currPage
				}
			}
			if checkArray(arr, nextPageNumbers) {
				correctArr = append(correctArr, arr)
				isValidArray = true
			}
		}
	}

	s := computeSumOfMid(correctArr)

	fmt.Printf("Result day 5 part 2: %d\n", s)
}

func readFile(name string) (map[string][]string, []string) {
	file, err := os.Open(name)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	nextPageNumbers := make(map[string][]string)
	var input []string
	r := bufio.NewReader(file)
	for {
		line, _, err := r.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		s := string(line)

		if strings.Contains(s, "|") {
			n := strings.Split(s, "|")
			n1 := n[0]
			n2 := n[1]
			v, ok := nextPageNumbers[n1]
			if ok {
				v = append(v, n2)
				nextPageNumbers[n1] = v
			} else {
				nextPageNumbers[n1] = []string{n2}
			}
		} else if len(s) > 0 {
			input = append(input, s)
		}
	}
	return nextPageNumbers, input
}

func computeSumOfMid(res [][]string) int {
	s := 0
	for _, r := range res {
		i := float64(len(r) / 2)
		mid := int(math.Ceil(i))
		num, _ := strconv.Atoi(r[mid])
		s += num
	}
	return s
}

func checkArray(arr []string, nextPageNum map[string][]string) bool {
	var isValidPage bool
	for i := 1; i < len(arr); i++ {
		isValidPage = false
		prevPage := arr[i-1]
		currPage := arr[i]
		v, ok := nextPageNum[prevPage]
		if ok {
			for _, n := range v {
				if n == currPage {
					isValidPage = true
				}
			}
			if !isValidPage {
				return false
			}
		} else {
			return false
		}
	}
	return isValidPage
}

package day3

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
)

var input string

func ComputeAOCDay3_1(name string) {
	input = readFile(name)

	exp := regexp.MustCompile(`mul\([0-9]+\,[0-9]+\)`)
	matches := exp.FindAllString(input, -1)

	exp = regexp.MustCompile(`[0-9]+`)
	res := 0
	for _, match := range matches {
		matches = exp.FindAllString(match, -1)
		a, _ := strconv.Atoi(matches[0])
		b, _ := strconv.Atoi(matches[1])
		res += a * b
	}

	fmt.Printf("Result day 3 part 1: %d\n", res)
}

func ComputeAOCDay3_2(name string) {
	res := 0
	idx := 0

	exp := regexp.MustCompile(`(?i)(mul\(\d+,\d+\)|don't|do)`)
	matches := exp.FindAllString(input, -1)

	exp = regexp.MustCompile(`[0-9]+`)
	for idx < len(matches) {
		if matches[idx] != "don't" && matches[idx] != "do" {
			m := exp.FindAllString(matches[idx], -1)
			a, _ := strconv.Atoi(m[0])
			b, _ := strconv.Atoi(m[1])
			res += a * b
			idx += 1
		} else if matches[idx] == "don't" {
			nextIdx := 0
			j := idx + 1
			for j < len(matches) {
				if matches[j] == "do" || j == len(matches)-1 {
					nextIdx = j + 1
					break
				}
				j += 1
			}
			idx = nextIdx
		} else {
			idx += 1
		}
	}
	fmt.Printf("Result day 3 part 2: %d\n", res)
}

func readFile(name string) string {
	file, err := os.Open(name)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	var s string
	r := bufio.NewReader(file)
	for {
		line, _, err := r.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		s += string(line)
	}
	return s
}

package day4

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

var input []string

func ComputeAOCDay4(name string) {
	input = readFile(name)

	runes := make(map[[2]int]string)
	neighbourPositions := make(map[[2]int][8][2]int)
	for row, r := range input {
		for col, c := range strings.Split(r, "") {
			runes[[2]int{row, col}] = c
			neighbours := [8][2]int{}
			for i := 0; i < 8; i += 1 {
				if i == 0 {
					x := row - 1
					y := col - 1
					if x < 0 || y < 0 {
						neighbours[i] = [2]int{-1, -1}
					} else {
						neighbours[i] = [2]int{x, y}
					}
				} else if i == 1 {
					x := row - 1
					y := col
					if x < 0 {
						neighbours[i] = [2]int{-1, -1}
					} else {
						neighbours[i] = [2]int{x, y}
					}
				} else if i == 2 {
					x := row - 1
					y := col + 1
					if x < 0 || y >= len(r) {
						neighbours[i] = [2]int{-1, -1}
					} else {
						neighbours[i] = [2]int{x, y}
					}
				} else if i == 3 {
					x := row
					y := col + 1
					if y >= len(r) {
						neighbours[i] = [2]int{-1, -1}
					} else {
						neighbours[i] = [2]int{x, y}
					}
				} else if i == 4 {
					x := row + 1
					y := col + 1
					if x >= len(input) || y >= len(r) {
						neighbours[i] = [2]int{-1, -1}
					} else {
						neighbours[i] = [2]int{x, y}
					}
				} else if i == 5 {
					x := row + 1
					y := col
					if x >= len(input) {
						neighbours[i] = [2]int{-1, -1}
					} else {
						neighbours[i] = [2]int{x, y}
					}
				} else if i == 6 {
					x := row + 1
					y := col - 1
					if x >= len(input) || y < 0 {
						neighbours[i] = [2]int{-1, -1}
					} else {
						neighbours[i] = [2]int{x, y}
					}
				} else if i == 7 {
					x := row
					y := col - 1
					if y < 0 {
						neighbours[i] = [2]int{-1, -1}
					} else {
						neighbours[i] = [2]int{x, y}
					}
				}
			}
			neighbourPositions[[2]int{row, col}] = neighbours
		}
	}

	results := 0
	for row, r := range input {
		for col, c := range strings.Split(r, "") {
			if c != "X" {
				continue
			}

			nps := neighbourPositions[[2]int{row, col}]
			for i, p := range nps {
				v := runes[p]
				if v != "M" {
					continue
				}

				npsP := neighbourPositions[[2]int{p[0], p[1]}]
				v = runes[npsP[i]]
				if v != "A" {
					continue
				}

				npsP = neighbourPositions[[2]int{npsP[i][0], npsP[i][1]}]
				v = runes[npsP[i]]
				if v != "S" {
					continue
				}

				results += 1
			}
		}
	}

	fmt.Printf("Result day 4 part 1: %d\n", results)

	results = 0
	for row, r := range input {
		for col, c := range strings.Split(r, "") {
			if c != "A" {
				continue
			}

			nps := neighbourPositions[[2]int{row, col}]
			n0, n4 := nps[0], nps[4]
			n2, n6 := nps[2], nps[6]

			r0, r4 := runes[n0], runes[n4]
			r2, r6 := runes[n2], runes[n6]
			if (r0 == "M" && r4 == "S" || r0 == "S" && r4 == "M") && (r2 == "M" && r6 == "S" || r2 == "S" && r6 == "M") {
				results += 1
			}
		}
	}

	fmt.Printf("Result day 4 part 2: %d\n", results)
}

func readFile(name string) []string {
	file, err := os.Open(name)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	s := []string{}
	r := bufio.NewReader(file)
	for {
		line, _, err := r.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		s = append(s, string(line))
	}
	return s
}

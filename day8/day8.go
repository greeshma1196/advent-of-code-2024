package day8

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode"
)

func ComputeAOCDay8_1(name string) {
	input := readFile(name)

	rows := len(input)
	cols := len(input[0])
	antennaPos := make(map[string][][2]int)

	for i, inp := range input {
		line := strings.Split(inp, "")
		for j, l := range line {
			if unicode.IsDigit(rune(l[0])) || unicode.IsLetter(rune(l[0])) {
				antennaPos[l] = append(antennaPos[l], [2]int{i, j})
			}
		}
	}

	res := 0

	validAntinodePos := make(map[[2]int]bool)
	for _, pos := range antennaPos {
		for i := 0; i < len(pos)-1; i++ {
			point1 := pos[i]
			for j := i + 1; j < len(pos); j++ {
				point2 := pos[j]
				antinodePos1, antinodePos2 := computeAntinodePos_1(point1, point2)
				if isValidAntinodePos(antinodePos1, rows, cols) {
					_, ok := validAntinodePos[antinodePos1]
					if !ok {
						validAntinodePos[antinodePos1] = true
						res += 1
					}
				}
				if isValidAntinodePos(antinodePos2, rows, cols) {
					_, ok := validAntinodePos[antinodePos2]
					if !ok {
						validAntinodePos[antinodePos2] = true
						res += 1
					}
				}
			}
		}
	}
	fmt.Printf("Result of day 8 part 1: %d\n", res)
}

func ComputeAOCDay8_2(name string) {
	input := readFile(name)

	rows := len(input)
	cols := len(input[0])
	antennaPos := make(map[string][][2]int)

	for i, inp := range input {
		line := strings.Split(inp, "")
		for j, l := range line {
			if unicode.IsDigit(rune(l[0])) || unicode.IsLetter(rune(l[0])) {
				antennaPos[l] = append(antennaPos[l], [2]int{i, j})
			}
		}
	}

	res := 0

	validAntinodePos := make(map[[2]int]bool)
	for _, pos := range antennaPos {
		for i := 0; i < len(pos)-1; i++ {
			point1 := pos[i]
			for j := i + 1; j < len(pos); j++ {
				point2 := pos[j]
				antinodePos := computeAntinodePos_2(point1, point2, rows, cols)
				for _, p := range antinodePos {
					_, ok := validAntinodePos[p]
					if !ok {
						validAntinodePos[p] = true
						res += 1
						fmt.Printf("%+v\n", p)
					}
				}
			}
		}
	}
	fmt.Printf("Result of day 8 part 2: %d\n", res)
}

func readFile(name string) []string {
	var t []string
	file, err := os.Open(name)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	r := bufio.NewReader(file)
	for {
		line, _, err := r.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		t = append(t, string(line))
	}

	return t
}

func computeAntinodePos_1(point1 [2]int, point2 [2]int) ([2]int, [2]int) {
	var antinodeP1 [2]int
	var antinodeP2 [2]int

	x1 := point1[0]
	x2 := point2[0]
	diffX := x2 - x1

	y1 := point1[1]
	y2 := point2[1]
	diffY := y2 - y1

	antinodeP1 = [2]int{x1 - diffX, y1 - diffY}
	antinodeP2 = [2]int{x2 + diffX, y2 + diffY}

	return antinodeP1, antinodeP2
}

func computeAntinodePos_2(point1 [2]int, point2 [2]int, rows int, cols int) [][2]int {
	var antinodePos [][2]int

	x1 := point1[0]
	x2 := point2[0]
	diffX := x2 - x1

	y1 := point1[1]
	y2 := point2[1]
	diffY := y2 - y1

	antinodePos = append(antinodePos, [2]int{x1, y1})
	antinodePos = append(antinodePos, [2]int{x2, y2})

	x1 = x1 - diffX
	y1 = y1 - diffY
	for x1 >= 0 && x1 < rows && y1 >= 0 && y1 < cols {
		antinodePos = append(antinodePos, [2]int{x1, y1})
		x1 = x1 - diffX
		y1 = y1 - diffY
	}

	x2 = x2 + diffX
	y2 = y2 + diffY
	for x2 >= 0 && x2 < rows && y2 >= 0 && y2 < cols {
		antinodePos = append(antinodePos, [2]int{x2, y2})
		x2 = x2 + diffX
		y2 = y2 + diffY
	}

	return antinodePos
}

func isValidAntinodePos(point [2]int, rows int, cols int) bool {
	if (point[0] >= 0 && point[0] < rows) && (point[1] >= 0 && point[1] < cols) {
		return true
	}
	return false
}

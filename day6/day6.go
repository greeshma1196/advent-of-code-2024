package day6

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func ComputeAOCDay6_1(name string) {
	input := readFile(name)

	res := stepCounter(input)

	fmt.Printf("Result day 6 part 1: %d\n", res)
}

func ComputeAOCDay6_2(name string) {
	input := readFile(name)

	direction, r, c := getStartPos(input)

	rows := len(input)
	cols := len(input[0])

	obsCount := 0

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if i == r && j == c {
				continue
			}
			if input[i][j] == "." {
				input[i][j] = "#"
				if isLoopPresent(input, direction, r, c) {
					obsCount += 1
				}
				input[i][j] = "."
			}
		}
	}

	fmt.Printf("Result day 6 part 2: %d\n", obsCount)
}

func stepCounter(input [][]string) int {
	direction_, row, col := getStartPos(input)

	rows := len(input)
	cols := len(input[0])

	direction := direction_
	inputp1 := input
	isExit := false
	r := row
	c := col
	for !isExit {
		if direction == "UP" {
			for r >= 0 {
				if r == 0 && inputp1[r][c] != "#" {
					inputp1[r][c] = "X"
					isExit = true
					break
				}
				if inputp1[r][c] == "#" {
					direction = "RIGHT"
					r += 1
					c += 1
					break
				} else {
					inputp1[r][c] = "X"
					r -= 1
				}
			}
		} else if direction == "RIGHT" {
			for c < cols {
				if c == cols-1 && inputp1[r][c] != "#" {
					inputp1[r][c] = "X"
					isExit = true
					break
				}
				if inputp1[r][c] == "#" {
					direction = "DOWN"
					r += 1
					c -= 1
					break
				} else {
					inputp1[r][c] = "X"
					c += 1
				}
			}
		} else if direction == "DOWN" {
			for r < rows {
				if r == rows-1 && inputp1[r][c] != "#" {
					inputp1[r][c] = "X"
					isExit = true
					break
				}
				if inputp1[r][c] == "#" {
					direction = "LEFT"
					r -= 1
					c -= 1
					break
				} else {
					inputp1[r][c] = "X"
					r += 1
				}
			}
		} else if direction == "LEFT" {
			for c >= 0 {
				if c == 0 && inputp1[r][c] != "#" {
					inputp1[r][c] = "X"
					isExit = true
					break
				}
				if inputp1[r][c] == "#" {
					direction = "UP"
					r -= 1
					c += 1
					break
				} else {
					inputp1[r][c] = "X"
					c -= 1
				}
			}
		}
		if isExit {
			break
		}
	}

	res := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if inputp1[i][j] == "X" {
				res += 1
			}
		}
	}

	return res
}

func isLoopPresent(input [][]string, direction string, r int, c int) bool {

	rows := len(input)
	cols := len(input[0])

	inputp1 := input
	isExit := false
	counter := 0
	for !isExit {
		if counter > 100000 {
			break
		}
		if direction == "UP" {
			for r >= 0 {
				if r == 0 && inputp1[r][c] != "#" {
					isExit = true
					break
				}
				if inputp1[r][c] == "#" {
					direction = "RIGHT"
					r += 1
					c += 1
					break
				} else {
					r -= 1
				}
			}
		} else if direction == "RIGHT" {
			for c < cols {
				if c == cols-1 && inputp1[r][c] != "#" {
					isExit = true
					break
				}
				if inputp1[r][c] == "#" {
					direction = "DOWN"
					r += 1
					c -= 1
					break
				} else {
					c += 1
				}
			}
		} else if direction == "DOWN" {
			for r < rows {
				if r == rows-1 && inputp1[r][c] != "#" {
					isExit = true
					break
				}
				if inputp1[r][c] == "#" {
					direction = "LEFT"
					r -= 1
					c -= 1
					break
				} else {
					r += 1
				}
			}
		} else if direction == "LEFT" {
			for c >= 0 {
				if c == 0 && inputp1[r][c] != "#" {
					isExit = true
					break
				}
				if inputp1[r][c] == "#" {
					direction = "UP"
					r -= 1
					c += 1
					break
				} else {
					c -= 1
				}
			}
		}
		if isExit {
			return false
		}
		counter += 1
	}

	return true
}

func getStartPos(input [][]string) (string, int, int) {
	direction := ""
	r, c := 0, 0
	for row, inp := range input {
		for col, char := range inp {
			if char == "^" {
				direction = "UP"
				input[row][col] = "X"
				r, c = row, col
				break
			} else if char == ">" {
				direction = "RIGHT"
				input[row][col] = "X"
				r, c = row, col
				break
			} else if char == "v" {
				direction = "DOWN"
				input[row][col] = "X"
				r, c = row, col
				break
			} else if char == "<" {
				direction = "LEFT"
				input[row][col] = "X"
				r, c = row, col
				break
			}
		}
	}

	return direction, r, c
}

func readFile(name string) [][]string {
	file, err := os.Open(name)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	var input [][]string
	r := bufio.NewReader(file)
	for {
		line, _, err := r.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		fields := strings.Split(string(line), "")
		input = append(input, fields)
	}
	return input
}

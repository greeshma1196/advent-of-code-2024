package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
	"strconv"
	"strings"
)

var left []int
var right []int

func main() {
	file, err := os.Open("input-day1.txt")
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
		fields := strings.Fields(string(line))
		lf, err := strconv.Atoi(fields[0])
		if err != nil {
			panic(err)
		}
		left = append(left, lf)

		rf, err := strconv.Atoi(fields[1])
		if err != nil {
			panic(err)
		}
		right = append(right, rf)
	}

	// part 1
	slices.Sort(left)
	slices.Sort(right)

	var diff []int

	for i, _ := range left {
		var d int
		l := left[i]
		r := right[i]
		if l >= r {
			d = l - r
		} else {
			d = r - l
		}

		diff = append(diff, d)
	}

	sDiff := 0

	for _, e := range diff {
		sDiff += e
	}

	fmt.Printf("Sum of Differences: %d\n", sDiff)

	// part 2
	mRight := map[int]int{}

	for _, e := range right {
		_, ok := mRight[e]
		if !ok {
			mRight[e] = 1
		} else {
			mRight[e] += 1
		}
	}

	mLeft := map[int]int{}
	for _, e := range left {
		_, ok := mLeft[e]
		if !ok {
			mLeft[e] = 0
		}
	}

	var sim []int
	for k, _ := range mLeft {
		n, ok := mRight[k]
		if ok {
			sim = append(sim, n*k)
		}
	}

	sSim := 0

	for _, e := range sim {
		sSim += e
	}

	fmt.Printf("Sum of Similarity: %d\n", sSim)

}

package day1

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
	"strconv"
	"strings"
)

func readFile(file *os.File) ([]int, []int) {
	r := bufio.NewReader(file)
	var left []int
	var right []int

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
	return left, right
}

func sumOfDiff(left []int, right []int) error {
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
	return nil
}

func sumOfSim(left []int, right []int) error {
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

	return nil
}

func ProcessAOCDay1(name string) error {

	file, err := os.Open(name)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	left, right := readFile(file)

	slices.Sort(left)
	slices.Sort(right)

	_ = sumOfDiff(left, right)
	_ = sumOfSim(left, right)

	return nil

}

package day2

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func readFile(name string) [][]int {
	file, err := os.Open(name)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	r := bufio.NewReader(file)
	var arr [][]int
	for {
		line, _, err := r.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		fields := strings.Fields(string(line))

		var a []int
		for _, f := range fields {
			n, err := strconv.Atoi(f)
			if err != nil {
				panic(err)
			}
			a = append(a, n)
		}

		arr = append(arr, a)
	}
	return arr
}

func checkSafety(arr []int) bool {

	isIncreasing, isDecreasing := false, false

	for i := 1; i < len(arr); i++ {
		d := arr[i] - arr[i-1]

		if d > 0 {
			isIncreasing = true
		} else if d < 0 {
			isDecreasing = true
		} else {
			return false
		}

		if isDecreasing && isIncreasing {
			return false
		}

		if d > 3 || d < -3 {
			return false
		}
	}

	return true
}

func ProcessAOCDay2(name string) error {

	arr := readFile(name)
	var unsafeArr [][]int
	res := 0
	for _, a := range arr {

		isSafe := checkSafety(a)

		if isSafe {
			res += 1
		} else {
			unsafeArr = append(unsafeArr, a)
		}
	}

	fmt.Printf("Number of safe reports: %d\n", res)
	fmt.Printf("Number of unsafe reports: %d\n", len(unsafeArr))

	for _, a := range unsafeArr {
		s := len(a)
		c := 0

		for c < s {
			var b []int
			for i, e := range a {
				if i != c {
					b = append(b, e)
				}
			}

			isSafe := checkSafety(b)
			if isSafe {
				res += 1
				break
			}

			c += 1
		}
	}

	fmt.Printf("Number of safe reports after using dampner: %d\n", res)

	return nil
}

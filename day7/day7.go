package day7

import (
	"bufio"
	"container/list"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var testValue map[string]string

func ComputeAOCDay7_1(name string) {
	testValue = readFile(name)

	res := 0

	type State struct {
		currentRes int
		idx        int
		exp        string
	}

	for k, v := range testValue {
		tVal, _ := strconv.Atoi(k)
		nums := strings.Split(v, " ")

		n0, _ := strconv.Atoi(nums[0])
		queue := list.New()
		queue.PushBack(State{currentRes: n0, idx: 0, exp: nums[0]})

		for queue.Len() > 0 {
			ele := queue.Front()
			state := ele.Value.(State)
			queue.Remove(ele)

			if state.idx == len(nums)-1 {
				if state.currentRes == tVal {
					res += tVal
					break
				}
				continue
			}
			nextNum, _ := strconv.Atoi(nums[state.idx+1])

			addRes := state.currentRes + nextNum
			addExp := fmt.Sprintf("(%s + %d)", state.exp, nextNum)
			queue.PushBack(State{currentRes: addRes, idx: state.idx + 1, exp: addExp})

			mulRes := state.currentRes * nextNum
			mulExp := fmt.Sprintf("(%s * %d)", state.exp, nextNum)
			queue.PushBack(State{currentRes: mulRes, idx: state.idx + 1, exp: mulExp})
		}

	}

	fmt.Printf("Result of day 7 part 1: %d\n", res)
}

func ComputeAOCDay7_2(name string) {
	testValue = readFile(name)

	res := 0

	type State struct {
		currentRes string
		idx        int
		exp        string
	}

	for k, v := range testValue {
		tVal := k
		nums := strings.Split(v, " ")

		queue := list.New()
		queue.PushBack(State{currentRes: nums[0], idx: 0, exp: nums[0]})

		for queue.Len() > 0 {
			ele := queue.Front()
			state := ele.Value.(State)
			queue.Remove(ele)

			if state.idx == len(nums)-1 {
				if state.currentRes == tVal {
					val, _ := strconv.Atoi(tVal)
					res += val
					break
				}
				continue
			}

			currResStr := state.currentRes
			currRes, _ := strconv.Atoi(currResStr)

			nextNumStr := nums[state.idx+1]
			nextNum, _ := strconv.Atoi(nextNumStr)

			concatRes := currResStr + nextNumStr
			concatExp := fmt.Sprintf("(%s || %d)", state.exp, nextNum)
			queue.PushBack(State{currentRes: concatRes, idx: state.idx + 1, exp: concatExp})

			addRes := currRes + nextNum
			addExp := fmt.Sprintf("(%s + %d)", state.exp, nextNum)
			queue.PushBack(State{currentRes: strconv.Itoa(addRes), idx: state.idx + 1, exp: addExp})

			mulRes := currRes * nextNum
			mulExp := fmt.Sprintf("(%s * %d)", state.exp, nextNum)
			queue.PushBack(State{currentRes: strconv.Itoa(mulRes), idx: state.idx + 1, exp: mulExp})
		}

	}

	fmt.Printf("Result of day 7 part 2: %d\n", res)
}

func readFile(name string) map[string]string {
	t := make(map[string]string)
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
		testVal, nums, _ := strings.Cut(string(line), ": ")
		t[testVal] = nums
	}

	return t
}

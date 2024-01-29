package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func isTermianlState(arr []int) bool {
	for _, n := range arr {
		if n != 0 {
			return false
		}
	}
	return true
}

func nextValue(arr []int) int {
	if isTermianlState(arr) {
		return 0
	} else {
		diffArr := make([]int, 0)
		for i := 1; i < len(arr); i++ {
			diffArr = append(diffArr, arr[i]-arr[i-1])
		}
		return arr[0] - nextValue(diffArr)
	}
	return 0
}

func readInput() [][]int {
	f, err := os.Open("Day09/input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)
	inputs := [][]int{}
	for scanner.Scan() {
		temp := strings.Split(scanner.Text(), " ")
		input := []int{}
		for _, s := range temp {
			x, _ := strconv.Atoi(s)
			input = append(input, x)
		}
		inputs = append(inputs, input)
	}
	return inputs
}

func sumOfNextValues(inputs [][]int) int {
	sum := 0
	for _, input := range inputs {
		sum += nextValue(input)
	}
	return sum
}

func main() {
	t1 := time.Now()
	fmt.Println(sumOfNextValues(readInput()))
	fmt.Println(time.Now().Sub(t1))
}

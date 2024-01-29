package main

import (
	"bufio"
	"fmt"
	"os"
)

type ValueNode struct {
	left  string
	right string
}

func readInput() (string, map[string]ValueNode) {
	f, err := os.Open("Day8/input.txt")
	if err != nil {
		panic(err)
	}
	m := map[string]ValueNode{}
	scanner := bufio.NewScanner(f)
	scanner.Scan()
	instructs := scanner.Text()
	scanner.Scan()
	for scanner.Scan() {
		temp := scanner.Text()
		m[temp[:3]] = ValueNode{left: temp[7:10], right: temp[12:15]}
	}
	return instructs, m
}

func part1() {
	steps := 0
	start := "AAA"
	terminal := "ZZZ"
	ins, m := readInput()
	nIns := len(ins)
	insRune := []rune(ins)
	currNode := start
	i := 0
	for currNode != terminal {
		if i == nIns {
			i = 0
		}
		if insRune[i] == 'L' {
			currNode = m[currNode].left
		} else if insRune[i] == 'R' {
			currNode = m[currNode].right
		}
		steps++
		i++
	}
	fmt.Println(steps)
}

func isTerminate(currs []string) bool {
	for _, curr := range currs {
		if curr[2] != 'Z' {
			return false
		}
	}
	return true
}

func part2() {
	steps := 0
	//starts := []string{"AAA", "SJA", "BXA", "QTA", "HCA", "LDA"}
	starts := []string{"LDA"}
	ins, m := readInput()
	nIns := len(ins)
	insRune := []rune(ins)
	currs := starts
	i := 0
	for !isTerminate(currs) {
		if i == nIns {
			i = 0
		}
		for j := 0; j < len(starts); j++ {
			if insRune[i] == 'L' {
				currs[j] = m[currs[j]].left
			} else if insRune[i] == 'R' {
				currs[j] = m[currs[j]].right
			}
		}
		steps++
		i++
		//fmt.Println(steps, currs)
	}
	fmt.Println(steps)
}

func main() {
	part2()
	//fmt.Println()
}

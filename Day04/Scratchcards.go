package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func twoListCreator(s string, rgx string) ([]int, []int) {
	splittedByBar := strings.Split(s, " | ")
	pattern := regexp.MustCompile(rgx)
	idx := pattern.FindStringIndex(splittedByBar[0])
	strWinningNumber := strings.Split(strings.Replace(strings.TrimSpace(splittedByBar[0][idx[1]:]), "  ", " ", -1), " ")
	strHavingNumber := strings.Split(strings.Replace(strings.TrimSpace(splittedByBar[1]), "  ", " ", -1), " ")
	winningNumber := make([]int, len(strWinningNumber))
	havingNumber := make([]int, len(strHavingNumber))
	for i := 0; i < len(strWinningNumber); i++ {
		x, _ := strconv.Atoi(strWinningNumber[i])
		winningNumber[i] = x
	}
	for i := 0; i < len(strHavingNumber); i++ {
		x, _ := strconv.Atoi(strHavingNumber[i])
		havingNumber[i] = x
	}
	return winningNumber, havingNumber
}

func contains(s *[]int, e int) bool {
	for i := 0; i < len(*s); i++ {
		if (*s)[i] == e {
			return true
		}
	}
	return false
}

func worthPoint(wn []int, hn []int) int {
	sum := 0
	for i := 0; i < len(hn); i++ {
		ok := contains(&wn, hn[i])
		if ok {
			if sum == 0 {
				sum = 1
			} else {
				sum *= 2
			}
		}
	}
	return sum
}

func numberOfMatchs(wn []int, hn []int) int {
	matchs := 0
	for i := 0; i < len(hn); i++ {
		ok := contains(&wn, hn[i])
		if ok {
			matchs++
		}
	}
	return matchs
}

func part1() {
	rgx := "Card[ ]*[0-9]*:"
	f, err := os.Open("Day04/input.txt")
	if err != nil {
		panic(err)
	}
	sum := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		s := scanner.Text()
		wn, hn := twoListCreator(s, rgx)
		sum += worthPoint(wn, hn)
	}
	fmt.Println(sum)
}

func part2() {
	var cardInstanceMap = map[int]int{}
	const NCards = 213
	rgx := "Card[ ]*[0-9]*:"
	f, err := os.Open("Day04/input.txt")
	if err != nil {
		panic(err)
	}
	sum := 0
	iterator := 1
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		cardInstanceMap[iterator] += 1
		s := scanner.Text()
		wn, hn := twoListCreator(s, rgx)
		nMatchs := numberOfMatchs(wn, hn)
		n := cardInstanceMap[iterator]
		for i := 1; i <= n; i++ {
			for j := 1; j <= nMatchs; j++ {
				cardInstanceMap[j+iterator]++
			}
		}
		iterator++
	}
	for i := 1; i <= NCards; i++ {
		sum += cardInstanceMap[i]
	}
	fmt.Println(sum)
}

func main() {
	part1()
}

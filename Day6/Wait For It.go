package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type race struct {
	time       int
	bestRecord int
}

func countWaysToWin(r race) int {
	sum := 0
	for i := 0; i < r.time; i++ {
		if i*(r.time-i) > r.bestRecord {
			sum++
		}
	}
	return sum
}

func fullCountWaysToWin(rs []race) int {
	res := 1
	for _, r := range rs {
		res *= countWaysToWin(r)
	}
	return res
}

func partA() {
	t1 := time.Now()
	pattern := regexp.MustCompile("[ ]+")
	f, err := os.Open("Day6/input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)
	scanner.Scan()
	s := strings.Split(pattern.ReplaceAllString(strings.TrimSpace(scanner.Text()[5:]), " "), " ")
	times := make([]int, len(s))
	for i, t := range s {
		times[i], _ = strconv.Atoi(t)
	}
	scanner.Scan()
	s = strings.Split(pattern.ReplaceAllString(strings.TrimSpace(scanner.Text()[9:]), " "), " ")
	distances := make([]int, len(s))
	for i, d := range s {
		distances[i], _ = strconv.Atoi(d)
	}
	races := []race{}
	for i := 0; i < len(times); i++ {
		races = append(races, race{times[i], distances[i]})
	}
	fmt.Println(fullCountWaysToWin(races))
	t2 := time.Now()
	fmt.Println(t2.Sub(t1))
}

func partB() {
	t1 := time.Now()
	pattern := regexp.MustCompile("[ ]+")
	f, err := os.Open("Day6/input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)
	scanner.Scan()
	s := pattern.ReplaceAllString(strings.TrimSpace(scanner.Text()[5:]), "")
	t, _ := strconv.Atoi(s)
	scanner.Scan()
	s = pattern.ReplaceAllString(strings.TrimSpace(scanner.Text()[9:]), "")
	d, _ := strconv.Atoi(s)
	race := race{t, d}
	fmt.Println(countWaysToWin(race))
	t2 := time.Now()
	fmt.Println(t2.Sub(t1))
}

func main() {
	partB()
}

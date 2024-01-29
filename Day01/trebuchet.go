package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var lettersToDigit = map[string]string{
	"zero":  "z0o",
	"one":   "o1e",
	"two":   "t2o",
	"three": "t3e",
	"four":  "f4r",
	"five":  "f5e",
	"six":   "s6x",
	"seven": "s7n",
	"eight": "e8t",
	"nine":  "n9e",
}

func replaceLetterNumberToDigits(s string) string {
	rgx := "zero|one|two|three|four|five|six|seven|eight|nine"
	pattern := regexp.MustCompile(rgx)
	letters := pattern.FindString(s)
	for letters != "" {
		s = strings.Replace(s, letters, lettersToDigit[letters], 1)
		letters = pattern.FindString(s)
	}
	return s
}

func extractValue(s string) int {
	r := []rune(s)
	start := 0
	end := len(r) - 1
	var first, last string
	var firstFound, lastFound bool
	for start <= end {
		if !firstFound && r[start] > 48 && r[start] < 58 {
			firstFound = true
			first = string(r[start])
		} else if !firstFound {
			start++
		}
		if !lastFound && r[end] > 48 && r[end] < 58 {
			lastFound = true
			last = string(r[end])
		} else if !lastFound {
			end--
		}
		if firstFound && lastFound {
			break
		}
	}
	res, _ := strconv.Atoi(first + last)
	return res
}

func part1() {
	f, err := os.Open("Day01/input.txt")
	if err != nil {
		panic(err)
	}
	sum := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		sum += extractValue(scanner.Text())
	}
	fmt.Println(sum)
}

func part2() {
	f, err := os.Open("Day01/input.txt")
	if err != nil {
		panic(err)
	}
	sum := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		sum += extractValue(replaceLetterNumberToDigits(scanner.Text()))
	}
	fmt.Println(sum)
}

func main() {
	part2()
}

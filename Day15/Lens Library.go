package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func hashAlgorithm(s string) int {
	current := 0
	chars := []rune(s)
	for _, char := range chars {
		current = ((current + int(char)) * 17) % 256
	}
	return current
}

func sumOfHash(strs []string) int {
	sum := 0
	for _, s := range strs {
		sum += hashAlgorithm(s)
	}
	return sum
}

func readData(path string) []string {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)
	scanner.Scan()
	strs := strings.Split(scanner.Text(), ",")
	return strs
}

func indexOfLabel(lenses []lens, label string) int {
	index := -1
	for i, l := range lenses {
		if l.label == label {
			index = i
		}
	}
	return index
}

func rmLens(lensesPtr *[]lens, index int) {
	copy((*lensesPtr)[index:], (*lensesPtr)[index+1:])
	(*lensesPtr)[len(*lensesPtr)-1] = lens{}
	*lensesPtr = (*lensesPtr)[:len(*lensesPtr)-1]
}

func initialSequence(strs []string) [][]lens {
	boxes := make([][]lens, 256)
	for _, s := range strs {
		pattern := regexp.MustCompile("[a-z]+")
		label := pattern.FindString(s)
		boxNo := hashAlgorithm(label)
		if strings.Contains(s, "=") {
			temp := string(s[len(s)-1])
			focalLength, _ := strconv.Atoi(temp)
			index := indexOfLabel(boxes[boxNo], label)
			if index == -1 {
				boxes[boxNo] = append(boxes[boxNo], lens{label: label, focalLength: focalLength})
			} else {
				boxes[boxNo][index] = lens{label: label, focalLength: focalLength}
			}
		} else {
			index := indexOfLabel(boxes[boxNo], label)
			if index != -1 {
				rmLens(&boxes[boxNo], index)
			}
		}
	}
	return boxes
}

func focusingPower(lenses []lens, boxNo int) int {
	sum := 0
	for i, lens := range lenses {
		sum += (i + 1) * lens.focalLength * boxNo
	}
	return sum
}

func part1() {
	strs := readData("Day15/input.txt")
	fmt.Println(sumOfHash(strs))
}

func part2() {
	sum := 0
	strs := readData("Day15/input.txt")
	boxes := initialSequence(strs)
	for i, box := range boxes {
		sum += focusingPower(box, i+1)
	}
	fmt.Println(sum)
}

func main() {
	part2()
}

type lens struct {
	label       string
	focalLength int
}

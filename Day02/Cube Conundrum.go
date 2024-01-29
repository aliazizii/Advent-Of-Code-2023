package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var cubeColorIDMap = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

type Cube struct {
	number int
	color  string
}

func TheGameNumber(s string) int {
	pattern := regexp.MustCompile("Game ([0-9]*):")
	subStr := pattern.FindStringSubmatch(s)
	n, err := strconv.Atoi(strings.TrimSpace(subStr[1]))
	if err != nil {
		panic(err)
	}
	return n
}

func splitGames(s string) []string {
	games := strings.Split(s, ";")
	pattern := regexp.MustCompile("Game ([0-9]*):")
	matchIndex := pattern.FindStringIndex(games[0])
	games[0] = games[0][matchIndex[1]+1:]
	for i := range games {
		games[i] = strings.TrimSpace(games[i])
		games[i] = strings.Replace(games[i], ", ", ",", -1)
	}
	return games
}

func cubeArrConvertor(s []string) []Cube {
	cubes := []Cube{}
	pattern := regexp.MustCompile("([0-9]*) (red|green|blue)")
	for i := 0; i < len(s); i++ {
		temp_cubes := strings.Split(s[i], ",")
		for j := 0; j < len(temp_cubes); j++ {
			cube := pattern.FindStringSubmatch(temp_cubes[j])
			number, err := strconv.Atoi(cube[1])
			if err != nil {
				panic(err)
			}
			color := cube[2]
			cubes = append(cubes, Cube{
				number: number,
				color:  color,
			})
		}
	}
	return cubes
}

func isValidGame(cubes []Cube, n int) int {
	for _, cube := range cubes {
		if cubeColorIDMap[cube.color] < cube.number {
			return 0
		}
	}
	return n
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func partBCalculator(cubes []Cube) int {
	multiply := 1
	maxRed := 0
	maxBlue := 0
	maxGreen := 0
	for _, cube := range cubes {
		if cube.color == "red" {
			maxRed = max(maxRed, cube.number)
		}
		if cube.color == "blue" {
			maxBlue = max(maxBlue, cube.number)
		}
		if cube.color == "green" {
			maxGreen = max(maxGreen, cube.number)
		}
	}
	multiply = maxGreen * maxBlue * maxRed
	return multiply
}

func part1() {
	f, err := os.Open("Day02/input.txt")
	if err != nil {
		panic(err)
	}
	sum := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		s := scanner.Text()
		n := TheGameNumber(s)
		arr := splitGames(s)
		cubes := cubeArrConvertor(arr)
		sum += isValidGame(cubes, n)
	}
	fmt.Println(sum)
}

func part2a() {
	f, err := os.Open("Day02/input.txt")
	if err != nil {
		panic(err)
	}
	sum := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		s := scanner.Text()
		arr := splitGames(s)
		cubes := cubeArrConvertor(arr)
		sum += partBCalculator(cubes)
	}
	fmt.Println(sum)
}

func main() {
	part1()
}

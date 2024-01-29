package main

import (
	"bufio"
	"fmt"
	"os"
)

func readData(path string) *[]string {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	rows := make([]string, 0)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		s := scanner.Text()
		rows = append(rows, s)
	}
	return &rows
}

func getCol(rows *[]string, colID int) []rune {
	out := make([]rune, 0)
	for _, row := range *rows {
		out = append(out, rune(row[colID]))
	}
	return out
}

func noGalaxyFinder(rows *[]string) (map[int]bool, map[int]bool) {
	rowNoGalaxy := make(map[int]bool)
	colNoGalaxy := make(map[int]bool)
	for i, s := range *rows {
		flag := true
		chars := []rune(s)
		for _, char := range chars {
			if char != '.' {
				flag = false
				break
			}
		}
		if flag {
			rowNoGalaxy[i] = true
		}
	}
	for i := 0; i < len((*rows)[0]); i++ {
		flag := true
		chars := getCol(rows, i)
		for _, char := range chars {
			if char != '.' {
				flag = false
				break
			}
		}
		if flag {
			colNoGalaxy[i] = true
		}
	}
	return rowNoGalaxy, colNoGalaxy
}

func extractGalaxy(rows *[]string) []galaxy {
	res := make([]galaxy, 0)
	for i, s := range *rows {
		for j, char := range s {
			if char == '#' {
				res = append(res, galaxy{
					x: i,
					y: j,
				})
			}
		}
	}
	return res
}

func abs(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}

func minAndMax(a, b int) (int, int) {
	if a > b {
		return b, a
	}
	return a, b
}

func distance(g1, g2 galaxy, rowNoGalaxy, colNoGalaxy *map[int]bool) int {
	d := 0
	d += abs(g1.x - g2.x)
	d += abs(g1.y - g2.y)
	minX, maxX := minAndMax(g1.x, g2.x)
	minY, maxY := minAndMax(g1.y, g2.y)
	for i := minX + 1; i < maxX; i++ {
		if (*rowNoGalaxy)[i] {
			d += 999999
		}
	}
	for i := minY + 1; i < maxY; i++ {
		if (*colNoGalaxy)[i] {
			d += 999999
		}
	}
	return d
}

func sumOfDistance(galaxies []galaxy, rowNoGalaxy, colNoGalaxy *map[int]bool) int {
	sum := 0
	for i, g1 := range galaxies {
		for _, g2 := range galaxies[i+1:] {
			sum += distance(g1, g2, rowNoGalaxy, colNoGalaxy)
		}
	}
	return sum
}

func part1() {
	rows := readData("Day11/input.txt")
	galaxies := extractGalaxy(rows)
	rowNoGalaxy, colNoGalaxy := noGalaxyFinder(rows)
	fmt.Println(sumOfDistance(galaxies, &rowNoGalaxy, &colNoGalaxy))
}

func main() {
	part1()
}

type galaxy struct {
	x int
	y int
}

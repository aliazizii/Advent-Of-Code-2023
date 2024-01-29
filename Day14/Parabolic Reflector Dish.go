package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"slices"
	"strings"
)

func roundRockCounter(chars []rune) []int {
	rocks := make([]int, 0)
	tempStringArr := make([]string, 0)
	for _, char := range chars {
		tempStringArr = append(tempStringArr, string(char))
	}
	splitBySharp := strings.Split(strings.Join(tempStringArr, ""), "#")
	for _, s := range splitBySharp {
		rocks = append(rocks, strings.Count(s, "O"))
	}
	return rocks
}

func getCol(rows [][]rune, colID int) []rune {
	out := make([]rune, 0)
	for _, row := range rows {
		out = append(out, row[colID])
	}
	return out
}

func getRow(rows [][]rune, rowID int) []rune {
	return rows[rowID]
}

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

func tiltingColToNorthWest(chars []rune) []rune {
	rocks := roundRockCounter(chars)
	modifiedChars := make([]rune, len(chars))
	j := 0
	cnt := 0
	for i, char := range chars {
		if char == '#' {
			j++
			cnt = 0
			modifiedChars[i] = '#'
		} else {
			if cnt < rocks[j] {
				modifiedChars[i] = 'O'
				cnt++
			} else {
				modifiedChars[i] = '.'
			}
		}
	}
	return modifiedChars
}

func tiltingColToSouthEast(chars []rune) []rune {
	rocks := roundRockCounter(chars)
	slices.Reverse(rocks)
	slices.Reverse(chars)
	modifiedChars := make([]rune, len(chars))
	j := 0
	cnt := 0
	for i, char := range chars {
		if char == '#' {
			j++
			cnt = 0
			modifiedChars[i] = '#'
		} else {
			if cnt < rocks[j] {
				modifiedChars[i] = 'O'
				cnt++
			} else {
				modifiedChars[i] = '.'
			}
		}
	}
	slices.Reverse(modifiedChars)
	return modifiedChars
}

func tiltingMapToNorth(problemMapIn [][]rune) [][]rune {
	problemMapOut := make([][]rune, len(problemMapIn))
	for i := 0; i < len(problemMapIn[0]); i++ {
		col := getCol(problemMapIn, i)
		modifiedCol := tiltingColToNorthWest(col)
		for j := 0; j < len(problemMapIn); j++ {
			problemMapOut[j] = append(problemMapOut[j], modifiedCol[j])
		}
	}
	return problemMapOut
}

func tiltingMapToEast(problemMapIn [][]rune) [][]rune {
	problemMapOut := make([][]rune, len(problemMapIn))
	for i := 0; i < len(problemMapIn); i++ {
		row := getRow(problemMapIn, i)
		modifiedRow := tiltingColToSouthEast(row)
		problemMapOut[i] = modifiedRow
	}
	return problemMapOut
}

func tiltingMapToSouth(problemMapIn [][]rune) [][]rune {
	problemMapOut := make([][]rune, len(problemMapIn))
	for i := 0; i < len(problemMapIn[0]); i++ {
		col := getCol(problemMapIn, i)
		modifiedCol := tiltingColToSouthEast(col)
		for j := 0; j < len(problemMapIn); j++ {
			problemMapOut[j] = append(problemMapOut[j], modifiedCol[j])
		}
	}
	return problemMapOut
}

func tiltingMapToWest(problemMapIn [][]rune) [][]rune {
	problemMapOut := make([][]rune, len(problemMapIn))
	for i := 0; i < len(problemMapIn); i++ {
		row := getRow(problemMapIn, i)
		modifiedRow := tiltingColToNorthWest(row)
		problemMapOut[i] = modifiedRow
	}
	return problemMapOut
}

func part1() {
	rows := readData("Day14/input.txt")
	problemMap := make([][]rune, len(*rows))
	for i, s := range *rows {
		for _, char := range s {
			problemMap[i] = append(problemMap[i], char)
		}
	}
	//minimumIteration := 0
	first := problemMap
	for i := 0; i < 90; i++ {
		a := problemMap
		problemMap = tiltingMapToNorth(problemMap)
		problemMap = tiltingMapToWest(problemMap)
		problemMap = tiltingMapToSouth(problemMap)
		problemMap = tiltingMapToEast(problemMap)
		b := problemMap
		if reflect.DeepEqual(a, b) || reflect.DeepEqual(first, b) {
			//minimumIteration = i
			fmt.Println("kir", i)
			break

		}
		fmt.Println("i=", i)
		printRes(problemMap)
	}
	//first = problemMap
	//for i := 0; i < 100; i++ {
	//	problemMap = tiltingMapToNorth(problemMap)
	//	problemMap = tiltingMapToWest(problemMap)
	//	problemMap = tiltingMapToSouth(problemMap)
	//	problemMap = tiltingMapToEast(problemMap)
	//	b := problemMap
	//	if reflect.DeepEqual(first, b) {
	//		//minimumIteration = i
	//		fmt.Println("kir", i)
	//		break
	//
	//	}
	//	fmt.Println("i=", i)
	//	printRes(problemMap)
	//
	//}
	printRes(problemMap)

}

func printRes(problemMap [][]rune) {
	sum := 0
	n := len(problemMap)
	for i, cells := range problemMap {
		for _, c := range cells {
			//fmt.Print(string(c))
			if c == 'O' {
				sum += n - i
			}
		}
		//fmt.Println()
	}
	fmt.Println(sum)
}

func main() {
	part1()
}

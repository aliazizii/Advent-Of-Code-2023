package main

import (
	"bufio"
	"fmt"
	"os"
)

func getCol(rows *[]string, colID int) []rune {
	out := make([]rune, 0)
	for _, row := range *rows {
		out = append(out, rune(row[colID]))
	}
	return out
}

func getRow(rows *[]string, rowID int) []rune {
	return []rune((*rows)[rowID])
}

func testEqual(a, b []rune) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func testEqualWithError(a, b []rune) int {
	err := 0
	for i, v := range a {
		if v != b[i] {
			err += 1
		}
	}
	return err
}

func isSymmetricFromCol(s *[]string, axis int, isLeft bool) bool {
	sumErr := 0
	//isSymmetric := true
	var l, r int
	if isLeft {
		l = 0
		r = (2 * axis) - 1
	} else {
		r = len((*s)[0]) - 1
		l = r - (2 * axis) + 1
	}
	for l <= r {
		leftCol := getCol(s, l)
		rightCol := getCol(s, r)
		//equality := testEqual(leftCol, rightCol)
		sumErr += testEqualWithError(leftCol, rightCol)
		//if !equality {
		//	isSymmetric = false
		//	break
		//} else {
		//	l++
		//	r--
		//}
		l++
		r--
	}
	if sumErr == 1 {
		return true
	}
	return false
	//return isSymmetric
}

func findColReflection(s *[]string) int {
	n := len((*s)[0])
	n /= 2
	flag := false
	for !flag && n > 0 {
		flag = isSymmetricFromCol(s, n, true)
		if flag {
			return n
		}
		flag = isSymmetricFromCol(s, n, false)
		if flag {
			return len((*s)[0]) - n
		}
		n--
	}
	return 0
}

func isSymmetricFromRow(s *[]string, axis int, isLeft bool) bool {
	sumErr := 0
	//isSymmetric := true
	var l, r int
	if isLeft {
		l = 0
		r = (2 * axis) - 1
	} else {
		r = len(*s) - 1
		l = r - (2 * axis) + 1
	}
	for l <= r {
		leftCol := getRow(s, l)
		rightCol := getRow(s, r)
		sumErr += testEqualWithError(leftCol, rightCol)
		//equality := testEqual(leftCol, rightCol)
		//if !equality {
		//	isSymmetric = false
		//	break
		//} else {
		//	l++
		//	r--
		//}
		l++
		r--
	}
	if sumErr == 1 {
		return true
	}
	return false
	//return isSymmetric
}

func findRowReflection(s *[]string) int {
	n := len(*s)
	n /= 2
	flag := false
	for !flag && n > 0 {
		flag = isSymmetricFromRow(s, n, true)
		if flag {
			return n
		}
		flag = isSymmetricFromRow(s, n, false)
		if flag {
			return len(*s) - n
		}
		n--
	}
	return 0
}

func readData(path string) [][]string {
	res := [][]string{}
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)
	pattern := []string{}
	for scanner.Scan() {
		s := scanner.Text()
		if s == "" {
			res = append(res, pattern)
			pattern = nil
			continue
		}
		pattern = append(pattern, s)
	}
	res = append(res, pattern)
	return res
}

func part1() {
	sum := 0
	res := readData("Day13/input.txt")
	for _, pattern := range res {
		colReflection := findColReflection(&pattern)
		if colReflection > 0 {
			sum += colReflection
			continue
		}
		rowReflection := findRowReflection(&pattern)
		sum += 100 * rowReflection
	}
	fmt.Println(sum)
}

func main() {
	part1()
}

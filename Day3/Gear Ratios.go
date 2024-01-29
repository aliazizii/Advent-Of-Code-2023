package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type gear struct {
	row int
	col int
}

var gearToEngineNumberMap = map[gear][]int{}

func readData(path string) *[]string {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	rows := []string{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		s := scanner.Text()
		rows = append(rows, s)
	}
	return &rows
}

func borderDataByDot(ps *[]string) {
	for i := 0; i < len(*ps); i++ {
		s := "."
		s += (*ps)[i]
		s += "."
		(*ps)[i] = s
	}

	temp_string := ""
	for i := 0; i < len((*ps)[0]); i++ {
		temp_string += "."
	}
	*ps = append((*ps)[:1], *ps...)
	(*ps)[0] = temp_string
	temp_string = ""
	for i := 0; i < len((*ps)[len(*ps)-1]); i++ {
		temp_string += "."
	}
	*ps = append(*ps, temp_string)
}

func isValidPartnumber(ps *[]string, i int, idx []int) bool {
	rgx := "[^0-9.]"
	top := (*ps)[i-1][idx[0]-1 : idx[1]+1]
	bottom := (*ps)[i+1][idx[0]-1 : idx[1]+1]
	left := (*ps)[i][idx[0]-1 : idx[0]]
	right := (*ps)[i][idx[1] : idx[1]+1]
	flag, _ := regexp.MatchString(rgx, top)
	if flag {
		return true
	}
	flag, _ = regexp.MatchString(rgx, bottom)
	if flag {
		return true
	}
	flag, _ = regexp.MatchString(rgx, left)
	if flag {
		return true
	}
	flag, _ = regexp.MatchString(rgx, right)
	if flag {
		return true
	}
	return false
}

func addToMap(i int, idx []int, engineNumber int, indices []int, leftOrRight bool) {
	g := gear{
		row: i,
		col: indices[0] + idx[0] - 1,
	}
	if leftOrRight {
		g = gear{
			row: i,
			col: indices[0],
		}
	}

	_, ok := gearToEngineNumberMap[g]
	if ok {
		gearToEngineNumberMap[g] = append(gearToEngineNumberMap[g], engineNumber)
	} else {
		gearToEngineNumberMap[g] = []int{engineNumber}
	}
}

func updateMap(ps *[]string, i int, idx []int, engineNumber int) {
	rgx := "[*]"
	top := (*ps)[i-1][idx[0]-1 : idx[1]+1]
	bottom := (*ps)[i+1][idx[0]-1 : idx[1]+1]
	left := (*ps)[i][idx[0]-1 : idx[0]]
	right := (*ps)[i][idx[1] : idx[1]+1]
	pattern := regexp.MustCompile(rgx)
	allSubstringMatches := pattern.FindAllStringIndex(top, -1)
	for _, indices := range allSubstringMatches {
		addToMap(i-1, idx, engineNumber, indices, false)
	}
	allSubstringMatches = pattern.FindAllStringIndex(bottom, -1)
	for _, indices := range allSubstringMatches {
		addToMap(i+1, idx, engineNumber, indices, false)
	}
	if left == "*" {
		addToMap(i, idx, engineNumber, []int{idx[0] - 1}, true)
	}
	if right == "*" {
		addToMap(i, idx, engineNumber, []int{idx[1]}, true)
	}
}

func iterateOverEngineNumbers(ps *[]string) {
	pattern := regexp.MustCompile("\\d+")
	for i := 1; i < len(*ps)-1; i++ {
		allSubstringMatches := pattern.FindAllStringIndex((*ps)[i], -1)
		for _, idx := range allSubstringMatches {
			engineNumber, _ := strconv.Atoi((*ps)[i][idx[0]:idx[1]])
			updateMap(ps, i, idx, engineNumber)
		}
	}
}

func partB() {
	ps := readData("Day3/input.txt")
	borderDataByDot(ps)
	iterateOverEngineNumbers(ps)
	sum := 0
	for _, ints := range gearToEngineNumberMap {
		if len(ints) == 2 {
			sum += ints[0] * ints[1]
		}
	}
	fmt.Println(sum)
}

func sumOfAllEnginePartNumber(ps *[]string) int {
	sum := 0
	pattern := regexp.MustCompile("\\d+")
	for i := 1; i < len(*ps)-1; i++ {
		allSubstringMatches := pattern.FindAllStringIndex((*ps)[i], -1)
		for _, idx := range allSubstringMatches {
			isValid := isValidPartnumber(ps, i, idx)
			if isValid {
				engineNumber, _ := strconv.Atoi((*ps)[i][idx[0]:idx[1]])
				sum += engineNumber
			}
		}
	}
	return sum
}

func partA() {
	ps := readData("Day3/input.txt")
	borderDataByDot(ps)
	sum := sumOfAllEnginePartNumber(ps)
	fmt.Println(sum)
}

func main() {
	//s := ".467*.+.*114.."
	//s = s[2:]
	//pattern := regexp.MustCompile("[*]")
	//allSubstringMatches := pattern.FindAllStringIndex(s, -1)
	//fmt.Println(allSubstringMatches)
	//ss := s[2:3]
	//fmt.Printf("%T, %v", ss, ss)
	partB()
	//for g, ints := range gearToEngineNumberMap {
	//	fmt.Println(g, " ", ints)
	//}
	//
	//for _, row := range *ps {
	//	fmt.Println(row)
	//}
}

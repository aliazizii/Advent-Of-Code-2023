package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
	"time"
)

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

func findStart(pipeMaze *[]string) (int, int) {
	for i, row := range *pipeMaze {
		for j, cell := range row {
			if cell == 'S' {
				return i, j
			}
		}
	}
	return 0, 0
}

func findSType(x, y int, pipeMazePtr *[]string) rune {
	top := (*pipeMazePtr)[x-1][y]
	bottom := (*pipeMazePtr)[x+1][y]
	left := (*pipeMazePtr)[x][y-1]
	right := (*pipeMazePtr)[x][y+1]
	topFlag := false
	bottomFlag := false
	leftFlag := false
	rightFlag := false
	if top != '.' && (top == '|' || top == '7' || top == 'F') {
		topFlag = true
	}
	if bottom != '.' && (bottom == '|' || bottom == 'J' || bottom == 'L') {
		bottomFlag = true
	}
	if left != '.' && (left == '-' || left == 'F' || left == 'L') {
		leftFlag = true
	}
	if right != '.' && (right == '-' || right == 'J' || right == '7') {
		rightFlag = true
	}
	if topFlag && bottomFlag {
		return '|'
	} else if topFlag && leftFlag {
		return 'J'
	} else if topFlag && rightFlag {
		return 'L'
	} else if rightFlag && leftFlag {
		return '-'
	} else if bottomFlag && rightFlag {
		return 'F'
	} else if bottomFlag && leftFlag {
		return '7'
	}
	return '.'
}

func twoConnector(p pipe, pipeMaze *[][]pipe) []pipe {
	connectors := make([]pipe, 2)
	if p.kind == '-' {
		connectors[0] = (*pipeMaze)[p.x][p.y+1]
		connectors[1] = (*pipeMaze)[p.x][p.y-1]
	} else if p.kind == '|' {
		connectors[0] = (*pipeMaze)[p.x-1][p.y]
		connectors[1] = (*pipeMaze)[p.x+1][p.y]
	} else if p.kind == '7' {
		connectors[1] = (*pipeMaze)[p.x][p.y-1]
		connectors[0] = (*pipeMaze)[p.x+1][p.y]
	} else if p.kind == 'F' {
		connectors[0] = (*pipeMaze)[p.x][p.y+1]
		connectors[1] = (*pipeMaze)[p.x+1][p.y]
	} else if p.kind == 'L' {
		connectors[1] = (*pipeMaze)[p.x-1][p.y]
		connectors[0] = (*pipeMaze)[p.x][p.y+1]
	} else if p.kind == 'J' {
		connectors[0] = (*pipeMaze)[p.x-1][p.y]
		connectors[1] = (*pipeMaze)[p.x][p.y-1]
	}
	return connectors
}

func stepsToFarthestPoint(x, y int, pipeMaze *[][]pipe) int {
	steps := 1
	tc := twoConnector((*pipeMaze)[x][y], pipeMaze)
	prevPipe := (*pipeMaze)[x][y]
	currPipe := tc[0]
	for currPipe != (*pipeMaze)[x][y] {
		tc = twoConnector(currPipe, pipeMaze)
		var tempPipe pipe
		for _, p := range tc {
			if p != prevPipe {
				tempPipe = p
			}
		}
		prevPipe = currPipe
		currPipe = tempPipe
		steps++
	}
	return steps / 2
}

func convertor(pipeMazePtr *[]string) [][]pipe {
	pipeMaze := [][]pipe{}
	for i, s := range *pipeMazePtr {
		pipeRow := make([]pipe, 0)
		for j, pipeKind := range s {
			pipeRow = append(pipeRow, pipe{i, j, pipeKind})
		}
		pipeMaze = append(pipeMaze, pipeRow)
	}
	return pipeMaze
}

func loopPipe(x, y int, pipeMaze *[][]pipe) []pipe {
	loop := make([]pipe, 0)
	tc := twoConnector((*pipeMaze)[x][y], pipeMaze)
	prevPipe := (*pipeMaze)[x][y]
	loop = append(loop, prevPipe)
	currPipe := tc[0]
	for currPipe != (*pipeMaze)[x][y] {
		tc = twoConnector(currPipe, pipeMaze)
		var tempPipe pipe
		for _, p := range tc {
			if p != prevPipe {
				tempPipe = p
			}
		}
		loop = append(loop, currPipe)
		prevPipe = currPipe
		currPipe = tempPipe
	}
	return loop
}

func findLTRBOfCell(loop *[]pipe, pipeMaze *[][]pipe, p pipe) []pipe {
	res := []pipe{}
	for i := p.x - 1; i > 0; i-- {
		//p.y
		if slices.Contains(*loop, (*pipeMaze)[i][p.y]) {
			res = append(res, (*pipeMaze)[i][p.y])
			break
		}
	}
	for i := p.y - 1; i > 0; i-- {
		if slices.Contains(*loop, (*pipeMaze)[p.x][i]) {
			res = append(res, (*pipeMaze)[p.x][i])
			break
		}
	}
	for i := p.x + 1; i < len((*pipeMaze)[p.x]); i++ {
		if slices.Contains(*loop, (*pipeMaze)[i][p.y]) {
			res = append(res, (*pipeMaze)[i][p.y])
			break
		}
	}
	for i := p.y + 1; i < len(*pipeMaze); i++ {
		if slices.Contains(*loop, (*pipeMaze)[p.x][i]) {
			res = append(res, (*pipeMaze)[p.x][i])
			break
		}
	}
	return res
}

func isClockWise(loop *[]pipe, ltrb []pipe) bool {
	flag := false
	ltrbOrder := ""
	surroundings := make([]surrounding, 4)
	surroundings[0] = surrounding{
		index: slices.Index(*loop, ltrb[0]),
		ltrb:  "l",
	}
	surroundings[1] = surrounding{
		index: slices.Index(*loop, ltrb[1]),
		ltrb:  "t",
	}
	surroundings[2] = surrounding{
		index: slices.Index(*loop, ltrb[2]),
		ltrb:  "r",
	}
	surroundings[3] = surrounding{
		index: slices.Index(*loop, ltrb[3]),
		ltrb:  "b",
	}
	sort.Slice(surroundings, func(i, j int) bool {
		return surroundings[i].index < surroundings[j].index
	})
	for _, s := range surroundings {
		ltrbOrder += s.ltrb
	}
	if ltrbOrder == "ltrb" || ltrbOrder == "trbl" || ltrbOrder == "rblt" || ltrbOrder == "bltr" {
		flag = true
	}
	return flag
}

func numberOfCellInsideLoop(loop *[]pipe, pipeMaze *[][]pipe) int {
	sum := 0
	for i := 1; i < len(*pipeMaze)-1; i++ {
		for j := 1; j < len((*pipeMaze)[i])-1; j++ {
			if !slices.Contains(*loop, (*pipeMaze)[i][j]) {
				ltrb := findLTRBOfCell(loop, pipeMaze, (*pipeMaze)[i][j])
				if len(ltrb) < 4 {
					continue
				} else {
					flag := isClockWise(loop, ltrb)
					if flag {
						sum++
					}
				}
			}
		}
	}
	return sum
}

func part1() {
	pipeMazePtr := readData("Day10/input.txt")
	borderDataByDot(pipeMazePtr)
	x, y := findStart(pipeMazePtr)
	pipeTemp := findSType(x, y, pipeMazePtr)
	runesTemp := []rune((*pipeMazePtr)[x])
	runesTemp[y] = pipeTemp
	(*pipeMazePtr)[x] = string(runesTemp)
	pipeMaze := convertor(pipeMazePtr)
	fmt.Println(stepsToFarthestPoint(x, y, &pipeMaze))
}

func part2() {
	pipeMazePtr := readData("Day10/input.txt")
	borderDataByDot(pipeMazePtr)
	x, y := findStart(pipeMazePtr)
	pipeTemp := findSType(x, y, pipeMazePtr)
	runesTemp := []rune((*pipeMazePtr)[x])
	runesTemp[y] = pipeTemp
	(*pipeMazePtr)[x] = string(runesTemp)
	pipeMaze := convertor(pipeMazePtr)
	loop := loopPipe(x, y, &pipeMaze)
	fmt.Println(numberOfCellInsideLoop(&loop, &pipeMaze))
}

func main() {
	t1 := time.Now()
	part2()
	fmt.Println(time.Now().Sub(t1))
}

type pipe struct {
	x    int
	y    int
	kind rune
}

type surrounding struct {
	index int
	ltrb  string
}

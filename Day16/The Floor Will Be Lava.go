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

func convertor(rows *[]string) *[][]tile {
	problemMap := [][]tile{}
	for i, s := range *rows {
		tileRow := make([]tile, 0)
		upDownBorder := false
		if i == 0 || i == len(*rows)-1 {
			upDownBorder = true
		}
		for j, t := range s {
			leftRightBorder := false
			if j == 0 || j == len(s)-1 {
				leftRightBorder = true
			}
			if upDownBorder || leftRightBorder {
				tileRow = append(tileRow, tile{
					x:          i,
					y:          j,
					tileKind:   t,
					isInBorder: true,
				})
			} else {
				tileRow = append(tileRow, tile{
					x:        i,
					y:        j,
					tileKind: t,
				})
			}
		}
		problemMap = append(problemMap, tileRow)
	}
	return &problemMap
}

func (t tile) nextTiles() []tile {
	nexts := make([]tile, 0)
	switch t.currDir {
	case UP:
		switch t.tileKind {
		case '.', '|':
			nexts = append(nexts, tile{x: t.x - 1, y: t.y, currDir: UP})
			return nexts
		case '\\':
			nexts = append(nexts, tile{x: t.x, y: t.y - 1, currDir: LEFT})
			return nexts
		case '/':
			nexts = append(nexts, tile{x: t.x, y: t.y + 1, currDir: RIGHT})
			return nexts
		case '-':
			nexts = append(nexts, tile{x: t.x, y: t.y - 1, currDir: LEFT})
			nexts = append(nexts, tile{x: t.x, y: t.y + 1, currDir: RIGHT})
			return nexts
		}
	case RIGHT:
		switch t.tileKind {
		case '.', '-':
			nexts = append(nexts, tile{x: t.x, y: t.y + 1, currDir: RIGHT})
			return nexts
		case '\\':
			nexts = append(nexts, tile{x: t.x + 1, y: t.y, currDir: DOWN})
			return nexts
		case '/':
			nexts = append(nexts, tile{x: t.x - 1, y: t.y, currDir: UP})
			return nexts
		case '|':
			nexts = append(nexts, tile{x: t.x + 1, y: t.y, currDir: DOWN})
			nexts = append(nexts, tile{x: t.x - 1, y: t.y, currDir: UP})
			return nexts
		}
	case LEFT:
		switch t.tileKind {
		case '.', '-':
			nexts = append(nexts, tile{x: t.x, y: t.y - 1, currDir: LEFT})
			return nexts
		case '\\':
			nexts = append(nexts, tile{x: t.x - 1, y: t.y, currDir: UP})
			return nexts
		case '/':
			nexts = append(nexts, tile{x: t.x + 1, y: t.y, currDir: DOWN})
			return nexts
		case '|':
			nexts = append(nexts, tile{x: t.x + 1, y: t.y, currDir: DOWN})
			nexts = append(nexts, tile{x: t.x - 1, y: t.y, currDir: UP})
			return nexts
		}
	case DOWN:
		switch t.tileKind {
		case '.', '|':
			nexts = append(nexts, tile{x: t.x + 1, y: t.y, currDir: DOWN})
			return nexts
		case '\\':
			nexts = append(nexts, tile{x: t.x, y: t.y + 1, currDir: RIGHT})
			return nexts
		case '/':
			nexts = append(nexts, tile{x: t.x, y: t.y - 1, currDir: LEFT})
			return nexts
		case '-':
			nexts = append(nexts, tile{x: t.x, y: t.y - 1, currDir: LEFT})
			nexts = append(nexts, tile{x: t.x, y: t.y + 1, currDir: RIGHT})
			return nexts
		}
	}
	return nexts
}

func iterateOverProblemMap(problemMap *[][]tile, startTile tile) {
	xStart := startTile.x
	yStart := startTile.y
	currDir := startTile.currDir
	if (*problemMap)[xStart][yStart].isInBorder {
		return
	}
	switch currDir {
	case LEFT:
		if (*problemMap)[xStart][yStart].leftSeen {
			return
		} else {
			(*problemMap)[xStart][yStart].leftSeen = true
		}
	case RIGHT:
		if (*problemMap)[xStart][yStart].rightSeen {
			return
		} else {
			(*problemMap)[xStart][yStart].rightSeen = true
		}
	case UP:
		if (*problemMap)[xStart][yStart].upSeen {
			return
		} else {
			(*problemMap)[xStart][yStart].upSeen = true
		}
	case DOWN:
		if (*problemMap)[xStart][yStart].downSeen {
			return
		} else {
			(*problemMap)[xStart][yStart].downSeen = true
		}
	}
	nexts := startTile.nextTiles()
	for _, next := range nexts {
		next.tileKind = (*problemMap)[next.x][next.y].tileKind
		iterateOverProblemMap(problemMap, next)
	}
}

func countEnergizedTiles(problemMap *[][]tile) int {
	sum := 0
	for _, tiles := range *problemMap {
		for _, t := range tiles {
			if t.rightSeen || t.downSeen || t.leftSeen || t.upSeen {
				sum++
			}
		}
	}
	return sum
}

func part1() {
	rows := readData("Day16/input.txt")
	borderDataByDot(rows)
	problemMap := convertor(rows)
	startTile := tile{x: 1, y: 4, currDir: DOWN, tileKind: (*problemMap)[1][4].tileKind}
	iterateOverProblemMap(problemMap, startTile)
	sum := countEnergizedTiles(problemMap)
	fmt.Println(sum)
}

func part2() {
	rows := readData("Day16/input.txt")
	borderDataByDot(rows)
	problemMap := convertor(rows)
	startTiles := make([]tile, 0)
	for _, t := range (*problemMap)[1][1 : len((*problemMap)[0])-1] {
		temp := t
		temp.currDir = DOWN
		startTiles = append(startTiles, temp)
	}
	for _, t := range (*problemMap)[len(*problemMap)-2][1 : len((*problemMap)[0])-1] {
		temp := t
		temp.currDir = UP
		startTiles = append(startTiles, temp)
	}
	for i := 1; i < len(*problemMap)-1; i++ {
		temp := (*problemMap)[i][1]
		temp.currDir = RIGHT
		startTiles = append(startTiles, temp)
		temp = (*problemMap)[i][len(*(problemMap))-2]
		temp.currDir = LEFT
		startTiles = append(startTiles, temp)
	}

	maxSum := 0
	for _, startTile := range startTiles {
		problemMap := convertor(rows)
		iterateOverProblemMap(problemMap, startTile)
		sum := countEnergizedTiles(problemMap)
		maxSum = max(sum, maxSum)
	}
	fmt.Println(maxSum)
}

func main() {
	//
	for {

	}
}

const (
	UP = iota + 1
	RIGHT
	DOWN
	LEFT
)

type tile struct {
	x          int
	y          int
	tileKind   rune
	isInBorder bool
	leftSeen   bool
	rightSeen  bool
	upSeen     bool
	downSeen   bool
	currDir    int
}

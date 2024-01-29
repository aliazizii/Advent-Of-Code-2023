package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

type mapper struct {
	destRangeStart   int
	sourceRangeStart int
	lengthRange      int
}

type listMapper []mapper

func mapSourceTodest(i int, lmapper listMapper) int {
	dest := i
	for _, m := range lmapper {
		if i >= m.sourceRangeStart && i < m.sourceRangeStart+m.lengthRange {
			dest = m.destRangeStart + (i - m.sourceRangeStart)
			break
		}
	}
	return dest
}

func locOfSeed(i int, mappers []listMapper) int {
	var loc = i
	for _, lm := range mappers {
		loc = mapSourceTodest(loc, lm)
	}
	return loc
}

func minLoc(seeds []int, mappers []listMapper) int {
	var minLocation = math.MaxInt
	for _, seed := range seeds {
		loc := locOfSeed(seed, mappers)
		minLocation = min(minLocation, loc)
	}
	return minLocation
}

func main() {
	t1 := time.Now()
	f, err := os.Open("Day5/input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)
	scanner.Scan()
	s := scanner.Text()
	s = s[7:]
	splitS := strings.Split(s, " ")
	seeds := make([]int, 0)
	for _, split := range splitS {
		x, err := strconv.Atoi(split)
		if err != nil {
			panic(err)
		}
		seeds = append(seeds, x)
	}
	scanner.Scan()
	mappers := []listMapper{}
	for scanner.Scan() {
		lmapper := listMapper{}
		for scanner.Scan() {
			if scanner.Text() == "" {
				break
			}
			numbers := strings.Split(scanner.Text(), " ")
			mapperData := make([]int, 3)
			for i, number := range numbers {
				x, err := strconv.Atoi(number)
				if err != nil {
					panic(err)
				}
				mapperData[i] = x
			}
			m := mapper{destRangeStart: mapperData[0], sourceRangeStart: mapperData[1], lengthRange: mapperData[2]}
			lmapper = append(lmapper, m)
		}
		mappers = append(mappers, lmapper)
	}
	arr := []int{}
	seedRanges := []seedRange{
		{565778304, 341771914},
		{1736484943, 907429186},
		{3928647431, 87620927},
		{311881326, 149873504},
		{1588660730, 119852039},
		{1422681143, 13548942},
		{1095049712, 216743334},
		{3671387621, 186617344},
		{3055786218, 213191880},
		{2783359478, 44001797},
	}

	for _, sr := range seedRanges {
		minL := math.MaxInt
		for i := sr.start; i < sr.start+sr.length; i++ {
			minL = min(minL, locOfSeed(i, mappers))
		}
		arr = append(arr, minL)
	}
	minL := math.MaxInt
	for _, loc := range arr {
		minL = min(minL, loc)
	}
	fmt.Println(minL)
	t2 := time.Now()
	totalTime := t2.Sub(t1)
	fmt.Println("Total time of execution program: ", totalTime)
}

type seedRange struct {
	start  int
	length int
}

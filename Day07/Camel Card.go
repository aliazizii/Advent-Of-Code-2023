package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	HIGHCARD = iota + 1
	ONEPAIR
	TWOPAIR
	THREEKIND
	FULLHOUSE
	FOURKIND
	FIVEKIND
)

func readInput() ([]string, map[string]int) {
	f, err := os.Open("Day07/input.txt")
	if err != nil {
		panic(err)
	}
	handsToBid := make(map[string]int)
	handsList := make([]string, 0)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		handsList = append(handsList, line[0])
		bid, _ := strconv.Atoi(line[1])
		handsToBid[line[0]] = bid
	}
	return handsList, handsToBid
}

func HandsDescriptor(s string) []int {
	m := make(map[rune]int)
	res := make([]int, 0)
	r := []rune(s)
	for _, rns := range r {
		m[rns]++
	}
	for _, v := range m {
		res = append(res, v)
	}
	sort.Ints(res)
	return res
}

func HandsDescriptorPartB(s string) []int {
	m := make(map[rune]int)
	res := make([]int, 0)
	r := []rune(s)
	for _, rns := range r {
		if rns != 'J' {
			m[rns]++
		}
	}
	for _, v := range m {
		res = append(res, v)
	}
	sort.Ints(res)
	x := strings.Count(s, "J")
	if x == 5 {
		res = append(res, 5)
		return res
	}
	res[len(res)-1] += x
	return res
}

func testEq(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func fiveOfaKindChecker(l []int) bool {
	return testEq(l, []int{5})
}

func fourOfaKindChecker(l []int) bool {
	return testEq(l, []int{1, 4})
}

func fullHouseChecker(l []int) bool {
	return testEq(l, []int{2, 3})
}

func threeOfaKindChecker(l []int) bool {
	return testEq(l, []int{1, 1, 3})
}

func twoPairChecker(l []int) bool {
	return testEq(l, []int{1, 2, 2})
}

func onePairChecker(l []int) bool {
	return testEq(l, []int{1, 1, 1, 2})
}

func highCardChecker(l []int) bool {
	return testEq(l, []int{1, 1, 1, 1, 1})
}

func typeOfHand(s string) int {
	var ok bool
	l := HandsDescriptorPartB(s)
	ok = fiveOfaKindChecker(l)
	if ok {
		return FIVEKIND
	}
	ok = fourOfaKindChecker(l)
	if ok {
		return FOURKIND
	}
	ok = fullHouseChecker(l)
	if ok {
		return FULLHOUSE
	}
	ok = threeOfaKindChecker(l)
	if ok {
		return THREEKIND
	}
	ok = twoPairChecker(l)
	if ok {
		return TWOPAIR
	}
	ok = onePairChecker(l)
	if ok {
		return ONEPAIR
	}
	ok = highCardChecker(l)
	if ok {
		return HIGHCARD
	}
	return 0
}

func HandsTypeCreator(hands []string) []HandsType {
	ht := []HandsType{}
	for _, hand := range hands {
		t := typeOfHand(hand)
		ht = append(ht, HandsType{
			hand: hand,
			typ:  t,
		})
	}
	return ht
}

func partA() {
	hl, htb := readInput()
	m := map[rune]int{
		'2': 2,
		'3': 3,
		'4': 4,
		'5': 5,
		'6': 6,
		'7': 7,
		'8': 8,
		'9': 9,
		'T': 10,
		'J': 1,
		'Q': 12,
		'K': 13,
		'A': 14,
	}
	ht := HandsTypeCreator(hl)
	sort.Slice(ht, func(i, j int) bool {
		if ht[i].typ != ht[j].typ {
			return ht[i].typ < ht[j].typ
		} else {
			k := 0
			for ht[i].hand[k] == ht[j].hand[k] {
				k++
			}
			return m[rune(ht[i].hand[k])] < m[rune(ht[j].hand[k])]
		}
	})
	sum := 0
	for i, handsType := range ht {
		sum += (i + 1) * htb[handsType.hand]
	}
	fmt.Println(sum)
	//for _, handsType := range ht {
	//	fmt.Println(handsType)
	//}
}

func partB() {
	//m := map[rune]int{
	//	'2': 2,
	//	'3': 3,
	//	'4': 4,
	//	'5': 5,
	//	'6': 6,
	//	'7': 7,
	//	'8': 8,
	//	'9': 9,
	//	'T': 10,
	//	'J': 1,
	//	'Q': 12,
	//	'K': 13,
	//	'A': 14,
	//}
	//hl, htb := readInput()
}

func main() {
	partA()
}

type HandsType struct {
	hand string
	typ  int
}

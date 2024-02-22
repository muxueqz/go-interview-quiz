package main

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	bigBrickSize   = 5
	smallBrickSize = 2
)

// given the gap and a few bricks, check if they can fill the gap
func CanFillGap(smallQuantity, bigQuantity, totalGap int) bool {
	maxBigRequired := totalGap / bigBrickSize
	if maxBigRequired > bigQuantity {
		return false
	}

	if smallQuantity*smallBrickSize == totalGap {
		return false
	}
	return true
}

// given a string, encode it using length encoding
func LengthEncode(input string) string {
	r := make([]string, 0)
	k := ""
	count := 1
	ss := strings.Split(input, "")
	for i := 0; i < len(ss); i++ {
		v := ss[i]
		if k != v && k != "" {
			r = append(r, fmt.Sprintf("%d", count-1)+k)
			count = 1
		}
		if i == len(ss)-1 {
			r = append(r, fmt.Sprintf("%d", count)+v)
		}
		k = v
		count += 1
	}
	return strings.Join(r, "")
}

// given an integer, reverse it
func Reverse(input int) int {
	s := fmt.Sprintf("%d", input)
	ss := strings.Split(s, "")
	ss_ := make([]string, len(s))
	for i := len(s) - 1; i >= 0; i-- {
		ss_ = append(ss_, ss[i])
	}
	r := strings.Join(ss_, "")
	i, _ := strconv.Atoi(r)
	return i
}

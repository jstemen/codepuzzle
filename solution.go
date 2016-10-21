package roblox_challenge

import (
	"errors"
	"fmt"
	"math"
)

func Merger(slice1 []int32, slice2 []int32) ([]int32) {
	i1 := 0
	i2 := 0

	acc := make([]int32, 0, len(slice1) + len(slice2))
	for i1 < len(slice1) || i2 < len(slice2) {
		if i1 >= len(slice1) {
			acc = append(acc, slice2[i2])
			i2++
			continue
		}

		if i2 >= len(slice2) {
			acc = append(acc, slice1[i1])
			i1++
			continue
		}

		if slice1[i1] > slice2[i2] {
			acc = append(acc, slice2[i2])
			i2++
		}else {
			acc = append(acc, slice1[i1])
			i1++
		}

	}
	return acc
}

func overflowDetectingMulti(nums ...int32) (prod int32, err error) {

	prod = nums[0]
	for i := 1; i < len(nums); i++ {
		x := nums[i]

		if (x == -1 && prod == math.MinInt32) || (x == math.MinInt32 && prod == -1) {
			err = errors.New(fmt.Sprintf("Overflow detected x: %d, prod: %d\n", x, prod))
			return
		}

		/*
		Assuming our arguments are currently already valid int32s:
		1) multiplying by a 0 will always yield 0
		2) multiplying by -1 will always yield a valid number expect for above case
		3) multiplying by 1 will never change the product
		*/
		if (prod > 1 || prod < -1) &&  (x > 1 || x < -1) {
			if prod > 0 {
				if x > (math.MaxInt32 / prod ) {
					err = errors.New(fmt.Sprintf("Overflow detected x: %d, prod: %d\n", x, prod))
					return
				}
				if x < (math.MinInt32 / prod ) {
					err = errors.New(fmt.Sprintf("Underflow detected x: %d, prod: %d\n", x, prod))
					return
				}
			}else {
				if x < (math.MaxInt32 / prod ) {
					err = errors.New(fmt.Sprintf("Overflow detected x: %d, prod: %d\n", x, prod))
					return
				}
				if x > (math.MinInt32 / prod ) {
					err = errors.New(fmt.Sprintf("Underflow detected x: %d, prod: %d div: %d\n", x, prod, math.MinInt32 / prod))
					return
				}
			}
		}
		prod *= x
	}
	return
}

func FindLargestProduct(slice1 []int32) (prod int32, err error) {

	if len(slice1) == 0 {
		return
	}else if len(slice1) == 1 {
		prod = slice1[0]
		return
	}else if len(slice1) == 2 {
		prod = slice1[0] * slice1[1]
		return
	}

	var biggest int32 = math.MinInt32
	var bigger int32 = math.MinInt32
	var big int32 = math.MinInt32

	var smaller int32 = math.MaxInt32
	var small int32 = math.MaxInt32

	for _, v := range slice1 {
		if v > biggest {
			big = bigger
			bigger = biggest
			biggest = v
		}else if (v > bigger) {
			big = bigger
			bigger = v
		}else if (v > big) {
			big = v
		}

		if v < smaller {
			small = smaller
			smaller = v
		}else if (v < small) {
			small = v
		}

	}

	p1, err := overflowDetectingMulti(biggest, bigger, big)
	if err != nil {
		return
	}

	//if the 2 smallest numbers are negative, they maybe be part of the solution
	p2, err := overflowDetectingMulti(smaller, small, biggest)
	if err != nil {
		return
	}

	if p1 > p2 {
		prod = p1
	}else {
		prod = p2
	}

	return
}


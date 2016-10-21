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

	fmt.Printf("Multiplying #%v\n", nums)
	prod = nums[0]
	for i := 1; i < len(nums); i++ {
		x := nums[i]

		fmt.Printf("x:%d * prod:%d = %d\n", x, prod, x * prod)
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
				fmt.Println("prod is greater than 0 ")
				if x > (math.MaxInt32 / prod ) {
					err = errors.New(fmt.Sprintf("Overflow detected x: %d, prod: %d\n", x, prod))
					return
				}
				fmt.Printf("(math.MinInt32 / prod ) = %d \n", (math.MinInt32 / prod ))
				if x < (math.MinInt32 / prod ) {
					err = errors.New(fmt.Sprintf("Underflow detected x: %d, prod: %d\n", x, prod))
					return
				}
			}else {
				fmt.Println("prod is less than 0 ")
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

	var pos1 int32 = math.MinInt32
	var pos2 int32 = math.MinInt32
	var pos3 int32 = math.MinInt32

	var neg1 int32 = math.MaxInt32
	var neg2 int32 = math.MaxInt32

	for _, v := range slice1 {
		if v > pos1 {
			pos3 = pos2
			pos2 = pos1
			pos1 = v
		}else if (v > pos2) {
			pos3 = pos2
			pos2 = v
		}else if (v > pos3) {
			pos3 = v
		}

		if v < neg1 {
			neg2 = neg1
			neg1 = v
		}else if (v < neg2) {
			neg2 = v
		}

	}

	p1, err := overflowDetectingMulti(pos1, pos2, pos3)
	if err != nil {
		return
	}

	//Multiplying negative numbers first should prevent underflow
	p2, err := overflowDetectingMulti(neg1, neg2, pos1)
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


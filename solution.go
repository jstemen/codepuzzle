package roblox_challenge

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


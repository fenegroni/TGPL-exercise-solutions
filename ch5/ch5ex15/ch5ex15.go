package ch5ex15

func Min(v1 int, vn ...int) int {
	min := v1
	for _, v := range vn {
		if v < min {
			min = v
		}
	}
	return min
}

func Max(v1 int, vn ...int) int {
	max := v1
	for _, v := range vn {
		if v > max {
			max = v
		}
	}
	return max
}

package exercise5_15

func FastMin(vals ...int) int {
	min := vals[0]
	for _, v := range vals[1:] {
		if v < min {
			min = v
		}
	}
	return min
}

func Min(vals ...int) (int, bool) {
	if len(vals) == 0 {
		return 0, false
	}
	min := vals[0]
	for _, v := range vals[1:] {
		if v < min {
			min = v
		}
	}
	return min, true
}

func Max(vals ...int) int {
	return 0
}

package exercise5_15

func Min(vals ...int) int {
	min := vals[0]
	for _, v := range vals[1:] {
		if v < min {
			min = v
		}
	}
	return min
}

func Max(vals ...int) int {
	return 0
}

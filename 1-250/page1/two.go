package main

func two() int {
	i := 1
	j := 2
	sum := 0
	for i < 4000000 {
		if j % 2 == 0 {
			sum += j
		}
		k := j + i
		i = j
		j = k
	}
	return sum
}

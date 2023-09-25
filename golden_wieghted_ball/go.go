package main

import (
	"fmt"
	"math"
)

type Balls []int

func two_balls_approach(balls Balls) int {
	i := 0
	for {
		if len(balls)/2 < i {
			return 0
		}
		b1 := balls[i]
		b2 := balls[i+1]
		if b1 != b2 {
			min := math.Min(float64(b1), float64(b2))
			return int(min)
		}
		i += 2
	}
}

func binary_approach(balls Balls) int {
	if len(balls) == 0 {
		return 0
	}
	l := balls[:len(balls)/2]
	r := balls[len(balls)/2:]
	if len(balls) == 4 {
		// means we have reached a point where there are 2
		// balls on either side of the scale
		min_l := math.Min(float64(l[0]), float64(l[1]))
		min_r := math.Min(float64(r[0]), float64(r[1]))
		return int(math.Min(min_l, min_r))
	} else {
		if sum(&l) < sum(&r) {
			return binary_approach(l)
		} else {
			return binary_approach(r)
		}
	}
}

func sum(balls *Balls) int {
	res := 0
	for _, b := range *balls {
		res += b
	}
	return res
}

func main() {
	balls := []int{8, 8, 8, 8, 8, 7, 8, 8}
	var result int
	result = two_balls_approach(balls)
	fmt.Println(result)

	result = binary_approach(balls)
	fmt.Println(result)
}

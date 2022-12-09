package main

import (
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	X, Y int
}

func main() {
	f, _ := os.Open("input.txt")
	b, _ := io.ReadAll(f)
	s := string(b)

	rope := make([]*Point, 10)
	for i := 0; i < 10; i++ {
		rope[i] = &Point{0, 0}
	}
	tail := rope[9]
	head := rope[0]
	visited := make(map[Point]bool)

	visited[*tail] = true

	for _, line := range strings.Split(s, "\n") {
		if line == "" {
			continue
		}
		split := strings.Split(line, " ")
		times, _ := strconv.Atoi(split[1])
		movement := split[0]
		for i := 0; i < times; i++ {
			switch movement {
			case "U":
				head.X--
			case "D":
				head.X++
			case "L":
				head.Y--
			case "R":
				head.Y++
			}
			for i := 1; i < 10; i++ {
				prev := rope[i-1]
				act := rope[i]
				if math.Abs(float64(act.X)-float64(prev.X)) <= 1 && math.Abs(float64(act.Y)-float64(prev.Y)) <= 1 {
					break
				}
				if prev.Y < act.Y {
					act.Y--
				} else if prev.Y > act.Y {
					act.Y++
				}
				if prev.X < act.X {
					act.X--
				} else if prev.X > act.X {
					act.X++
				}
			}
			visited[*tail] = true
		}
	}
	r := 0
	for _, v := range visited {
		if v {
			r++
		}
	}
	println(r)
}

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
	tail := &Point{200, 0}
	head := &Point{200, 0}
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
			prev := &Point{head.X, head.Y}
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
			if math.Abs(float64(tail.X)-float64(head.X)) <= 1 && math.Abs(float64(tail.Y)-float64(head.Y)) <= 1 {
				continue
			}
			tail = prev
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

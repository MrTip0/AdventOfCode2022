package main

import (
	"container/list"
	"io"
	"math"
	"os"
	"strings"
)

type Point struct {
	x, y int
}

func main() {
	f, _ := os.Open("input.txt")
	b, _ := io.ReadAll(f)
	s := string(b)

	end := strings.Index(s, "E")
	start := strings.Index(s, "S")
	s = strings.Replace(s, "E", "z", 1)
	s = strings.Replace(s, "S", "`", 1)

	m := make([][]rune, 0)
	r := make([][]int, 0)

	for _, v := range strings.Split(s, "\n") {
		if v == "" {
			continue
		}
		m = append(m, []rune(v))
		vec := make([]int, len(v))
		for i := 0; i < len(v); i++ {
			vec[i] = math.MaxInt
		}
		r = append(r, vec)
	}

	line_length := len(m[0])
	line := start / (line_length + 1)
	column := start % (line_length + 1)

	r[line][column] = 0

	q := list.New()
	q.PushBack(&Point{line, column})

	for q.Len() > 0 {
		node := q.Front()
		val := node.Value.(*Point)
		q.Remove(node)
		act := m[val.x][val.y]
		act_res := r[val.x][val.y]
		if val.x > 0 {
			if m[val.x-1][val.y] <= act+1 && r[val.x-1][val.y] > act_res+1 {
				r[val.x-1][val.y] = act_res + 1
				q.PushBack(&Point{val.x - 1, val.y})
			}
		}
		if val.x < len(m)-1 {
			if m[val.x+1][val.y] <= act+1 && r[val.x+1][val.y] > act_res+1 {
				r[val.x+1][val.y] = act_res + 1
				q.PushBack(&Point{val.x + 1, val.y})
			}
		}
		if val.y > 0 {
			if m[val.x][val.y-1] <= act+1 && r[val.x][val.y-1] > act_res+1 {
				r[val.x][val.y-1] = act_res + 1
				q.PushBack(&Point{val.x, val.y - 1})
			}
		}
		if val.y < line_length-1 {
			if m[val.x][val.y+1] <= act+1 && r[val.x][val.y+1] > act_res+1 {
				r[val.x][val.y+1] = act_res + 1
				q.PushBack(&Point{val.x, val.y + 1})
			}
		}
	}

	line = end / (line_length + 1)
	column = end % (line_length + 1)

	println(r[line][column])
}

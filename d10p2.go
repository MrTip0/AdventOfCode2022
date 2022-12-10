package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	x := 1
	r := 0
	actions := make([]int, 0, 220)

	f, _ := os.Open("input.txt")
	b, _ := io.ReadAll(f)
	s := string(b)

	for _, v := range strings.Split(s, "\n") {
		if v == "" {
			continue
		}
		v := strings.Split(v, " ")
		actions = append(actions, 0)
		if v[0] == "addx" {
			val, _ := strconv.Atoi(v[1])
			actions = append(actions, val)
		}
	}
	for i := 1; i < 241; i++ {
		if i%40 >= x && i%40 <= x+2 {
			fmt.Print("#")
		} else {
			fmt.Print(".")
		}
		if i%40 == 0 {
			fmt.Println()
		}
		x += actions[i-1]
	}
	fmt.Println(r)
}

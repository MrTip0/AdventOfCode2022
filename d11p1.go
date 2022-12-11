package main

import (
	"container/list"
	"fmt"
	"io"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type Monkey struct {
	Items           *list.List
	Operation       func(val int) int
	Module          int
	IfTrue, IfFalse int
	Inspected       int
}

func main() {
	f, _ := os.Open("input.txt")
	b, _ := io.ReadAll(f)
	s := string(b)

	input := strings.Split(s, "\n\n")

	monkeys := make([]*Monkey, 0, 10)
	for _, m := range input {
		if r, e := regexp.MatchString(`^\s*$`, m); r || e != nil {
			continue
		}
		lines := strings.Split(m, "\n")
		monkey := &Monkey{}
		monkey.Items = GetItems(lines[1])
		monkey.Operation = CreateFunction(lines[2])
		monkey.Module = SplitAndRead(lines[3], "by ")
		monkey.IfTrue = SplitAndRead(lines[4], "monkey ")
		monkey.IfFalse = SplitAndRead(lines[5], "monkey ")
		monkeys = append(monkeys, monkey)
	}

	for i := 0; i < 20; i++ {
		DoRound(monkeys)
	}

	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].Inspected > monkeys[j].Inspected
	})

	fmt.Println(monkeys[0].Inspected * monkeys[1].Inspected)
}

func CreateFunction(description string) func(val int) int {
	description = strings.Split(description, "= ")[1]
	tokens := strings.Split(description, " ")
	var operator func(a, b int) int
	switch tokens[1] {
	case "+":
		operator = func(a, b int) int {
			return a + b
		}
	case "*":
		operator = func(a, b int) int {
			return a * b
		}
	}
	return func(val int) int {
		var a, b int
		if tokens[0] == "old" {
			a = val
		} else {
			a, _ = strconv.Atoi(tokens[0])
		}
		if tokens[2] == "old" {
			b = val
		} else {
			b, _ = strconv.Atoi(tokens[2])
		}
		return operator(a, b)
	}
}

func GetItems(data string) *list.List {
	data = strings.Split(data, ": ")[1]
	elements := strings.Split(data, ", ")
	r := list.New()
	for _, e := range elements {
		v, _ := strconv.Atoi(e)
		r.PushBack(v)
	}
	return r
}

func SplitAndRead(description, sep string) int {
	description = strings.Split(description, sep)[1]
	r, _ := strconv.Atoi(description)
	return r
}

func DoRound(V []*Monkey) {
	for _, monkey := range V {
		for {
			if monkey.Items.Len() == 0 {
				break
			}
			next := monkey.Items.Front()
			val := next.Value.(int)
			monkey.Items.Remove(next)
			val = monkey.Operation(val) / 3
			if val%monkey.Module == 0 {
				V[monkey.IfTrue].Items.PushBack(val)
			} else {
				V[monkey.IfFalse].Items.PushBack(val)
			}
			monkey.Inspected++
		}
	}
}

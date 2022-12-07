package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	value    int
	children map[string]*Node
	parent   *Node
}

func main() {
	root := &Node{
		value:    0,
		children: make(map[string]*Node, 0),
		parent:   nil,
	}
	ptr := root

	f, _ := os.Open("input.txt")
	b, _ := ioutil.ReadAll(f)

	for _, v := range strings.Split(string(b), "\n") {
		if v == "" {
			continue
		}
		if strings.HasPrefix(v, "$ cd ") {
			v = strings.Split(v, " ")[2]
			if v == "/" {
				ptr = root
			} else if v == ".." {
				ptr = ptr.parent
			} else if v != "." {
				ptr = ptr.children[v]
			}
		} else if v != "$ ls" {
			divided := strings.Split(v, " ")
			if divided[0] == "dir" {
				ptr.children[divided[1]] = &Node{
					value:    0,
					children: make(map[string]*Node, 0),
					parent:   ptr,
				}
			} else {
				val, _ := strconv.Atoi(divided[0])
				ptr.children[divided[1]] = &Node{
					value:    val,
					children: nil,
					parent:   ptr,
				}
			}
		}
	}

	vals := make([]int, 0)
	CountValues(&vals, root)
	var r int

	for _, v := range vals {
		if v < 100000 {
			r += v
		}
	}

	fmt.Printf("%d\n", r)
}

func CountValues(V *[]int, ptr *Node) int {
	r := 0
	for _, val := range ptr.children {
		if val.children == nil {
			r += val.value
		} else {
			r += CountValues(V, val)
		}
	}
	*V = append(*V, r)
	return r
}

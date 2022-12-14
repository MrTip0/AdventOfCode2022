package main

import (
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.Open("input.txt")
	b, _ := io.ReadAll(f)
	s := string(b)
	lowerY := math.MinInt
	var matrix [200][600]string
	x, y := -1, -1
	for _, line := range strings.Split(s, "\n") {
		if line == "" {
			continue
		}
		x, y = -1, -1
		couples := strings.Split(line, " -> ")
		c := strings.Split(couples[0], ",")
		x, _ = strconv.Atoi(c[0])
		y, _ = strconv.Atoi(c[1])
		if y > lowerY {
			lowerY = y
		}
		matrix[y][x] = "#"
		couples = couples[1:]
		for _, couple := range couples {
			c = strings.Split(couple, ",")
			newX, _ := strconv.Atoi(c[0])
			newY, _ := strconv.Atoi(c[1])

			xMov := int(math.Max(-1, math.Min(1, float64(newX-x))))
			for x != newX {
				x += xMov
				matrix[y][x] = "#"
			}
			yMov := int(math.Max(-1, math.Min(1, float64(newY-y))))
			for y != newY {
				y += yMov
				matrix[y][x] = "#"
			}
			if y > lowerY {
				lowerY = y
			}
		}
	}

	r := 0
	cont := true
	for cont {
		r++
		newSandY, newSandX := 0, 500
		for {
			if matrix[newSandY+1][newSandX] == "" {
				matrix[newSandY][newSandX] = ""
				newSandY++
				matrix[newSandY][newSandX] = "o"
				if newSandY > lowerY {
					cont = false
					break
				}
			} else if matrix[newSandY+1][newSandX-1] == "" {
				matrix[newSandY][newSandX] = ""
				newSandY++
				newSandX--
				matrix[newSandY][newSandX] = "o"
				if newSandY > lowerY {
					cont = false
					break
				}
			} else if matrix[newSandY+1][newSandX+1] == "" {
				matrix[newSandY][newSandX] = ""
				newSandY++
				newSandX++
				matrix[newSandY][newSandX] = "o"
				if newSandY > lowerY {
					cont = false
					break
				}
			} else {
				break
			}
		}
	}

	fmt.Println(r - 1)
}

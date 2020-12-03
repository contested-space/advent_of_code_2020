package main

import (
	"fmt"
	"strings"
	"os"
	"io/ioutil"
)

func main() {
	fileName := os.Args[1]
	data, _ := ioutil.ReadFile(fileName)
	lines := strings.Split(string(data), "\n")

	slopeRight := []int{1, 3, 5, 7, 1}
	slopeDown := []int{1, 1, 1, 1, 2}
	cycle := len(lines[0])

	nbTrees := []int{0, 0, 0, 0, 0}

	tree := "#"

	for slope := 0; slope < 5; slope++ {
		pos := 0
		for i := 0; i<len(lines) - 1; i += slopeDown[slope] {
			if lines[i][pos:pos+1] == tree {
				nbTrees[slope]++
			}
			pos = (pos + slopeRight[slope]) % cycle
		}
	}

	total := nbTrees[0] * nbTrees[1] * nbTrees[2] * nbTrees[3] * nbTrees[4]

	fmt.Printf("Multiple of numbers of trees: %d\n", total)

}

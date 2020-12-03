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

	pos := 0

	slopeRight := 3
	slopeDown := 1
	cycle := len(lines[0])

	nbTrees := 0

	tree := "#"

	for i := 0; i<len(lines) - 1; i += slopeDown {
		if lines[i][pos:pos+1] == tree {
			nbTrees++
		}
		pos = (pos + slopeRight) % cycle
	}

	fmt.Printf("Number of trees: %d\n", nbTrees)

}

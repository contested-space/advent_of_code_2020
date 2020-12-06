package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"strings"
)


func main() {
	fileName := os.Args[1]
	data, _ := ioutil.ReadFile(fileName)
	lines := strings.Split(string(data), "\n")

	var form = make(map[rune]int)
	count := 0
	nbPeople := 0

	for _, line := range(lines) {
		if line == "" {
			count += countYesses(form, nbPeople)
			resetDeclarationForm(form)
			nbPeople = 0
		} else {
			updateForm(form, line)
			nbPeople++
		}
	}
	fmt.Printf("Total sum of yesses: %d\n", count)
}

func updateForm(form map[rune]int, line string) {
	for _, char := range(line) {
		form[char]++
	}
}

func countYesses(form map[rune]int, nbPeople int) int {
	count := 0
	for _ , value := range(form) {
		if value == nbPeople {
			count++
		}
	}
	return count
}

func resetDeclarationForm(form map[rune]int) {
	for key, _ := range(form) {
		form[key] = 0
	}
}

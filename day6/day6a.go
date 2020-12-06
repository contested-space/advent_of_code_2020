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

	var form = make(map[rune]bool)
	count := 0

	for _, line := range(lines) {
		if line == "" {
			count += countYesses(form)
			resetDeclarationForm(form)
		} else {
			updateForm(form, line)
		}
	}
	fmt.Printf("Total sum of yesses: %d\n", count)
}

func updateForm(form map[rune]bool, line string) {
	for _, char := range(line) {
		form[char] = true
	}
}

func countYesses(form map[rune]bool) int {
	count := 0
	for _ , value := range(form) {
		if value == true {
			count++
		}
	}
	return count
}

func resetDeclarationForm(form map[rune]bool) {
	for key, _ := range(form) {
		form[key] = false
	}
}

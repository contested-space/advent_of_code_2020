package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	fileName := os.Args[1]
	data, _ := ioutil.ReadFile(fileName)
	numbersStrings := strings.Split(string(data), "\n")
	numbers := numberize(numbersStrings)

	for i, num1 := range numbers {
		for j, num2 := range numbers[i:] {
			for _, num3 := range numbers[j:] {
				if num1+num2+num3 == 2020 {
					fmt.Printf("%d\n", num1*num2*num3)
				}
			}
		}
	}
}

func numberize(in []string) []int {
	number_array := make([]int, len(in))
	for i, str := range in {
		num, _ := strconv.Atoi(str)
		number_array[i] = num
	}
	return number_array
}

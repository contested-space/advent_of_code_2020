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
	lines := strings.Split(string(data), "\n")

	nbValid := 0

	for _, line := range(lines) {
		passwordPair := strings.Split(line, ":")
		if len(passwordPair) < 2 {
			continue
		}
//		fmt.Printf("passwordPair : %s\n", passwordPair[1])
		rule := passwordPair[0]
		password := passwordPair[1]
		rulePair := strings.Split(rule, " ")
		ruleRange := rulePair[0]
		ruleChar := rulePair[1]
		rangePair := strings.Split(ruleRange, "-")
		pos1, _ := strconv.Atoi(rangePair[0])
		pos2, _ := strconv.Atoi(rangePair[1])

		//fmt.Printf("line : %s ", line)

		if valid(pos1, pos2, ruleChar, password) {
			nbValid++
			fmt.Printf("valid\n")
		} else {
			fmt.Printf("invalid\n")
		}

	}

	fmt.Printf("%d valid passwords\n", nbValid)


}

func valid(pos1 int, pos2 int, ruleChar string, password string) bool {
	fmt.Printf("ruleChar: %s, pos1: %d, pos1 value: %s, pos2: %d, pos2 value: %s ", ruleChar, pos1, password[pos1-1:pos1], pos2, password[pos2-1:pos2])
	if pos1 > len(password) && pos2 > len(password) {
		fmt.Printf("fail condition 1 ")
		return false
	}
  if password[pos1:pos1+1] == password[pos2:pos2+1] {
		fmt.Printf("fail condition 2 ")
		return false
	}
	if password[pos1:pos1+1] == ruleChar {
		return true
	}
	if password[pos2:pos2+1] == ruleChar {
		return true
	}
	fmt.Printf("fail everything ")
	return false
}

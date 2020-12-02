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

	for _, line := range lines {
		passwordPair := strings.Split(line, ":")
		if len(passwordPair) < 2 {
			continue
		}
		rule := passwordPair[0]
		password := passwordPair[1]
		rulePair := strings.Split(rule, " ")
		ruleRange := rulePair[0]
		ruleChar := rulePair[1]
		rangePair := strings.Split(ruleRange, "-")
		pos1, _ := strconv.Atoi(rangePair[0])
		pos2, _ := strconv.Atoi(rangePair[1])

		if valid(pos1, pos2, ruleChar, password) {
			nbValid++
		}
	}

	fmt.Printf("%d valid passwords\n", nbValid)

}

func valid(pos1 int, pos2 int, ruleChar string, password string) bool {
	if pos1 > len(password) && pos2 > len(password) {
		return false
	}
	// haha I kept the blank space so I don't need to adjust my indexes
	if password[pos1:pos1+1] == password[pos2:pos2+1] {
		return false
	}
	if password[pos1:pos1+1] == ruleChar {
		return true
	}
	if password[pos2:pos2+1] == ruleChar {
		return true
	}
	return false
}

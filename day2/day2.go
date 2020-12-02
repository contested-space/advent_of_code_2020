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
		fmt.Printf("rule : %s\n", rule)
		password := passwordPair[1]
		rulePair := strings.Split(rule, " ")
		ruleRange := rulePair[0]
		ruleChar := rulePair[1]
		rangePair := strings.Split(ruleRange, "-")
		rangeMin, _ := strconv.Atoi(rangePair[0])
		rangeMax, _ := strconv.Atoi(rangePair[1])

		if valid(rangeMin, rangeMax, ruleChar, password) {
			nbValid++
		}

	}

	fmt.Printf("%d valid passwords\n", nbValid)


}

func valid(rangeMin int, rangeMax int, ruleChar string, password string) bool {
	count := strings.Count(password, ruleChar)
	return count >= rangeMin && count <= rangeMax
}

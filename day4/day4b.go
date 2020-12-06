package main

import (
	"fmt"
	"strings"
	"os"
	"io/ioutil"
	"strconv"
)

type Passport struct {
    Byr, Iyr, Eyr, Hgt, Hcl, Ecl, Pid, Cid bool
}

func main() {
	fileName := os.Args[1]
	data, _ := ioutil.ReadFile(fileName)
	lines := strings.Split(string(data), "\n")

	passport := Passport{}
	resetPassport(&passport)

	nbValid := 0

	for _, line := range(lines) {
		if line == "" {
			if checkValidity(passport) {
				nbValid++
			}
			resetPassport(&passport)
			continue
		}
		paramStrings := strings.Split(line, " ")
		for _, paramString := range(paramStrings) {
			paramPair := strings.Split(paramString, ":")
			validateParam(paramPair[0], paramPair[1], &passport)
		}
	}

	fmt.Printf("Number of valid passports: %d\n", nbValid)
}

func validateParam(param string, value string, p *Passport) {
	switch param {
	case "byr":
		p.Byr = validateByr(value)
	case "iyr":
		p.Iyr = validateIyr(value)
	case "eyr":
		p.Eyr = validateEyr(value)
	case "hgt":
		p.Hgt = validateHgt(value)
	case "hcl":
		p.Hcl = validateHcl(value)
	case "ecl":
		p.Ecl = validateEcl(value)
	case "pid":
		p.Pid = validatePid(value)
	case "cid":
		p.Cid = true
	default:
	}
}

func validateByr(byr string) bool {
	numByr, _ := strconv.Atoi(byr)
	return len(byr) == 4 && numByr >= 1920 && numByr <= 2002
}

func validateIyr(iyr string) bool {
	numIyr, _ := strconv.Atoi(iyr)
	return len(iyr) == 4 && numIyr >= 2010 && numIyr <= 2020
}

func validateEyr(eyr string) bool {
	numEyr, _ := strconv.Atoi(eyr)
	return len(eyr) == 4 && numEyr >= 2020 && numEyr <= 2030
}

func validateHgt(hgt string) bool {
	if hgt[len(hgt)-2:] == "cm" {
		sizeCm, _ := strconv.Atoi(hgt[:len(hgt) - 2])
		return sizeCm >= 150 && sizeCm <= 193
	}
	if hgt[len(hgt)-2:] == "in" {
		sizeIn, _ := strconv.Atoi(hgt[:len(hgt) - 2])
		return sizeIn >= 59 && sizeIn <= 176
	}
	return false
}

func validateHcl(hcl string) bool {
	if hcl[0] != '#' {
		return false
	}
	for _, char := range(hcl[1:]) {
		if !validateHclChar(char) {
			return false
		}
	}
	return true
}

func validateHclChar(char rune) bool {
	return char >= '0' && char <= '9' || char >= 'a' && char <= 'f'
}

func checkValidity(p Passport) bool{
	return p.Byr && p.Iyr && p.Eyr && p.Hgt && p.Hcl && p.Ecl && p.Pid
}

func validateEcl(ecl string) bool {
	validColors := [7]string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}

	for _, col := range validColors {
		if col == ecl {
			return true
		}
	}
	return false
}

func validatePid(pid string) bool {
	if len(pid) != 9 {
		return false
	}
	for _, char := range(pid) {
		if char < '0' || char > '9' {
			return false
		}
	}
	return true
}

func resetPassport(p *Passport) {
	p.Byr = false
	p.Iyr = false
	p.Eyr = false
	p.Hgt = false
	p.Hcl = false
	p.Ecl = false
	p.Pid = false
	p.Cid = false
}

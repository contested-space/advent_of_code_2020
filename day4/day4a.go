package main

import (
	"fmt"
	"strings"
	"os"
	"io/ioutil"
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
			validateParam(paramPair[0], &passport)
		}
	}

	fmt.Printf("Number of valid passports: %d\n", nbValid)
}

func validateParam(param string, p *Passport) {
	switch param {
	case "byr":
		p.Byr = true
	case "iyr":
		p.Iyr = true
	case "eyr":
		p.Eyr = true
	case "hgt":
		p.Hgt = true
	case "hcl":
		p.Hcl = true
	case "ecl":
		p.Ecl = true
	case "pid":
		p.Pid = true
	case "cid":
		p.Cid = true
	default:
	}
}

func checkValidity(p Passport) bool {
	return p.Byr && p.Iyr && p.Eyr && p.Hgt && p.Hcl && p.Ecl && p.Pid
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

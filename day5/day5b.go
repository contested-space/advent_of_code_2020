package main

import (
	"fmt"
	"strings"
	"os"
	"io/ioutil"
	"math"
	"sort"
)


func main() {
	fileName := os.Args[1]
	data, _ := ioutil.ReadFile(fileName)
	lines := strings.Split(string(data), "\n")

	var wantedId int
	var seatIds []int

	for i, line := range(lines) {
		if len(line) != 10 || i >= len(lines) - 2 {
			continue
		}
		seatIds = append(seatIds, seatCodeToId(line))
	}

	sort.Ints(seatIds)

	for i:= 1; i < len(seatIds); i++ {
		if seatIds[i] > seatIds[i-1] + 1  {
			wantedId = seatIds[i] - 1
		}
	}
	fmt.Printf("max id: %d\n", wantedId)
}

func seatCodeToId(code string) int {
	rowNum := rowCodeToPosition(code[:7])
	colNum := columnCodeToPosition(code[7:])
	return 8 * rowNum + colNum


}

func rowCodeToPosition(code string) int {
	seatPos := 0
	for i, char := range(code) {
		seatPos += letterToVal(i, char)
	}
	return seatPos
}

func letterToVal(pos int, char rune) int {
	switch char {
	case 'F':
		return 0
	case 'B':
		return int(math.Exp2(float64(6-pos)))
	case 'L':
		return 0
	case 'R':
		return int(math.Exp2(float64(2-pos)))
	default:
		return 0
	}
}

func columnCodeToPosition(code string) int {
	seatPos := 0
	for i, char := range(code) {
		seatPos += letterToVal(i, char)
	}
	return seatPos
}

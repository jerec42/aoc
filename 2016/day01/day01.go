package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {

	body, err := os.Open("input01")
	if err != nil {
		log.Fatalf("unable to open file: %v", err)
	}

	scanner := bufio.NewScanner(body)

	scanner.Scan()
	line := scanner.Text()

	instructs := strings.Split(line, ", ")

	dir := [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	//represents direction - n,e,s,w

	type Location struct {
		X, Y int
	}
	locations := map[Location]bool{}

	pos := Location{0, 0}
	var part2 Location
	found := false

	locations[pos] = true
	// position on board (x,y)

	facing := 0
	// index of dir (0 = n, 1, = e , 2 = s, 3 = w)

	for _, v := range instructs {
		//fmt.Println(i, v, facing, dir[facing], pos)

		if string(v[0]) == "L" {
			if facing == 0 {
				facing = 3
			} else {
				facing -= 1
			}
		} else if string(v[0]) == "R" {
			if facing == 3 {
				facing = 0
			} else {
				facing += 1
			}
		}
		steps, _ := strconv.Atoi(v[1:])

		for step := 1; step <= steps; step++ {

			pos.X += dir[facing][0]
			pos.Y += dir[facing][1]

			if found == false && locations[pos] == true {
				part2 = pos
				found = true
				fmt.Println("FOUND")
			} else {
				locations[pos] = true
			}
		}

	}

	//part 1 answer:
	fmt.Println(math.Abs(float64(pos.X)) + math.Abs(float64(pos.Y)))

	fmt.Println(math.Abs(float64(part2.X)) + math.Abs(float64(part2.Y)))

}

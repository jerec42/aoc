package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	var score = 0
	var score2 = 0

	var resultmap = map[string]int{
		"win":  6,
		"draw": 3,
		"loss": 0,
	}

	var scoremap = map[string]int{
		"X": 1,
		"Y": 2,
		"Z": 3,
	}

	body, err := os.Open("input02")
	if err != nil {
		log.Fatalf("unable to open file: %v", err)
	}

	scanner := bufio.NewScanner(body)

	for scanner.Scan() {
		var line = scanner.Text()

		game := strings.Split(line, " ")
		//fmt.Println(game[0], " <> ", game[1])

		var r = ""
		var r2 = ""
		var pick = ""

		switch game[0] {
		case "A":
			switch game[1] {
			case "X":
				r = "draw"
			case "Y":
				r = "win"
			case "Z":
				r = "loss"
			}
		case "B":
			switch game[1] {
			case "X":
				r = "loss"
			case "Y":
				r = "draw"
			case "Z":
				r = "win"
			}
		case "C":
			switch game[1] {
			case "X":
				r = "win"
			case "Y":
				r = "loss"
			case "Z":
				r = "draw"
			}
		}

		switch game[1] {
		case "X":
			r2 = "loss"
			switch game[0] {
			case "A":
				pick = "Z"
			case "B":
				pick = "X"
			case "C":
				pick = "Y"
			}
		case "Y":
			r2 = "draw"
			switch game[0] {
			case "A":
				pick = "X"
			case "B":
				pick = "Y"
			case "C":
				pick = "Z"
			}
		case "Z":
			r2 = "win"
			switch game[0] {
			case "A":
				pick = "Y"
			case "B":
				pick = "Z"
			case "C":
				pick = "X"
			}
		}
		score += scoremap[game[1]] + resultmap[r]
		score2 += scoremap[pick] + resultmap[r2]

	}

	fmt.Println(score)
	fmt.Println(score2)

}

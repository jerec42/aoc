package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	//"strings"
)

func main() {

	body, err := os.Open("input06")
	if err != nil {
		log.Fatalf("unable to open file: %v", err)
	}

	scanner := bufio.NewScanner(body)

	scanner.Scan()
	line := scanner.Text()

	for i := 3; i < len(line); i += 1 {
		//fmt.Println(i, line[i], line[i-3:i+1])
		if dup_detector(line[i-3:i+1]) == false {
			fmt.Println(i + 1)
			break
		}
	}

	for i := 13; i < len(line); i += 1 {
		//fmt.Println(i, line[i], line[i-3:i+1])
		if dup_detector(line[i-13:i+1]) == false {
			fmt.Println(i + 1)
			break
		}
	}

	//fmt.Println(line)

}

func dup_detector(str string) bool {
	charmap := make(map[rune]bool, len(str))

	for _, c := range str {
		charmap[c] = true
	}
	//fmt.Println(charmap)
	if len(charmap) < len(str) {
		return true
	}
	return false
}

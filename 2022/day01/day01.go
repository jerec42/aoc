package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {

	var max = 0
	var sum = 0
	values := make([]int, 0)

	body, err := os.Open("input01")
	if err != nil {
		log.Fatalf("unable to open file: %v", err)
	}

	scanner := bufio.NewScanner(body)

	for scanner.Scan() {

		// fmt.Println(scanner.Text())
		var line = scanner.Text()

		if line == "" {
			if max < sum {
				max = sum
			}
			values = append(values, sum)
			sum = 0
		} else {
			cur, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}
			sum += cur
		}
		//println(max, sum)

	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	
	// Part 1
	println(max)

	//Part 2
	sort.Ints(values)
	l := len(values)
	result := 0

	for _, v := range values[l-3 : l] {
		//println(v)
		result += v
	}
	println(result)

}

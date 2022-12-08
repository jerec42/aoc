package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	var total_priority = 0
	var total_grouppriority = 0

	prioritymap := make(map[string]int)

	for i := 1; i <= 26; i++ {
		prioritymap[string(rune(i+96))] = i
	}
	for i := 27; i <= 52; i++ {
		prioritymap[string(rune(i+38))] = i
	}

	body, err := os.Open("input03")
	if err != nil {
		log.Fatalf("unable to open file: %v", err)
	}

	scanner := bufio.NewScanner(body)

	for scanner.Scan() {
		var line = scanner.Text()

		//fmt.Println(len(line))

		var priority = 0
		for _, b := range line[len(line)/2:] {
			for _, d := range line[:len(line)/2] {
				if d == b {
					//fmt.Println("Found item:", string(d), "with priority", prioritymap[string(d)], " at pos:", c)
					priority = prioritymap[string(d)]
				}
			}
			//fmt.Println(a, ":", b, ":", string(b), prioritymap[string(b)])
		}
		total_priority += priority

	}

	body2, err2 := os.Open("input03")
	if err2 != nil {
		log.Fatalf("unable to open file: %v", err2)
	}

	scanner2 := bufio.NewScanner(body2)

	for scanner2.Scan() {

		var pack1 = scanner2.Text()
		scanner2.Scan()
		var pack2 = scanner2.Text()
		scanner2.Scan()
		var pack3 = scanner2.Text()

		group_priority := getcommon(countchars(pack1, prioritymap), countchars(pack2, prioritymap), countchars(pack3, prioritymap))

		total_grouppriority += group_priority
	}

	fmt.Println(total_priority)
	fmt.Println(total_grouppriority)
}

func getcommon(a []int, b []int, c []int) int {

	for i := 1; i <= 52; i++ {

		if a[i] > 0 && b[i] > 0 && c[i] > 0 {
			return (i)
		}
	}
	return (0)
}

func countchars(str string, prioritymap map[string]int) []int {
	// return an array that counts each instance of a character , assigned to index by priority value
	// .  priority values are 1-52 (a-z then A-Z).  0 index will always have value 0

	//fmt.Println(str, prioritymap)

	counts := make([]int, 53)

	for _, c := range str {
		counts[prioritymap[string(c)]]++
	}

	fmt.Println(counts)
	return (counts)

}

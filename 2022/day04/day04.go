package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	fulloverlap_count := 0
	overlap_count := 0

	body, err := os.Open("input04")
	if err != nil {
		log.Fatalf("unable to open file: %v", err)
	}

	scanner := bufio.NewScanner(body)

	for scanner.Scan() {
		var line = scanner.Text()

		sections := strings.Split(line, ",")

		as := strings.Split(sections[0], "-")
		zs := strings.Split(sections[1], "-")

		// fmt.Println(as, zs)
		a, _ := strconv.Atoi(as[0])
		b, _ := strconv.Atoi(as[1])
		y, _ := strconv.Atoi(zs[0])
		z, _ := strconv.Atoi(zs[1])

		if (a <= y && b >= z) || (y <= a && z >= b) {
			fulloverlap_count++
			overlap_count++
		} else if ( a >= y && a <= z) || (b >= y && b <= z) {
			overlap_count++
		}


	}
	fmt.Println(fulloverlap_count)
	fmt.Println(overlap_count)

}

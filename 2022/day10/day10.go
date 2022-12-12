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

	// part 1
	answer, h := part1()
	fmt.Println(answer)

	//part 2
	answer2 := part2(h)
	fmt.Println(answer2)
}

func part2(cpu_cycle []int) string {

	output := ""
	//dimensions of output
	h := 6
	w := len(cpu_cycle) / h

	for x := 0; x < len(cpu_cycle); x++ {

		delta := (x % w) - cpu_cycle[x]
		if delta < 2 && delta > -2 {
			output += "#"
		} else {
			output += " "
		}
	}

	for r := 0; r < h; r++ {
		fmt.Println(output[r*w : (r+1)*w])
	}
	return identify_4x6(output)
}

func part1() (int, []int) {

	X := 1
	t := 1
	start_i := 20
	start_interval := 40

	retval := 0

	history := make([]int, 0)
	history = append(history, 1)

	body, err := os.Open("input10")
	if err != nil {
		log.Fatalf("unable to open file: %v", err)
	}

	scanner := bufio.NewScanner(body)

	for scanner.Scan() {

		line := strings.Split(scanner.Text(), " ")

		command := line[0]

		if command == "noop" {
			t++
			history = append(history, X)
			if (t-start_i)%start_interval == 0 {
				retval += t * X
			}
		} else if command == "addx" {
			t++
			history = append(history, X)
			if (t-start_i)%start_interval == 0 {
				retval += t * X
			}

			t++
			a, _ := strconv.Atoi(line[1])
			X += a
			history = append(history, X)
			if (t-start_i)%start_interval == 0 {
				retval += t * X
			}
		}
	}
	return retval, history
}

func identify_4x6(display string) string {

	h := 6
	w := 4
	length := len(display) / h
	count := length / (w + 1)
	message := make([]string, 0)

	for letter := 0; letter < count; letter++ {

		fmt.Println(display[letter*(w+1) : letter*(w+1)+w])
		switch s1 := display[letter*(w+1) : letter*(w+1)+w]; {
		case s1 == "####":
			switch s2 := display[1*length+(letter*(w+1)) : 1*length+(letter*(w+1)+w)]; {
			case s2 == "#   ":
				switch s6 := display[5*length+(letter*(w+1)) : 5*length+(letter*(w+1)+w)]; {
				case s6 == "#   ":
					message = append(message, "F")
				case s6 == "####":
					message = append(message, "E")
				}
			default:
				message = append(message, "Z")
			}

		default:
			message = append(message, "?")
		}
	}

	fmt.Println("message:", message)
	return strings.Join(message, "")

}

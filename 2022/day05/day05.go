package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
	"unsafe"
)

func main() {

	body, err := os.Open("input05")
	if err != nil {
		log.Fatalf("unable to open file: %v", err)
	}

	scanner := bufio.NewScanner(body)

	stacks := make([][]string, 9)
	for i := 0; i < 9; i++ {
		stacks[i] = make([]string, 0)
	}

	section1 := true
	for section1 {
		scanner.Scan()
		line := scanner.Text()

		if scanner.Text()[0:2] == " 1" {
			section1 = false
			//fmt.Println("exitting section1")
			scanner.Scan() // read null line separating sections
		} else {
			for i := 1; i < 34; i = i + 4 {
				id := int(i / 4)

				if len(line) > i {
					if line[i:i+1] != " " {
						stacks[id] = append(stacks[id], line[i:i+1])
					}

				}
			}
		}
	}
	//fmt.Println(stacks)
	//loadcounts(stacks)
	//fmt.Println("========================")

	for scanner.Scan() {
		var line = scanner.Text()

		//fmt.Println(line)

		vals := strings.Split(line, " ")
		quantity, _ := strconv.Atoi(vals[1])
		from_stack, _ := strconv.Atoi(vals[3])
		to_stack, _ := strconv.Atoi(vals[5])

		from_stack_i := from_stack - 1
		to_stack_i := to_stack - 1

		//fmt.Println(quantity, from_stack_i, to_stack_i)

		load := make([]string, quantity)
		result_from := make([]string, len(stacks[from_stack_i])-quantity)

		copy(load, stacks[from_stack_i][:quantity])
		copy(result_from, stacks[from_stack_i][quantity:])

		//fmt.Println(load, result_from, to_stack_i)

		// For part 2 - Comment out the following function call
		reverse_slice(load)

		//fmt.Println(load, result_from, to_stack_i, append(load, stacks[to_stack_i]...))

		stacks[to_stack_i] = append(load, stacks[to_stack_i]...)
		stacks[from_stack_i] = result_from

		//fmt.Println(stacks)
		//loadcounts(stacks)

	}
	//fmt.Println(stacks)

	output := ""
	for _, s := range stacks {
		output += s[0]
	}
	fmt.Println(output)
}

func reverse_slice(s []string) {
	// MODIFIES s IN PLACE (i dont understand why) (now i understand why, maybe)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func loadcounts(s [][]string) {
	// used for debugging only - count the containers of each type.

	res := make(map[string]int, 0)
	t := 0
	for _, stack := range s {
		fmt.Printf("%+v\n", (*reflect.SliceHeader)(unsafe.Pointer(&stack)))
		//fmt.Println("i:", i, stack, len(stack), cap(stack))
		for _, val := range stack {
			t += 1
			res[val] += 1
			/* if res[val] == nil {
				res[val] += 1
			} else {
				res[val] = 1
			} */
		}
	}
	fmt.Println(t, res)
}

package main

import (
	"bufio"
	"fmt"
	"log"
	"math/big"
	"os"
	"regexp"
	"strconv"
)

type Monkey struct {
	Items        []*big.Int
	Op           string
	Op_value     int
	Test_divisor int
	True_Monkey  int
	False_Monkey int
	Inspects     int
}

func main() {

	// part 1
	answer := part(1, 20)
	fmt.Println(answer)

	// part 2
	answer2 := part(2, 10000)
	fmt.Println(answer2)
}

func part(part, rounds int) int {

	monkeys := make([]Monkey, 0)

	body, err := os.Open("input11")
	if err != nil {
		log.Fatalf("unable to open file: %v", err)
	}

	scanner := bufio.NewScanner(body)

	for scanner.Scan() {

		scanner.Text()

		scanner.Scan()
		itemtext := regexp.MustCompile(",? ").Split(scanner.Text(), -1)

		items := make([]*big.Int, 0)
		for _, v := range itemtext[4:] {
			a, _ := strconv.Atoi(v)
			items = append(items, big.NewInt(int64(a)))
		}

		scanner.Scan()
		regex := *regexp.MustCompile(`new = old ([^ ]+) ([^ ]+)`)
		res := regex.FindStringSubmatch(scanner.Text())
		oper := ""
		oper_val := 0

		if res[2] == "old" {
			if res[1] == "*" {
				oper = "^"
				oper_val = 2
			} else if res[1] == "+" {
				oper = "*"
				oper_val = 2
			}
		} else {
			oper = res[1]
			oper_val, _ = strconv.Atoi(res[2])
		}

		scanner.Scan()
		regex = *regexp.MustCompile(`Test: divisible by ([^ ]+)`)
		res = regex.FindStringSubmatch(scanner.Text())
		div, _ := strconv.Atoi(res[1])

		scanner.Scan()
		regex = *regexp.MustCompile(`true: throw to monkey ([^ ]+)`)
		res = regex.FindStringSubmatch(scanner.Text())
		iftrue, _ := strconv.Atoi(res[1])

		scanner.Scan()
		regex = *regexp.MustCompile(`false: throw to monkey ([^ ]+)`)
		res = regex.FindStringSubmatch(scanner.Text())
		iffalse, _ := strconv.Atoi(res[1])

		scanner.Scan()

		monkeys = append(monkeys, Monkey{Items: items, Op: oper, Op_value: oper_val, Test_divisor: div, True_Monkey: iftrue, False_Monkey: iffalse})

	}

	// throw_forward_counter := 0

	test_multiple := monkeys[0].Test_divisor
	for _, v := range monkeys[1:] {
		test_multiple *= v.Test_divisor
	}

	for round := 0; round < rounds; round++ {
		/*if round%50 == 0 {
			fmt.Println(round)
		}*/

		for m_i, monkey := range monkeys {

			for _, item := range monkey.Items {

				//fmt.Println("inspecting:", item, "from:", monkey)

				monkey.Inspects += 1

				switch monkey.Op {
				case "+":
					item.Add(item, big.NewInt(int64(monkey.Op_value)))
				case "*":
					item.Mul(item, big.NewInt(int64(monkey.Op_value)))
				case "^":
					//item = int(math.Pow(float64(item), float64(monkey.Op_value)))
					item.Mul(item, item)
				}
				//fmt.Println("afterop:", item, "from:", monkey)

				if part == 1 {
					item.Div(item, big.NewInt(int64(3)))
					//fmt.Println("relief:", item, "from:", monkey)
				}

				throwto := m_i
				_, m := new(big.Int).DivMod(item, big.NewInt(int64(monkey.Test_divisor)), big.NewInt(0))
				//fmt.Println("mod:", m)
				if m.Int64() == 0 {
					//fmt.Println("TRUE")
					throwto = monkey.True_Monkey
				} else {
					//fmt.Println("FALSE")
					throwto = monkey.False_Monkey
				}

				item.Mod(item, big.NewInt(int64(test_multiple)))

				monkeys[throwto].Items = append(monkeys[throwto].Items, item)
				/*if throwto > m_i {
					throw_forward_counter++
				}*/

				//fmt.Println("thrown:", item, "from:", monkey, "to:", monkeys[throwto])

			}
			monkey.Items = nil
			monkeys[m_i] = monkey
		}

		/*if (round+1)%1000 == 0 {
			//fmt.Println("round:", round)
			for i, v := range monkeys {
				fmt.Println(i, " : ", v.Inspects)
				//fmt.Println(i, " : ", v)
			}
			//fmt.Println(throw_forward_counter)
		}*/
	}
	top2 := make([]int, 2)
	for _, v := range monkeys {

		if v.Inspects >= top2[0] {
			top2[1] = top2[0]
			top2[0] = v.Inspects
		} else if v.Inspects >= top2[1] {
			top2[1] = v.Inspects
		}
	}
	return top2[0] * top2[1]
}

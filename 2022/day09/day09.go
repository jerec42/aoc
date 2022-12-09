package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Coord struct {
	X, Y int
}

type Player struct {
	Position Coord
	Trail    []Coord
}

type Move struct {
	Step     []int
	Distance int
}

func main() {

	// part 1
	answer := congaline(2)
	fmt.Println(answer)

	//part 2
	answer2 := congaline(10)
	fmt.Println(answer2)
}

func congaline(playercount int) int {

	stepmap := map[string][]int{"U": {0, -1}, "D": {0, 1}, "R": {1, 0}, "L": {-1, 0}}

	conga := make([]Player, 0)

	for c := 0; c < playercount; c += 1 {
		conga = append(conga, Player{Position: Coord{X: 0, Y: 0}})
	}
	//fmt.Println(conga)

	body, err := os.Open("input09")
	if err != nil {
		log.Fatalf("unable to open file: %v", err)
	}

	scanner := bufio.NewScanner(body)

	for scanner.Scan() {

		line := scanner.Text()

		move := strings.Split(line, " ")
		dir := move[0]
		steps, _ := strconv.Atoi(move[1])

		//fmt.Println(dir, steps)

		for x := 0; x < steps; x += 1 {

			h_move := Move{stepmap[dir], 1}
			conga[0] = moveplayer(conga[0], h_move)

			for i := 1; i < len(conga); i += 1 {

				if is_adjacent_9way(conga[i-1], conga[i]) == false {
					t_move := follow(conga[i-1], conga[i])
					conga[i] = moveplayer(conga[i], t_move)
				}
			}

		}

		/* for i, p := range conga {
			fmt.Println("I:", i, "P:", p.Position)
		} */
	}

	// add ending position to list of "visited sites"
	for i, p := range conga {
		p.Trail = append(p.Trail, p.Position)
		// note here that original conga is not modified by the above statement!  p is not a pointer
		conga[i].Trail = p.Trail
	}
	//fmt.Println(conga)

	visitmap := make(map[Coord]bool)

	sites_visited := 0
	for _, v := range conga[len(conga)-1].Trail {
		if visitmap[v] == false {
			sites_visited += 1
			visitmap[v] = true
		}
	}
	//fmt.Println("part 1:", sites_visited)
	return (sites_visited)
}

func moveplayer(A Player, M Move) Player {
	A.Trail = append(A.Trail, A.Position)
	A.Position.X += M.Distance * M.Step[0]
	A.Position.Y += M.Distance * M.Step[1]
	return (A)
}

func is_adjacent_9way(A Player, B Player) bool {

	if A.Position == B.Position {
		return (true)
	}
	steps := [][]int{{0, -1}, {1, -1}, {1, 0}, {1, 1}, {0, 1}, {-1, 1}, {-1, 0}, {-1, -1}}

	for _, s := range steps {
		// fmt.Println(steps[s])
		if B.Position.X+s[0] == A.Position.X && B.Position.Y+s[1] == A.Position.Y {
			return (true)
		}
	}
	return (false)
}

func follow(A Player, B Player) Move {
	// fmt.Println("follow:", A.Position, B.Position)
	if A.Position.X > B.Position.X {
		if A.Position.Y > B.Position.Y {
			return (Move{[]int{1, 1}, 1})
		} else if A.Position.Y < B.Position.Y {
			return (Move{[]int{1, -1}, 1})
		} else {
			return (Move{[]int{1, 0}, 1})
		}
	} else if A.Position.X < B.Position.X {
		if A.Position.Y > B.Position.Y {
			return (Move{[]int{-1, 1}, 1})
		} else if A.Position.Y < B.Position.Y {
			return (Move{[]int{-1, -1}, 1})
		} else {
			return (Move{[]int{-1, 0}, 1})
		}
	} else {
		if A.Position.Y > B.Position.Y {
			return (Move{[]int{0, 1}, 1})
		} else if A.Position.Y < B.Position.Y {
			return (Move{[]int{0, -1}, 1})
		} else {
			return (Move{[]int{0, 0}, 1})
		}
	}
}

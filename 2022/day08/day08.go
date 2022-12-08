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

type Tree struct {
	height  int
	visible bool
}

func main() {

	body, err := os.Open("input08")
	if err != nil {
		log.Fatalf("unable to open file: %v", err)
	}

	scanner := bufio.NewScanner(body)

	grid := map[Coord]Tree{}
	row_i := 0
	bound_x := 0
	bound_y := 0
	max_scenic_score := 0

	for scanner.Scan() {

		line := scanner.Text()

		vals := strings.Split(line, "")
		bound_x = len(vals)

		for i, v := range vals {
			h, _ := strconv.Atoi(v)
			grid[Coord{X: i, Y: row_i}] = Tree{height: h, visible: true}
		}
		row_i += 1
	}
	bound_y = row_i

	//fmt.Println(grid)
	//fmt.Println(bound_x, bound_y)

	for y := 0; y < (bound_y); y += 1 {
		for x := 0; x < (bound_x); x += 1 {
			score := 0
			grid, score = vischeck(grid, x, y, bound_x, bound_y)

			if score > max_scenic_score {
				max_scenic_score = score
			}
		}
	}

	treecount := 0
	for y := 0; y < bound_y; y += 1 {
		for x := 0; x < bound_x; x += 1 {
			if grid[Coord{X: x, Y: y}].visible == true {
				treecount += 1
			}
		}
	}
	//fmt.Println(grid)
	fmt.Println(treecount)
	fmt.Println(max_scenic_score)
}

func vischeck(forest map[Coord]Tree, x int, y int, bound_x int, bound_y int) (map[Coord]Tree, int) {

	//fmt.Println("vischeck: ", x, y, bound_x, bound_y)

	my_Tree := forest[Coord{X: x, Y: y}]

	blocked_n, blocked_s, blocked_e, blocked_w := false, false, false, false
	count_n, count_s, count_e, count_w := 0, 0, 0, 0 // scenic score in each direction.  how many trees in that direction can be seen?

	// check north
	for n := y - 1; n >= 0; n -= 1 {
		count_n += 1
		if forest[Coord{X: x, Y: n}].height >= my_Tree.height {
			//fmt.Println("north blocker:", x, n, "blocks:", x, y)
			blocked_n = true
			break
		}
	}

	// check south
	for s := y + 1; s < bound_y; s += 1 {
		count_s += 1
		if forest[Coord{X: x, Y: s}].height >= my_Tree.height {
			//fmt.Println("south blocker:", x, s, "blocks:", x, y)
			blocked_s = true
			break
		}
	}

	// check east
	for e := x + 1; e < bound_x; e += 1 {
		count_e += 1
		if forest[Coord{X: e, Y: y}].height >= my_Tree.height {
			//fmt.Println("east blocker:", e, y, "blocks:", x, y)
			blocked_e = true
			break
		}
	}

	// check west
	for w := x - 1; w >= 0; w -= 1 {
		count_w += 1
		if forest[Coord{X: w, Y: y}].height >= my_Tree.height {
			//fmt.Println("west blocker:", w, y, "blocks:", x, y)
			blocked_w = true
			break
		}
	}

	if blocked_n && blocked_s && blocked_e && blocked_w {
		//fmt.Println(x, y, " is blocked")
		my_Tree.visible = false
		forest[Coord{X: x, Y: y}] = my_Tree
	}

	//fmt.Println(x, y, count_n, count_s, count_e, count_w)
	score := count_n * count_s * count_e * count_w

	return forest, score
}

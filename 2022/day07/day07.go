package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type File struct {
	Name string
	Size int
}

type Dir struct {
	Name     string
	Files    []File
	Subs     []*Dir
	CompSize int
}

func main() {

	body, err := os.Open("input07")
	if err != nil {
		log.Fatalf("unable to open file: %v", err)
	}

	scanner := bufio.NewScanner(body)

	/* example static declaration of values:
	var filer = File{Name: "a.txt", Size: 1777}
	var file1 = File{Name: "b.txt", Size: 14848514}
	var file2 = File{Name: "c.txt", Size: 486}

	var bdir = Dir{Name: "b_directory", Files: []File{file1}}
	var cdir = Dir{Name: "c_directory", Files: []File{file2}}
	var root = Dir{Name: "root", Files: []File{filer}, Subs: []*Dir{&bdir, &cdir}}
	*/

	cwd := make([]*Dir, 0)
	tree := make([]*Dir, 0)
	mode := "init"

	for scanner.Scan() {

		line := scanner.Text()

		words := strings.Split(line, " ")
		//fmt.Println(mode, words, root)

		if words[0] == "$" {
			mode = "prompt"
		} else {
			mode = "list"
		}

		if mode == "prompt" {
			if words[1] == "cd" {
				if words[2] == ".." {

					cwd = cwd[:len(cwd)-1]

				} else {
					var leaf = Dir{Name: words[2], CompSize: 0}

					if len(cwd) > 0 {
						cwd[len(cwd)-1].Subs = append(cwd[len(cwd)-1].Subs, &leaf)
					}
					cwd = append(cwd, &leaf)
					tree = append(tree, &leaf)
				}

				//fmt.Println(cwd)
			}
		}
		if mode == "list" {
			if words[0] != "dir" {
				s, _ := strconv.Atoi(words[0])
				cwd[len(cwd)-1].Files = append(cwd[len(cwd)-1].Files, File{Name: words[1], Size: s})
				for _, v := range cwd {
					v.CompSize += s
				}
			}

		}

	}

	//printtree(tree)
	fmt.Println(sumqualified(tree))

	fmt.Println(part2(tree))
}

func printtree(tree []*Dir) {
	for _, v := range tree {
		fmt.Println(v.Name, v.Files, v.Subs, v.CompSize)
	}
}

func sumqualified(tree []*Dir) int {
	sum := 0
	for _, v := range tree {
		if v.CompSize <= 100000 {
			sum += v.CompSize
		}
	}
	return (sum)
}

func part2(tree []*Dir) int {

	capacity := 70000000
	need := 30000000 - (capacity - tree[0].CompSize)

	min_delete := capacity

	for _, v := range tree {

		freeup := v.CompSize - need
		if freeup > 0 && v.CompSize < min_delete {
			min_delete = v.CompSize
		}
	}
	return (min_delete)
}

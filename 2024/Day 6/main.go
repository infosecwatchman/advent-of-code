package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	content, err := os.ReadFile("input.txt") // the file is inside the local directory
	if err != nil {
		fmt.Println("Err")
	}
	// ^ > < v
	lines := strings.Split(string(content), "\n")
	var grid [][]string
	var location []int

	for linenum, line := range lines {
		chars := strings.Split(line, "")
		grid = append(grid, chars)
		for charloc, char := range chars {
			if char != "#" && char != "." {
				location = []int{linenum, charloc}
			}
		}
	}
	for _, line := range grid {
		fmt.Println(line)
	}

	transformInput(grid, location, lines)
	var newset string
	for _, line := range grid {
		fmt.Println(line)
		newset = fmt.Sprintf("%s\n%s", newset, strings.Join(line, ""))
	}
	regex := regexp.MustCompile(`X`)
	result := len(regex.FindAllString(newset, -1)) + 1
	//fmt.Println(grid)

	fmt.Printf("Part 1: %d\n", result)

}

func transformInput(grid [][]string, location []int, lines []string) {
	for {
		if grid[location[0]][location[1]] == "^" {
			if location[0] == 0 {
				break
			} else {
				if grid[location[0]-1][location[1]] != "#" {
					grid[location[0]][location[1]] = "X"
					grid[location[0]-1][location[1]] = "^"
					location[0] = location[0] - 1
				} else {
					grid[location[0]][location[1]] = "X"
					grid[location[0]][location[1]+1] = ">"
					location[1] = location[1] + 1
				}
			}
		} else if grid[location[0]][location[1]] == ">" {
			if location[1] == len(grid[location[0]])-1 {
				break
			} else {
				if grid[location[0]][location[1]+1] != "#" {
					grid[location[0]][location[1]] = "X"
					grid[location[0]][location[1]+1] = ">"
					location[1] = location[1] + 1
				} else {
					grid[location[0]][location[1]] = "X"
					grid[location[0]+1][location[1]] = "v"
					location[0] = location[0] + 1
				}
			}
		} else if grid[location[0]][location[1]] == "v" {
			if location[0] == len(lines)-1 {
				break
			} else {
				if grid[location[0]+1][location[1]] != "#" {
					grid[location[0]][location[1]] = "X"
					grid[location[0]+1][location[1]] = "v"
					location[0] = location[0] + 1
				} else {
					grid[location[0]][location[1]] = "X"
					grid[location[0]][location[1]-1] = "<"
					location[1] = location[1] - 1
				}
			}
		} else if grid[location[0]][location[1]] == "<" {
			if location[1] == 0 {
				break
			} else {
				if grid[location[0]][location[1]-1] != "#" {
					grid[location[0]][location[1]] = "X"
					grid[location[0]][location[1]-1] = "<"
					location[1] = location[1] - 1
				} else {
					grid[location[0]][location[1]] = "X"
					grid[location[0]-1][location[1]] = "^"
					location[0] = location[0] - 1
				}
			}
		}
	}
}

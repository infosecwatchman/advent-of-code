package main

import (
	"fmt"
	"os"
	"regexp"
	"slices"
	"strings"
)

type hit struct {
	count int
	match string
}

func main() {
	content, err := os.ReadFile("input.txt") // the file is inside the local directory
	if err != nil {
		fmt.Println("Err")
	}
	// ^ > < v
	lines := strings.Split(string(content), "\n")
	var grid [][]string
	var location []int
	var start []int
	var part2grid [][]string
	for _, line := range lines {
		chars := strings.Split(line, "")
		part2grid = append(part2grid, chars)
	}
	for linenum, line := range lines {
		chars := strings.Split(line, "")
		grid = append(grid, chars)
		for charloc, char := range chars {
			if char != "#" && char != "." {
				location = []int{linenum, charloc}
				start = []int{linenum, charloc}
			}
		}
	}
	//for _, line := range grid {
	//	fmt.Println(line)
	//}
	transformInput(grid, location, lines)
	var newset string
	for _, line := range grid {
		//fmt.Println(line)
		newset = fmt.Sprintf("%s\n%s", newset, strings.Join(line, ""))
	}
	regex := regexp.MustCompile(`X`)
	result := len(regex.FindAllString(newset, -1))
	//fmt.Println(grid)

	fmt.Printf("Part 1: %d\n", result)
	//Part 2
	result = 0
	for rownum, row := range grid {
		for columnnum, column := range row {
			if column != "#" {
				if !(rownum == start[0] && columnnum == start[1]) {
					tempstart := slices.Clone(start)
					var tempdata [][]string
					for _, line := range part2grid {
						templine := slices.Clone(line)
						tempdata = append(tempdata, templine)
					}
					tempdata[rownum][columnnum] = "#"
					if checkLoop(tempdata, tempstart, lines) {
						result++
					}
					tempdata[rownum][columnnum] = "."
				}
			}
		}
	}
	fmt.Printf("Part 2: %d", result)

}

func transformInput(grid [][]string, location []int, lines []string) {
	for {
		if grid[location[0]][location[1]] == "^" {
			if location[0] == 0 {
				grid[location[0]][location[1]] = "X"
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
				grid[location[0]][location[1]] = "X"
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
				grid[location[0]][location[1]] = "X"
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
				grid[location[0]][location[1]] = "X"
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

func checkLoop(data [][]string, loc []int, reference []string) bool {
	//var hits []string
	var hits []hit
	for {
		if len(hits) == 0 {
			hits = append(hits, hit{match: fmt.Sprintf("%d,%d", loc[0], loc[1]), count: 1})
		}
		if data[loc[0]][loc[1]] == "^" {
			if loc[0] == 0 {
				break
			} else {
				if data[loc[0]-1][loc[1]] != "#" {
					data[loc[0]][loc[1]] = "X"
					data[loc[0]-1][loc[1]] = "^"
					loc[0] = loc[0] - 1
				} else {
					data[loc[0]][loc[1]] = "X"
					data[loc[0]][loc[1]+1] = ">"
					for localnum, local := range hits {
						if local.match == fmt.Sprintf("%d,%d", loc[0], loc[1]) {
							hits[localnum].count = local.count + 1
							break
						} else if localnum+1 == len(hits) {
							hits = append(hits, hit{match: fmt.Sprintf("%d,%d", loc[0], loc[1]), count: 1})
							break
						}
					}
					loc[1] = loc[1] + 1
				}
			}
		} else if data[loc[0]][loc[1]] == ">" {
			if loc[1] == len(data[loc[0]])-1 {
				break
			} else {
				if data[loc[0]][loc[1]+1] != "#" {
					data[loc[0]][loc[1]] = "X"
					data[loc[0]][loc[1]+1] = ">"
					loc[1] = loc[1] + 1
				} else {
					data[loc[0]][loc[1]] = "X"
					data[loc[0]+1][loc[1]] = "v"
					for localnum, local := range hits {
						if local.match == fmt.Sprintf("%d,%d", loc[0], loc[1]) {
							hits[localnum].count = local.count + 1
							break
						} else if localnum+1 == len(hits) {
							hits = append(hits, hit{match: fmt.Sprintf("%d,%d", loc[0], loc[1]), count: 1})
							break
						}
					}
					loc[0] = loc[0] + 1
				}
			}
		} else if data[loc[0]][loc[1]] == "v" {
			if loc[0] == len(reference)-1 {
				break
			} else {
				if data[loc[0]+1][loc[1]] != "#" {
					data[loc[0]][loc[1]] = "X"
					data[loc[0]+1][loc[1]] = "v"
					loc[0] = loc[0] + 1
				} else {
					data[loc[0]][loc[1]] = "X"
					data[loc[0]][loc[1]-1] = "<"
					for localnum, local := range hits {
						if local.match == fmt.Sprintf("%d,%d", loc[0], loc[1]) {
							hits[localnum].count = local.count + 1
							break
						} else if localnum+1 == len(hits) {
							hits = append(hits, hit{match: fmt.Sprintf("%d,%d", loc[0], loc[1]), count: 1})
							break
						}
					}
					loc[1] = loc[1] - 1
				}
			}
		} else if data[loc[0]][loc[1]] == "<" {
			if loc[1] == 0 {
				break
			} else {
				if data[loc[0]][loc[1]-1] != "#" {
					data[loc[0]][loc[1]] = "X"
					data[loc[0]][loc[1]-1] = "<"
					loc[1] = loc[1] - 1
				} else {
					data[loc[0]][loc[1]] = "X"
					data[loc[0]-1][loc[1]] = "^"
					for localnum, local := range hits {
						if local.match == fmt.Sprintf("%d,%d", loc[0], loc[1]) {
							hits[localnum].count = local.count + 1
							break
						} else if localnum+1 == len(hits) {
							hits = append(hits, hit{match: fmt.Sprintf("%d,%d", loc[0], loc[1]), count: 1})
							break
						}
					}
					loc[0] = loc[0] - 1
				}
			}
		}
		//fmt.Println(hits)
		//fmt.Println("\n\n")
		for _, local := range hits {
			if local.count > 2000 {
				return true
			}
		}
		//for _, line := range data {
		//	fmt.Println(line)
		//}
		//fmt.Println("\n")
	}
	return false
}

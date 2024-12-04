package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	content, err := os.ReadFile("input.txt") // the file is inside the local directory
	if err != nil {
		fmt.Println("Err")
	}
	var part1 int
	lines := strings.Split(string(content), "\n")
	lines = slices.Delete(lines, len(lines)-1, len(lines))
	for linenum, line := range lines {
		chars := strings.Split(line, "")
		for charsel, char := range chars {
			//horizontal forward
			if charsel <= len(line)-4 {
				if char == "X" && chars[charsel+1] == "M" && chars[charsel+2] == "A" && chars[charsel+3] == "S" {
					part1++
				}
			}
			//horizontal backward
			if charsel >= 3 {
				if char == "X" && chars[charsel-1] == "M" && chars[charsel-2] == "A" && chars[charsel-3] == "S" {
					part1++
				}
			}

			if linenum >= 3 {
				//vertical upward
				if char == "X" && strings.Split(lines[linenum-1], "")[charsel] == "M" && strings.Split(lines[linenum-2], "")[charsel] == "A" && strings.Split(lines[linenum-3], "")[charsel] == "S" {
					part1++
				}
				//diagonal left upward
				if charsel >= 3 {
					if char == "X" && strings.Split(lines[linenum-1], "")[charsel-1] == "M" && strings.Split(lines[linenum-2], "")[charsel-2] == "A" && strings.Split(lines[linenum-3], "")[charsel-3] == "S" {
						part1++
					}
				}
				//diagonal right upward
				if charsel <= len(line)-4 {
					if char == "X" && strings.Split(lines[linenum-1], "")[charsel+1] == "M" && strings.Split(lines[linenum-2], "")[charsel+2] == "A" && strings.Split(lines[linenum-3], "")[charsel+3] == "S" {
						part1++
					}
				}
			}

			if linenum <= len(lines)-4 {
				//vertical downward
				if char == "X" && strings.Split(lines[linenum+1], "")[charsel] == "M" && strings.Split(lines[linenum+2], "")[charsel] == "A" && strings.Split(lines[linenum+3], "")[charsel] == "S" {
					part1++
				}
				//diagonal left downward
				if charsel >= 3 {
					if char == "X" && strings.Split(lines[linenum+1], "")[charsel-1] == "M" && strings.Split(lines[linenum+2], "")[charsel-2] == "A" && strings.Split(lines[linenum+3], "")[charsel-3] == "S" {
						part1++
					}
				}
				//diagonal right downward
				if charsel <= len(line)-4 {
					//fmt.Printf("%d x %d: %s\n", linenum, charsel, line)
					if char == "X" && strings.Split(lines[linenum+1], "")[charsel+1] == "M" && strings.Split(lines[linenum+2], "")[charsel+2] == "A" && strings.Split(lines[linenum+3], "")[charsel+3] == "S" {
						part1++
					}
				}
			}

		}
	}
	fmt.Printf("Part 1: %d\n", part1)

	var part2 int
	for linenum, line := range lines {
		if linenum > 0 && linenum < len(lines)-1 {
			chars := strings.Split(line, "")
			for charsel, char := range chars {
				if charsel > 0 && charsel < len(line)-1 && char == "A" {
					var match1, match2 bool = false, false
					upleft := strings.Split(lines[linenum-1], "")[charsel-1]
					upright := strings.Split(lines[linenum-1], "")[charsel+1]
					downleft := strings.Split(lines[linenum+1], "")[charsel-1]
					downright := strings.Split(lines[linenum+1], "")[charsel+1]
					if upleft == "M" && downright == "S" {
						match1 = true
					}
					if upleft == "S" && downright == "M" {
						match1 = true
					}
					if downleft == "M" && upright == "S" {
						match2 = true
					}
					if downleft == "S" && upright == "M" {
						match2 = true
					}
					if match1 && match2 {
						part2++
					}
				}
			}
		}
	}
	fmt.Printf("Part 2: %d\n", part2)

}

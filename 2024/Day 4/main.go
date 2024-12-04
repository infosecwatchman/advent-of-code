package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	content, err := os.ReadFile("input.txt") // the file is inside the local directory
	if err != nil {
		fmt.Println("Err")
	}
	var result int
	lines := strings.Split(string(content), "\n")
	for linenum, line := range lines {
		chars := strings.Split(line, "")
		for charsel, char := range chars {
			//horizontal forward
			if charsel <= len(line)-4 {
				if char == "X" && chars[charsel+1] == "M" && chars[charsel+2] == "A" && chars[charsel+3] == "S" {
					result++
				}
			}
			//horizontal backward
			if charsel >= 3 {
				if char == "X" && chars[charsel-1] == "M" && chars[charsel-2] == "A" && chars[charsel-3] == "S" {
					result++
				}
			}

			if linenum >= 3 {
				//vertical upward
				if char == "X" && strings.Split(lines[linenum-1], "")[charsel] == "M" && strings.Split(lines[linenum-2], "")[charsel] == "A" && strings.Split(lines[linenum-3], "")[charsel] == "S" {
					result++
				}
				//diagonal left upward
				if charsel >= 3 {
					if char == "X" && strings.Split(lines[linenum-1], "")[charsel-1] == "M" && strings.Split(lines[linenum-2], "")[charsel-2] == "A" && strings.Split(lines[linenum-3], "")[charsel-3] == "S" {
						result++
					}
				}
				//diagonal right upward
				if charsel <= len(line)-4 {
					if char == "X" && strings.Split(lines[linenum-1], "")[charsel+1] == "M" && strings.Split(lines[linenum-2], "")[charsel+2] == "A" && strings.Split(lines[linenum-3], "")[charsel+3] == "S" {
						result++
					}
				}
			}

			if linenum <= len(lines)-4 {
				//vertical downward
				if char == "X" && strings.Split(lines[linenum+1], "")[charsel] == "M" && strings.Split(lines[linenum+2], "")[charsel] == "A" && strings.Split(lines[linenum+3], "")[charsel] == "S" {
					result++
				}
				//diagonal left downward
				if charsel >= 3 {
					if char == "X" && strings.Split(lines[linenum+1], "")[charsel-1] == "M" && strings.Split(lines[linenum+2], "")[charsel-2] == "A" && strings.Split(lines[linenum+3], "")[charsel-3] == "S" {
						result++
					}
				}
				//diagonal right downward
				if charsel <= len(line)-4 {
					//fmt.Printf("%d x %d: %s\n", linenum, charsel, line)
					if char == "X" && strings.Split(lines[linenum+1], "")[charsel+1] == "M" && strings.Split(lines[linenum+2], "")[charsel+2] == "A" && strings.Split(lines[linenum+3], "")[charsel+3] == "S" {
						result++
					}
				}
			}

		}
	}
	fmt.Println(result)

}

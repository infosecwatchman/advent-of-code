package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	content, err := os.ReadFile("input.txt") // the file is inside the local directory
	if err != nil {
		fmt.Println("Err")
	}
	part1 := GetNumMatches(string(content))
	fmt.Printf("Part 1: %d\n", part1)

	part2regex := regexp.MustCompile(`(?:^|do\(\))(.*?)(?:don't\(\)|$)`)
	var part2string string
	matches := part2regex.FindAllString(strings.Join(strings.Split(string(content), "\n"), ""), -1)
	for _, match := range matches {
		part2string = part2string + match
	}
	part2 := GetNumMatches(part2string)
	fmt.Printf("Part 2: %d\n", part2)
}

func GetNumMatches(content string) int {
	var result int
	regex, err := regexp.Compile(`(mul\([0-9]+,[0-9]+\))`)
	if err != nil {
		fmt.Println(err)
	}
	matches := regex.FindAllString(content, -1)
	for _, submatch := range matches {
		//fmt.Println(submatch)
		numregex, err := regexp.Compile(`([0-9]+)`)
		if err != nil {
			fmt.Println(err)
		}
		nummatches := numregex.FindAllString(submatch, -1)
		num1, _ := strconv.Atoi(nummatches[0])
		num2, _ := strconv.Atoi(nummatches[1])
		result = result + (num1 * num2)
	}
	return result
}

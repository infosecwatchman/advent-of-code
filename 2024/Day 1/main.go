package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"slices"
	"strconv"
	"strings"
)

func main() {
	content, err := ioutil.ReadFile("input.txt") // the file is inside the local directory
	if err != nil {
		fmt.Println("Err")
	}
	var pairs1, pairs2 []int
	var result int
	for _, line := range strings.Split(string(content), "\n") {
		//fmt.Println(line)
		pairs := strings.SplitN(line, "   ", 2)
		if len(pairs) == 2 {
			num0, _ := strconv.Atoi(pairs[0])
			num1, _ := strconv.Atoi(pairs[1])
			pairs1 = append(pairs1, num0)
			pairs2 = append(pairs2, num1)
		}
	}
	slices.Sort(pairs1)
	slices.Sort(pairs2)
	for index, pair := range pairs1 {
		result = result + int(math.Abs(float64(pair-pairs2[index])))
		//fmt.Printf("%d - %d = %d\n", pair, pairs2[index], int(math.Abs(float64(pair-pairs2[index]))))
	}
	fmt.Printf("Part 1: %d\n", result)
	result = 0
	for _, pair1 := range pairs1 {
		count := 0
		for _, pair2 := range pairs2 {
			if pair1 == pair2 {
				count = count + 1
			}
		}
		result = result + (pair1 * count)
	}
	fmt.Printf("Part 2: %d", result)
	//fmt.Println(pairs1)
	//fmt.Println(string(content)) // This is some content
}

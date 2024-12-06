package main

import (
	"fmt"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func main() {
	content, err := os.ReadFile("input.txt") // the file is inside the local directory
	if err != nil {
		fmt.Println("Err")
	}
	lines := strings.Split(string(content), "\n")
	var pages, ruleset, badrules []string
	var goodrules [][]string
	var result int
	for _, line := range lines {
		if strings.Contains(line, "|") {
			pages = append(pages, line)
		}
		if strings.Contains(line, ",") {
			ruleset = append(ruleset, line)
		}
	}
	for _, ruleline := range ruleset {
		rules := strings.Split(ruleline, ",")
		var valid bool = false
		for _, page := range pages {
			firstPage := strings.Split(page, "|")[0]
			secondPage := strings.Split(page, "|")[1]
			if strings.Contains(ruleline, firstPage) && strings.Contains(ruleline, secondPage) {
				if slices.Index(rules, secondPage) > slices.Index(rules, firstPage) {
					valid = true
					goodrules = append(goodrules, rules)
				} else {
					badrules = append(badrules, ruleline)
					valid = false
					break
				}
			}
		}
		if valid {
			ruleint, _ := strconv.Atoi(rules[(len(rules) / 2)])
			result = result + ruleint
		}
	}
	fmt.Printf("Part 1: %d\n", result)
	result = 0

	for _, badrule := range badrules {
		set := strings.Split(badrule, ",")
		var compariter [][]string
		for _, items1 := range set {
			compariter = append(compariter, []string{items1, "0"})
		}
		for itemsint1, items1 := range set {
			for itemsint2, items2 := range set {
				if itemsint1 != itemsint2 {
					// 1/2
					pagelocal := slices.Index(pages, fmt.Sprintf("%s|%s", items1, items2))
					if pagelocal != -1 {
						converted, _ := strconv.Atoi(compariter[itemsint2][1])
						converted++
						compariter[itemsint2][1] = strconv.Itoa(converted)
					}
					// 2/1
					pagelocal = slices.Index(pages, fmt.Sprintf("%s|%s", items2, items1))
					if pagelocal != -1 {
						converted, _ := strconv.Atoi(compariter[itemsint1][1])
						converted++
						compariter[itemsint1][1] = strconv.Itoa(converted)
					}
				}
			}
		}
		sort.SliceStable(compariter, func(i, j int) bool {
			num1, _ := strconv.Atoi(compariter[i][1])
			num2, _ := strconv.Atoi(compariter[j][1])
			return num1 < num2
		})
		ruleint, _ := strconv.Atoi(compariter[(len(compariter) / 2)][0])
		result = result + ruleint
	}
	fmt.Printf("Part 2: %d\n", result)

}

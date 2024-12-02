package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	content, err := os.ReadFile("input.txt") // the file is inside the local directory
	if err != nil {
		fmt.Println("Err")
	}
	var safereports int
	var part2safereports int
	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		chars := strings.Split(line, " ")
		safe := true
		safe = CharEval(chars, safe)
		if safe {
			//fmt.Fprintf(os.Stdout, "%s SAFE \033[0m %d:\t%s\t %d \n", "\033[0;32m", linenum, line, len(chars))
			safereports++
		} else {
			//fmt.Fprintf(os.Stdout, "%s UNSAFE \033[0m %d:\t%s\t %d \n", "\033[0;31m", linenum, line, len(chars))
		}
	}
	fmt.Printf("Part 1: %d\n", safereports)

	for _, line := range lines {
		chars := strings.Split(line, " ")
		safe := true
		safe = CharEval(chars, safe)
		if safe {
			//fmt.Fprintf(os.Stdout, "%s SAFE \033[0m %d:\t%s\t %d \n", "\033[0;32m", linenum, line, len(chars))
			part2safereports++
		} else {
			safe = true
			for charnum, _ := range chars {
				newchars := slices.Clone(chars)
				newchars = slices.Delete(newchars, charnum, charnum+1)
				//fmt.Printf("%d: Original Chars: %s\n", linenum+1, chars)
				//fmt.Printf("%d: new Chars: %s\n", linenum+1, newchars)
				if CharEval(newchars, safe) {
					part2safereports++
					break
				}

			}
			//fmt.Fprintf(os.Stdout, "%s UNSAFE \033[0m %d:\t%s\t %d \n", "\033[0;31m", linenum, line, len(chars))
		}
	}
	fmt.Printf("Part 2: %d\n", part2safereports)

}

func CharEval(chars []string, safe bool) bool {
	var increasing bool = true
	for charnum, char := range chars {
		currentchar, _ := strconv.Atoi(char)
		var nextchar int
		if charnum == len(chars)-1 {
			safe = true
			break
		} else {
			nextchar, _ = strconv.Atoi(chars[charnum+1])
			if currentchar == nextchar {
				//fmt.Printf("%d: Fail on Current char %d, no difference in next char.\n", linenum, currentchar)
				safe = false
				break
			}
		}
		if charnum == 0 {
			nextmatch := currentchar - nextchar
			if int(math.Abs(float64(nextmatch))) > 3 {
				//fmt.Printf("%d: Fail on Current char %d, difference in next char is greater than 3.\n", linenum, currentchar)
				safe = false
				break
			} else {
				if nextmatch < 0 {
					increasing = true
				} else {
					increasing = false
				}
			}
		} else {
			nextmatch := currentchar - nextchar
			if int(math.Abs(float64(nextmatch))) > 3 {
				//fmt.Printf("%d: Fail on Current char %d, difference in next char is greater than 3.\n", linenum, currentchar)
				safe = false
				break
			} else {
				if int(math.Abs(float64(nextmatch))) <= 3 {
					if nextmatch < 0 && increasing {
						increasing = true
					} else if nextmatch > 0 && increasing == false {
						increasing = false
					} else {
						//fmt.Printf("%d: Fail on Current char %d next char %d, alternating direction.\n", linenum, currentchar, nextchar)
						safe = false
						break
					}
				}
			}
		}
	}
	return safe
}

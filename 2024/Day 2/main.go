package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	content, err := os.ReadFile("input.txt") // the file is inside the local directory
	if err != nil {
		fmt.Println("Err")
	}
	var safereports int
	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		chars := strings.Split(line, " ")
		safe := true
		for charnum, char := range chars {
			currentchar, _ := strconv.Atoi(char)
			previouschar, _ := strconv.Atoi(chars[charnum-1])
			nextchar, _ := strconv.Atoi(chars[charnum+1])
			var increasing bool
			if currentchar == 0 {
				nextmatch := currentchar - nextchar
				if int(math.Abs(float64(nextmatch))) > 3 || int(math.Abs(float64(nextmatch))) == 0 {
					safe = false
					break
				}

			} else if currentchar == len(chars)-1 {
				nextmatch := currentchar - nextchar
				previousmatch := currentchar - previouschar

			} else {

			}
		}
	}

}

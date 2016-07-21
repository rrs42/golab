package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/emirpasic/gods/sets"
	"github.com/emirpasic/gods/stacks/arraystack"
)

func parseSaveFile(fname string, keywords *sets.Set) {
	s := arraystack.New()

	f, err := os.Open(fname)
	if err != nil {
		panic(err.Error())
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	isParameter := func(line string) bool {
		match, err := regexp.MatchString("\\w+ = \\w*", line)
		if err != nil {
			panic(err.Error())
		}

		return match
	}

	var prevLine string
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "{" && keywords.Contains(prevLine) {
			s.Push(prevLine)
		} else if isParameter(line) {
			parts := strings.Split(line, " = ")
			fmt.Println(parts)
		} else {
			prevLine = line
		}
	}

	fmt.Println("\nLast Token")
	fmt.Print(s)
}

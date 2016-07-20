// Kerbal Knowledge Online
//
// Parse a KSP save file, extract the roster and determine the best way to
//    advance each one to level 3
//

package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"

	set "github.com/emirpasic/gods/sets/hashset"
	stack "github.com/emirpasic/gods/stacks/linkedliststack"
)

type vessel struct {
	name       string
	vesselType string
}

func main() {
	keySections := set.New()

	keySections.Add("GAME")
	keySections.Add("VESSEL")
	keySections.Add("ROSTER")

	parseSaveFile2("persistent.sfs", keySections)
}

func parseSaveFile2(fname string, keywords *set.Set) {
	s := stack.New()

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

func parseSaveFile(fname string, keywords *set.Set) {
	s := stack.New()

	f, err := os.Open(fname)
	if err != nil {
		panic(err.Error())
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)

	var prevToken string
	for scanner.Scan() {
		token := scanner.Text()
		if token == "{" && keywords.Contains(prevToken) {
			s.Push(prevToken)
		} else if token == "}" {
			frame, _ := s.Pop()
			if frame != nil {
				fmt.Println(frame)
			}
		} else {
			prevToken = token
		}
	}

	fmt.Println("\nLast Token")
	fmt.Print(s)
}

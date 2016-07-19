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

	"github.com/emirpasic/gods/sets/hashset"
	"github.com/emirpasic/gods/stacks/linkedliststack"
)

func main() {
	keySections := hashset.New()

	keySections.Add("GAME")
	keySections.Add("VESSEL")
	keySections.Add("ROSTER")

	openSaveFile("persistent.sfs", keySections)
}

func openSaveFile(fname string, keywords *hashset.Set) {
	s := linkedliststack.New()

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
			fmt.Println(frame)
		} else {
			prevToken = token
		}
	}

	fmt.Println("\nLast Token")
	fmt.Print(s)
}

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
)

func main() {
	openSaveFile("persistent.sfs")
}

func openSaveFile(fname string) {
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
		if token == "{" {
			fmt.Print(prevToken)
		} else {
			prevToken = token
		}
	}
}

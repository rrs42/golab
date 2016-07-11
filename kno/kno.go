// Kerbal Knowledge Online
//
// Parse a KSP save file, extract the roster and determine the best way to
//    advance each one to level 3
//

package main

import "os"

func main() {

}

func openSaveFile(fname string) {
	f, err := os.Open(fname)
	if err != nil {
		panic(err.Error())
	}

	_ = f
}

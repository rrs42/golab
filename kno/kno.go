// Kerbal Knowledge Online
//
// Parse a KSP save file, extract the roster and determine the best way to
//    advance each one to level 3
//

package main

import set "github.com/emirpasic/gods/sets/hashset"

func main() {

	keySections := set.New()

	keySections.Add("GAME")
	keySections.Add("VESSEL")
	keySections.Add("ROSTER")

	parseSaveFile("persistent.sfs", keySections)
}

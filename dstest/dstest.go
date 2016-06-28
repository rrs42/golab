package main

import (
	"fmt"

	"github.com/emirpasic/gods/lists/arraylist"
	"github.com/emirpasic/gods/trees/redblacktree"
)

func main() {
	list := arraylist.New()
	list.Add(2, 1, 3)
	fmt.Println(list.Values())

	it := list.Iterator()
	for it.Next() {
		key, value := it.Index(), it.Value()
		fmt.Println(key, value)
	}

	tree := redblacktree.NewWithIntComparator()

	tree.Put(1, "x")
	tree.Put(2, "a")
	tree.Put(3, "h")
	tree.Put(4, "e")

	fmt.Println(tree)
}

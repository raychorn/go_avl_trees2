package main

// https://github.com/emirpasic/gods/blob/master/trees/avltree/avltree_test.go

import (
	"fmt"
	"strings"
	avltree "github.com/emirpasic/gods/trees/avltree"
)

func main() {
	tree := avltree.NewWithIntComparator()
	tree.Put(13, 5)
	tree.Put(8, 3)
	tree.Put(17, 7)
	tree.Put(1, 1)
	tree.Put(11, 4)
	tree.Put(15, 6)
	tree.Put(25, 9)
	tree.Put(6, 2)
	tree.Put(22, 8)
	tree.Put(27, 10)

	it := tree.Iterator()
	count := 0
	for it.Next() {
		count++
		key := it.Key()
		value := it.Value()

		fmt.Printf("%d -> %d\n", key, value)
	}
	fmt.Println(strings.Repeat("-", 30))

	it = tree.Iterator()
	it.Begin()
	count = 0
	for it.Next() {
		count++
		key := it.Key()
		value := it.Value()

		fmt.Printf("%d -> %d\n", key, value)
	}
	fmt.Println(strings.Repeat("-", 30))
	fmt.Println(tree.String())
	fmt.Println("Done !!!")
}
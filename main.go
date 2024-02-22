package main

import (
	"fmt"
)

type Node struct {
	Left

	Right
}

type Queue struct {
	Node
}

type Cache struct {
	Queue Queue
	Hash Hash
}

type Hash map[string]*Node

func main() {
	fmt.Println("Start Cache")

	cache := NewCache()

	for _, word := range []string{"banana", "papaya", "potato", "tomato", "tree", "cat"} {
		cache.Check(word)
		cache.Display()
	}
}
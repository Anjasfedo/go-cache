package main

import (
	"fmt"
)

type Node struct {
	Left  *Node
	Val   string
	Right *Node
}

type Queue struct {
	Head *Node
	Length int
	Tail *Node
}

type Cache struct {
	Queue Queue
	Hash  Hash
}

type Hash map[string]*Node

func NewCache() Cache {
	return Cache{Queue: NewQueue(), hash: Hash{}}
}

func NewQueue() Queue {

}

func main() {
	fmt.Println("Start Cache")

	cache := NewCache()

	for _, word := range []string{"banana", "papaya", "potato", "tomato", "tree", "cat"} {
		cache.Check(word)
		cache.Display()
	}
}

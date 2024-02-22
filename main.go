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
	Head   *Node
	Length int
	Tail   *Node
}

type Cache struct {
	Queue Queue
	Hash  Hash
}

type Hash map[string]*Node

func NewCache() Cache {
	return Cache{Queue: NewQueue(), Hash: Hash{}}
}

func NewQueue() Queue {
	head := &Node{}

	tail := &Node{}

	head.Right = tail
	tail.Left = head

	return Queue{Head: head, Tail: tail}
}

func main() {
	fmt.Println("Start Cache")

	cache := NewCache()

	for _, word := range []string{"banana", "papaya", "potato", "tomato", "tree", "cat"} {
		cache.Check(word)
		cache.Display()
	}
}

package main

import (
	"fmt"
)

const SIZE = 5

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

func (c *Cache) Check(str string) {
	node := &Node{}

	if val, ok := c.Hash[str]; ok {
		node = c.Remove(val)
	} else {
		node = &Node{Val: str}
	}

	c.Add(node)

	c.Hash[str] = node
}

func (c *Cache) Remove(n *Node) *Node {
	fmt.Printf("remove: %s\n", n.Val)

	left := n.Left
	right := n.Right

	left.Right = right
	right.Left = left

	c.Queue.Length--

	delete(c.Hash, n.Val)

	return n
}

func (c *Cache) Add(n *Node) {
	fmt.Printf("Add: %s\n", n.Val)

	tmp := c.Queue.Head.Right

	c.Queue.Head.Right = n

	n.Left = c.Queue.Head

	n.Right = tmp

	tmp.Left = n

	c.Queue.Length++

	if c.Queue.Length > SIZE {
		c.Remove(c.Queue.Tail.Left)
	}
}

func main() {
	fmt.Println("Start Cache")

	cache := NewCache()

	for _, word := range []string{"banana", "papaya", "potato", "tomato", "tree", "cat"} {
		cache.Check(word)
		cache.Display()
	}
}

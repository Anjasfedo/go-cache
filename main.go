package main // Package declaration, main is a special package for executable programs

import (
	"fmt" // Importing the fmt package for formatted I/O
)

const SIZE = 5 // Constant declaration for cache size

type Node struct { // Defining a struct type Node for binary tree nodes
	Left  *Node  // Pointer to the left child node
	Val   string // Value of the node
	Right *Node  // Pointer to the right child node
}

type Queue struct { // Defining a struct type Queue for a doubly linked list
	Head   *Node // Pointer to the head node
	Length int   // Length of the queue
	Tail   *Node // Pointer to the tail node
}

type Cache struct { // Defining a struct type Cache to hold the queue and hash map
	Queue Queue // Queue for managing nodes
	Hash  Hash  // Hash map for quick access to nodes
}

type Hash map[string]*Node // Defining a type Hash, which is a map from string keys to Node pointers

func NewCache() Cache { // Constructor function for creating a new cache
	return Cache{Queue: NewQueue(), Hash: Hash{}} // Initializing a new cache with an empty queue and hash map
}

func NewQueue() Queue { // Constructor function for creating a new queue
	head := &Node{} // Creating a new node for the head of the queue

	tail := &Node{} // Creating a new node for the tail of the queue

	head.Right = tail // Setting the right pointer of the head to the tail
	tail.Left = head  // Setting the left pointer of the tail to the head

	return Queue{Head: head, Tail: tail} // Returning the initialized queue
}

func (c *Cache) Check(str string) { // Method for checking if a string is in the cache
	node := &Node{} // Creating a new node

	if val, ok := c.Hash[str]; ok { // Checking if the string exists in the hash map
		node = c.Remove(val) // If yes, remove the corresponding node from the cache
	} else {
		node = &Node{Val: str} // If no, create a new node with the string value
	}

	c.Add(node) // Adding the node to the cache

	c.Hash[str] = node // Updating the hash map with the new node
}

func (c *Cache) Remove(n *Node) *Node { // Method for removing a node from the cache
	fmt.Printf("remove: %s\n", n.Val) // Printing a message indicating removal

	left := n.Left   // Storing the left pointer of the node
	right := n.Right // Storing the right pointer of the node

	left.Right = right // Adjusting the pointers to remove the node
	right.Left = left

	c.Queue.Length-- // Decreasing the length of the queue

	delete(c.Hash, n.Val) // Removing the node from the hash map

	return n // Returning the removed node
}

func (c *Cache) Add(n *Node) { // Method for adding a node to the cache
	fmt.Printf("Add: %s\n", n.Val) // Printing a message indicating addition

	tmp := c.Queue.Head.Right // Storing the current head's right pointer

	c.Queue.Head.Right = n // Setting the new node as the head's right child

	n.Left = c.Queue.Head // Adjusting pointers to add the node to the queue

	n.Right = tmp
	tmp.Left = n

	c.Queue.Length++ // Increasing the length of the queue

	if c.Queue.Length > SIZE { // Checking if the queue exceeds the cache size
		c.Remove(c.Queue.Tail.Left) // If yes, remove the least recently used node
	}
}

func (c *Cache) Display() { // Method for displaying the cache
	c.Queue.Display() // Delegating the display operation to the queue
}

func (q *Queue) Display() { // Method for displaying the queue
	node := q.Head.Right // Starting from the first node after the head

	fmt.Printf("%d - [", q.Length) // Printing the length of the queue

	for i := 0; i < q.Length; i++ { // Looping through each node in the queue
		fmt.Printf("{%s}", node.Val) // Printing the value of the node

		if i < q.Length-1 { // Checking if it's not the last node
			fmt.Printf("<-->") // Printing a connector
		}

		node = node.Right // Moving to the next node
	}

	fmt.Println("]") // Printing the closing bracket
}

func main() { // Main function, entry point of the program
	fmt.Println("Start Cache") // Printing a message indicating the start of the cache

	cache := NewCache() // Creating a new cache instance

	for _, word := range []string{"banana", "papaya", "potato", "tomato", "tree", "cat", "geprek"} { // Iterating over a list of words
		cache.Check(word) // Checking each word in the cache
		cache.Display()   // Displaying the cache after each word is checked
	}
}

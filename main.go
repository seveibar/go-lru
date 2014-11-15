package main

import (
	"fmt"
)

type Node struct {
	next, prev *Node
	value      int
}

type IntLRU struct {
	start, end *Node
	size       int
	nodes      map[int]*Node
}

type LRU interface {
	push(value int)
	mostRecent() int
	leastRecent() int
	print()
}

func (lru *IntLRU) push(value int) {
	if node, nodeExists := lru.nodes[value]; nodeExists {
		if node.prev != nil {
			node.prev.next = node.next
		}
		if node.next == nil {
			// It's the last node
			lru.end = node.prev
		} else {
			node.next.prev = node.prev
		}

		lru.start.prev = node
		node.prev = nil
		node.next = lru.start
		lru.start = node
	} else {
		node = &Node{lru.start, nil, value}
		lru.nodes[value] = node
		if lru.start != nil {
			lru.start.prev = node
		}
		lru.start = node
		if lru.end == nil {
			lru.end = lru.start
		}
	}

	for len(lru.nodes) > lru.size {

		nextEnd := lru.end.prev
		delete(lru.nodes, lru.end.value)
		lru.end = nextEnd
		lru.end.next = nil
	}
}

func (lru *IntLRU) leastRecent() int {
	return lru.end.value
}

func (lru *IntLRU) mostRecent() int {
	return lru.start.value
}

func (lru *IntLRU) print() {
	if lru.start != nil {
		fmt.Print(lru.start.value)
	}
	for node := lru.start.next; node != nil; node = node.next {
		fmt.Print(", ", node.value)
	}
	fmt.Println("")
}

func createLRU(size int) LRU {
	return &IntLRU{nil, nil, size, make(map[int]*Node)}
}

func main() {
	m := createLRU(4)

	m.push(1)
	m.push(2)
	m.push(3)
	m.push(4)
	m.push(5)
	m.push(4)

	m.print()
	fmt.Println("Most Recent: ", m.mostRecent())
	fmt.Println("Least Recent: ", m.leastRecent())
}

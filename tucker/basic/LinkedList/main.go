package main

import "fmt"

type Node struct {
	next *Node
	val  int
}

type LinkedList struct {
	root *Node
	tail *Node
}

func (l *LinkedList) AddNode(val int) {
	if l.root == nil {
		l.root = &Node{val: val}
		l.tail = l.root
		return
	}
	l.tail.next = &Node{val: val}
	l.tail = l.tail.next
}

func (l *LinkedList) RemoveNode(node *Node) {
	if node == l.root {
		l.root = l.root.next
		node.next = nil
		return
	}
	prev := l.root
	for prev.next != node {
		prev = prev.next
	}

	if node == l.tail {
		prev.next = nil
		l.tail = prev
	} else {
		prev.next = prev.next.next
	}
	node.next = nil
}

func (l LinkedList) PrintNodes() {
	for node.next != nil {
		fmt.Print(node.val, " -> ")
		node = node.next
	}
	fmt.Println(node.val)
}

func main() {
	ll := LinkedList{root: }

	root = &Node{val: 0}
	tail = root

	for i := 1; i < 10; i++ {
		tail = AddNode(tail, i)
	}

	PrintNodes(root)
	root, tail = RemoveNode(root.next.next, root, tail)
	PrintNodes(root)
	root, tail = RemoveNode(root, root, tail)
	PrintNodes(root)
	root, tail = RemoveNode(tail, root, tail)
	PrintNodes(root)
}

func PrintNodes(node *Node) {
	for node.next != nil {
		fmt.Print(node.val, " -> ")
		node = node.next
	}
	fmt.Println(node.val)
}

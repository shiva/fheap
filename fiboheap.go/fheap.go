// Copyright 2010 Shivanand Velmurugan. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This package provides fibonacci heap operations for any type that implements
// fheap.Interface.
//
package main

import (
	"container/list"
	"fmt"
)

type FHeap struct {
	minKeyRoot	*FHeapNode
	count		int
	maxDegree	int
}

type FHeapNode struct {
	degree			int
	mark			bool
	prev, next		*FHeapNode
	child, parent	*FHeapNode
	key				int
	Value			interface{}		// untouched
}

func (n *FHeapNode) addLast(destNode *FHeapNode) {
	// check for tail
	n.prev.next = destNode
	destNode.prev = n.prev
	destNode.next = n
	n.prev = destNode
}

func (n *FHeapNode) Next() *FHeapNode {
	return n.next
}

func (n *FHeapNode) print() {
	fmt.Print(n.key)
}

func MakeHeap() *FHeap {
	fh := new(FHeap)
	return fh
}

func (fh *FHeap) min() *FHeapNode {
	return fh.minKeyRoot;
} 

func (fh *FHeap) insert(key int, value interface{}) {
	node := new(FHeapNode)
	node.key = key
	node.Value = value
	node.prev = node
	node.next = node
	
	if fh.minKeyRoot == nil {
		fh.minKeyRoot = node
	} else {
		fh.minKeyRoot.addLast(node)
		if key < fh.minKeyRoot.key {
			// this is minkey
			fh.minKeyRoot = node
		}
	}
	fh.count++
}

func (root *FHeapNode) printTree() {	
	q := list.New()	
	q.PushBack(root)
	
	prevParent := root.parent
	for e := q.Front(); e != nil; e = e.Next() {
		treeNode := e.Value.(*FHeapNode)
		
		if (treeNode.parent != prevParent) {
			fmt.Println()
		}
		prevParent = treeNode.parent
		treeNode.print()
		mainchild := treeNode.child
		if mainchild != nil {
			q.PushBack(mainchild)
			
			for c := mainchild.next; c != mainchild; c = c.next {
				// add all children to queue
				q.PushBack(c)
			}
		
		}
	}
}


/* function to test FHeap and FHeapNode behaviour */
func test_FHeapNode_print() {
	fmt.Println("test_FHeapNode_print: BEGIN")
	
	n := new(FHeapNode)
	n.key = 1
	n.prev = n
	n.next = n
	
	n.print()
	fmt.Println()
	fmt.Println("test_FHeapNode_print: SUCCESS")
}

func test_printTree() {
	fmt.Println("test_printTree: BEGIN")
	
	// create tree
	n := new(FHeapNode)
	n.key = 1
	
	// level 1 children
	l1n1 := new(FHeapNode)
	l1n1.key = 2
	l1n1.parent = n
	n.child = l1n1
	
	l1n2 := new(FHeapNode)
	l1n2.key = 3
	l1n2.parent = n
	
	l1n1.prev = l1n2
	l1n1.next = l1n2
	l1n2.prev = l1n1
	l1n2.next = l1n1
	// level 2 children
	
	l2n1 := new(FHeapNode)
	l2n1.key = 4
	
	l2n2 := new(FHeapNode)
	l2n2.key = 5
	
	l2n1.prev = l2n2
	l2n1.next = l2n2
	l2n2.prev = l2n1
	l2n2.next = l2n1
	l1n1.child = l2n1
	l2n1.parent = l1n1
	l2n2.parent = l1n1
	
	
	l2n3 := new(FHeapNode)
	l2n3.key = 6
	
	l2n4 := new(FHeapNode)
	l2n4.key = 7
	
	l2n3.prev = l2n4
	l2n3.next = l2n4
	l2n4.prev = l2n3
	l2n4.next = l2n3
	l1n2.child = l2n3
	l2n3.parent = l1n2
	l2n4.parent = l1n2
	
	// print
	n.printTree()
	fmt.Println()
	fmt.Println("test_printTree: SUCCESS")
}

func test_FHeap_insert() {
	// test for min root, and insertion order
	fmt.Println("test_FHeap_insert: BEGIN")
	fh := new(FHeap)
	
	fh.insert(2, nil)
	fh.insert(5, nil)
	fh.insert(6, nil)
	fh.insert(1, nil)
	fh.insert(4, nil)
	fh.insert(3, nil)
			
	minrt := fh.min()
	fmt.Println("min root: ", minrt.key)

	n := minrt
	fmt.Print("roots : ", n.key)
	for e := n.Next(); e != n; e = e.Next() {
		fmt.Print(", ")
		fmt.Print(e.key)
	}
	fmt.Println()
	fmt.Println("test_FHeap_insert: SUCCESS")
}

func main() {
	test_FHeapNode_print()
	fmt.Println()
	test_printTree()
	fmt.Println()
	test_FHeap_insert()
}
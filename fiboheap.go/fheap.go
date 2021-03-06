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

func (n *FHeapNode) addChild(child *FHeapNode) {
	if (n.child != nil) {
		n.child.addLast(child)
	} else {
		child.prev = child
		child.next = child
		n.child = child
	}
}

func (n *FHeapNode) Next() *FHeapNode {
	return n.next
}

// op called on the head of the list
// specify node to delete
// returns the head of the list (to allow for n.remove(n1).remove(n2) .. etc
func (root *FHeapNode) remove(node *FHeapNode) *FHeapNode {
	// ensure all reqd values are not null
	newroot := root
	if (node.prev == nil) {
		// deleting root
		node = nil
		newroot = node.next
	} else if (node.next == nil) {
		// deleting tail
		node = nil
	} else {
		// both are not null, disconnect
		node.prev.next = node.next
		node.next.prev = node
		node = nil
	}	
	
	return newroot
}

func (n *FHeapNode) print() {
	fmt.Print(n.key)
}



/* Fibonacci Heap implementation */
func MakeHeap() *FHeap {
	fh := new(FHeap)
	return fh
}

func (fh *FHeap) min() *FHeapNode {
	return fh.minKeyRoot;
} 

func (fh *FHeap) add_to_root_list(node *FHeapNode) {
	if fh.minKeyRoot == nil {
		fh.minKeyRoot = node
	} else {
		fh.minKeyRoot.addLast(node)
		if node.key < fh.minKeyRoot.key {
			// this is minkey
			fh.minKeyRoot = node
		}
	}
}

func (fh *FHeap) insert(key int, value interface{}) {
	node := new(FHeapNode)
	node.key = key
	node.Value = value
	node.prev = node
	node.next = node

	fh.add_to_root_list(node)
	fh.count++
}

// add y as a child of x (remove from root list)
func (fh *FHeap) link(y *FHeapNode, x *FHeapNode) {
	// remove y from the root list of H
	if (y.parent != nil) {
		fmt.Println("node to link, cannot have a parent")
	}
	
	fh.minKeyRoot.remove(y)
	// make y a child of x
	y.parent = x
	x.addChild(y)
	fh.maxDegree++
	y.mark = false	
}

func (x *FHeapNode) swap(y *FHeapNode) {
	x_prev := x.prev
	x_next := x.next
	
	x.prev = y.prev
	x.next = y.next
	y.prev = x_prev
	y.next = x_next
}

func (fh *FHeap) consolidate() {
	var degArr [10] *FHeapNode
	for w := fh.minKeyRoot; w != nil; w = w.Next() {
		x := w
		d := x.degree
		for degArr[d] != nil {
			y := degArr[d]
			if ( x.key > y.key) {
				x.swap(y)
			}
			fh.link(y, x)
			degArr[d] = nil
		}
		degArr[d] = x
	}
	fh.minKeyRoot = nil
	for i := 0; i < fh.maxDegree; i++ {
		if degArr[i] != nil {
			fh.add_to_root_list(degArr[i])
		}
	}
}

func (fh *FHeap) extract_min() *FHeapNode {
	z := fh.min()
	// add each child of z to root list
	first_child := z.child
	first_child.parent = nil
	fh.add_to_root_list(first_child)
	for c := first_child.next; c != first_child; c = c.next {
		fh.add_to_root_list(c)
		c.parent = nil
	}
	
	// delete z from root list
	z.prev.next = z.next
	z.next.prev = z.prev
	
	if z == z.next {
		fh.minKeyRoot = nil
	} else {
		fh.minKeyRoot = z.next
		fh.consolidate()
	}
	fh.count--
	
	return z
}

/* helper to print a tree, starting from the root list in a f-heap */
func (root *FHeapNode) printTree() {	
	q := list.New()	
	marker := new(FHeapNode)
	marker.degree = -1
	q.PushBack(marker)
	q.PushBack(root)

	prevNode := root.parent
	for e := q.Front(); e != nil; e = e.Next() {
		treeNode := e.Value.(*FHeapNode)
	
		if treeNode.degree == -1 {		// marker node
			// requeue at the end to indicate end of next level
			if prevNode != nil && prevNode.degree == -1 {
				break
			}
			
			fmt.Println("")
			q.PushBack(treeNode)
			prevNode = treeNode
			continue
		} 	

		if prevNode != nil && prevNode.parent != nil && prevNode.parent.degree != -1 && prevNode.parent != treeNode.parent {
			fmt.Print(", ")
		} else {
 			fmt.Print(" ")
		}
				
		treeNode.print()
		

		mainchild := treeNode.child
		if mainchild != nil {
			q.PushBack(mainchild)

			for c := mainchild.next; c != mainchild; c = c.next {
				// add all children to queue
				q.PushBack(c)
			}

		}
		
		prevNode = treeNode
	}
}

/* Tests */

/* f-heap tests */
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
	
	fmt.Println("Tree for root: ", minrt.key)
	minrt.printTree()
	fmt.Println()
	
	for e := n.Next(); e != n; e = e.Next() {
		fmt.Println("Tree for root: ", e.key)
		e.printTree()
		fmt.Println()
	}
	fmt.Println("test_FHeap_insert: SUCCESS")
}


func test_extractmin() {
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
	
	fmt.Println("Tree for root: ", minrt.key)
	minrt.printTree()
	fmt.Println()
	
	for e := n.Next(); e != n; e = e.Next() {
		fmt.Println("Tree for root: ", e.key)
		e.printTree()
		fmt.Println()
	}
	
	
	fmt.Println("test extract min .... ")
	test := fh.extract_min()
	fmt.Println("Extacted min key ", test.key)
	
	
	
	
	
	
	minrt = fh.min()
	fmt.Println("min root: ", minrt.key)

	n = minrt
	fmt.Print("roots : ", n.key)
	for e := n.Next(); e != n; e = e.Next() {
		fmt.Print(", ")
		fmt.Print(e.key)
	}
	fmt.Println()
	
	fmt.Println("Tree for root: ", minrt.key)
	minrt.printTree()
	fmt.Println()
	
	for e := n.Next(); e != n; e = e.Next() {
		fmt.Println("Tree for root: ", e.key)
		e.printTree()
		fmt.Println()
	}
	
	
	
	fmt.Println("test_FHeap_insert: SUCCESS")
}

/* test for printTree */
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


/* main */
func main() {
	test_FHeapNode_print()	
	fmt.Println()
	test_printTree()
	fmt.Println()
	test_FHeap_insert()
}
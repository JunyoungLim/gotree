package gotree

import (
	"log"
)

// BST is a struct for a Binary Search Tree.
type BST struct {
	root *BSTNode
}

// BSTNode is a Binary Search Tree Node.
type BSTNode struct {
	value Element
	left  *BSTNode
	right *BSTNode
}

// NewBST returns a new Binary Search Tree created from the given
// list of elements.
func NewBST(elems ...Element) *BST {
	tree := &BST{}
	for _, e := range elems {
		tree.Insert(e)
	}
	return tree
}

// Insert inserts a new element to the head BST node.
func (tree *BST) Insert(elem Element) {
	node := &BSTNode{value: elem}
	tree.InsertNode(node)
}

// InsertNode inserts a new BST node to the head node.
func (tree *BST) InsertNode(node *BSTNode) {
	if tree == nil {
		log.Printf("Insertion of %+v skipped because the tree is nil.\n", node)
		return
	}
	if tree.root == nil {
		tree.root = node
		return
	}
	tree.root.insert(node)
}

// insert is an insertion operation for a BSTNode.
func (head *BSTNode) insert(node *BSTNode) {
	if head == nil {
		log.Printf("Insertion of %+v skipped because the head is nil.\n", node)
		return
	}

	if head.value.Less(node.value) {
		if head.right == nil {
			head.right = node
		} else {
			head.right.insert(node)
		}
	} else if node.value.Less(head.value) {
		if head.left == nil {
			head.left = node
		} else {
			head.left.insert(node)
		}
	}
}

// Delete deletes the node with the given element as its value.
func (tree *BST) Delete(elem Element) {
	tree.root = tree.root.delete(elem)
}

// delete is a deletion operation for a BSTNode.
// It returns the updated node, which is a replacement for the given head node.
func (head *BSTNode) delete(elem Element) *BSTNode {
	if head == nil {
		return head
	}

	if elem.Less(head.value) {
		head.left = head.left.delete(elem)
	} else if head.value.Less(elem) {
		head.right = head.right.delete(elem)
	} else {
		if head.left == nil {
			return head.right
		} else if head.right == nil {
			return head.left
		}

		head.value = head.right.minValue()
		head.right = head.right.delete(head.value)
	}

	return head
}

// minValue finds a node with the minimum value under the given BST node
// and returns its value.
func (head *BSTNode) minValue() Element {
	minVal := head.value
	for head.left != nil {
		head = head.left
		minVal = head.value
	}
	return minVal
}

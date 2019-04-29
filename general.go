package gotree

import "errors"

// GeneralNode defines a tree node for a general purpose tree.
type GeneralNode struct {
	value    Element
	children []*GeneralNode
	parent   *GeneralNode
}

// NewGeneralNode returns a new general tree node with the given value
// and the given list of child nodes, if any.
func NewGeneralNode(value Element, children ...*GeneralNode) *GeneralNode {
	return &GeneralNode{
		value:    value,
		children: children,
	}
}

// Parent returns the parent tree node of this node.
func (node *GeneralNode) Parent() *GeneralNode {
	return node.parent
}

// Children returns the list of child tree nodes of this node.
func (node *GeneralNode) Children() []*GeneralNode {
	return node.children
}

// Value returns the element value stored in this node.
func (node *GeneralNode) Value() Element {
	return node.value
}

// SetValue sets the new value to this node.
func (node *GeneralNode) SetValue(value Element) {
	node.value = value
}

// NumberOfChildren returns the total number of child nodes of this node.
func (node *GeneralNode) NumberOfChildren() int {
	return len(node.children)
}

// HasChildren returns true if this node has one or more child nodes,
// false otherwise.
func (node *GeneralNode) HasChildren() bool {
	return node.NumberOfChildren() > 0
}

// NumberOfDescendants returns the total number of descendant nodes of this node.
func (node *GeneralNode) NumberOfDescendants() int {
	count := node.NumberOfChildren()
	for _, child := range node.children {
		count += child.Size()
	}
	return count
}

// Size returns the total number of nodes under this node, including itself.
func (node *GeneralNode) Size() int {
	return node.NumberOfDescendants() + 1
}

// AddChild inserts a child node to this node.
func (node *GeneralNode) AddChild(child *GeneralNode) {
	child.parent = node
	node.children = append(node.children, child)
}

// AddChildAt inserts a child node to this node at the given index.
func (node *GeneralNode) AddChildAt(child *GeneralNode, index int) error {
	if index < 0 || index > len(node.children) {
		return errors.New("index out of range")
	}
	child.parent = node
	node.children = append(node.children, nil)
	copy(node.children[index+1:], node.children[index:])
	return nil
}

// RemoveChildren resets the list of child nodes.
func (node *GeneralNode) RemoveChildren() {
	node.children = nil
}

// RemoveChildAt deletes the child node at the given index.
func (node *GeneralNode) RemoveChildAt(index int) error {
	if index < 0 || index >= len(node.children) {
		return errors.New("index out of range")
	}
	copy(node.children[index:], node.children[index+1:])
	node.children[len(node.children)-1] = nil
	node.children = node.children[:len(node.children)-1]
	return nil
}

// GetChildAt returns the child node at the given index.
func (node *GeneralNode) GetChildAt(index int) (*GeneralNode, error) {
	if index < 0 || index >= len(node.children) {
		return nil, errors.New("index out of range")
	}
	return node.children[index], nil
}

// Find searches for the descendant node containing the given element value.
// If found, it returns the node. If not, it returns a nil pointer.
func (node *GeneralNode) Find(elem Element) *GeneralNode {
	if node == nil {
		return nil
	}
	if node.value.Equals(elem) {
		return node
	}
	for _, child := range node.children {
		if match := child.Find(elem); match != nil {
			return match
		}
	}
	return nil
}

// Equals returns true if the given node equals to this node, false otherwise.
func (node *GeneralNode) Equals(other *GeneralNode) bool {
	if node == nil || other == nil {
		return node == other
	}
	return node.value.Equals(other.value)
}

// GeneralTree defines a general purpose tree.
// It can have any number of child nodes, without any invariants or restrictions.
type GeneralTree struct {
	root *GeneralNode
}

// NewGeneralTree returns a new general tree with the given root node.
func NewGeneralTree(root *GeneralNode) *GeneralTree {
	return &GeneralTree{root}
}

// Root returns the root node of this tree.
func (tree *GeneralTree) Root() *GeneralNode {
	return tree.root
}

// SetRoot sets the new root node to this tree.
func (tree *GeneralTree) SetRoot(root *GeneralNode) {
	tree.root = root
}

// Size returns the total number of nodes under this tree.
func (tree *GeneralTree) Size() int {
	if tree.root == nil {
		return 0
	}
	return tree.root.Size()
}

// Find searches for the node of this tree containing the given element value.
// If found, it returns the node. If not, it returns a nil pointer.
func (tree *GeneralTree) Find(elem Element) *GeneralNode {
	return tree.root.Find(elem)
}

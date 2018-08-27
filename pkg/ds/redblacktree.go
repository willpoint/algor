package ds

type color int

const (
	black color = iota
	red
)

// RedBlackTree RBT is basically a BST with a different
// property that ensures it is always balanced
// with the height of the tree being O(log n)
// The additional properties of an RBT are
// Color and Black Height of a given node x `bh(x)`
// The properties of an RBT are:
// 1. Every node is either a black or red node
// 2. The root node is a black node
// 3. The leaf nodes (NIL) are black
// 4. A red node has two black children
// 5. Every simple path from a node x to the
//    descendant nodes have the same number of black nodes
// A node with no parent or child node have their corresponding
// pointers pointing to NIL - which are always black in the RBT
type RedBlackTree struct {
	Root *RBNode
}

// RBNode is an RBT Node
type RBNode struct {
	Color       color
	blackHeight int
	Parent      *RBNode
	Left        *RBNode
	Right       *RBNode
}

package ds

import (
	"fmt"
	"math"
)

// BNode is a binary search tree node
type BNode struct {
	Key    int
	Parent *BNode
	Left   *BNode
	Right  *BNode

	// flag is used to help improve the asymptotic
	// performance during inserts when n identical
	// inserts are made to a an initially empty BST
	flag bool
}

// BST is a binary search tree with root
// A BST has an invariant - that for all nodes x:
// if y is in the left subtree of x
// the Key(y) <= Key(x)
// and if y is in the right subtree of x
// then Key(y) >= Key(x)
type BST struct {
	Root *BNode
	Num  int // number of nodes in the tree
}

// NewBST ...
func NewBST() *BST {
	return &BST{}
}

// Insert inserts into the BST
func (t *BST) Insert(e int) {

	var root *BNode
	nn := &BNode{
		Key: e,
	}

	t.Num++
	curr := t.Root
	for curr != nil {
		root = curr
		if curr.Key > nn.Key {
			curr = curr.Left
		} else {
			curr = curr.Right
		}
	}
	nn.Parent = root
	if root == nil {
		t.Root = nn
	} else if root.Key > nn.Key {
		root.Left = nn
	} else {
		root.Right = nn
	}
}

// InsertOptimized inserts a new node with key e
// Insert tries to improve the asymptotic performance
// when n items with identical keys are inserted
// into an initially empty binary BST
// by using a strategy that checks if the key e is equal to
// the the its parent node
// by using a boolean flag - a local property of the
// parent node curr -
func (t *BST) InsertOptimized(e int) {

	var root *BNode
	nn := &BNode{
		Key: e,
	}

	t.Num++
	curr := t.Root
	for curr != nil {
		root = curr
		if nn.Key == curr.Key {
			curr.flag = !curr.flag
			if curr.flag {
				curr = curr.Right
			} else {
				curr = curr.Left
			}
		} else if nn.Key < curr.Key {
			curr = curr.Left
		} else {
			curr = curr.Right
		}
	}
	nn.Parent = root
	if root == nil {
		t.Root = nn
	} else if nn.Key == root.Key {
		root.flag = !root.flag
		if root.flag {
			root.Left = nn
		} else {
			root.Right = nn
		}
	} else if nn.Key < root.Key {
		root.Left = nn
	} else {
		root.Right = nn
	}
}

// Min returns the min key of the bst
func (t *BST) Min() int {
	curr := t.Root
	for curr.Left != nil {
		curr = curr.Left
	}
	return curr.Key
}

// Max returns the max key of the bst
func (t *BST) Max() int {
	curr := t.Root
	for curr.Right != nil {
		curr = curr.Right
	}
	return curr.Key
}

// InOrderWalk ...
func (t *BST) InOrderWalk() []int {
	var r []int
	var walk func(*BNode)
	walk = func(bn *BNode) {
		if bn != nil {
			walk(bn.Left)
			r = append(r, bn.Key)
			walk(bn.Right)
		}
	}
	walk(t.Root)
	return r
}

// Height computes the height of a bst from its root
func (t *BST) Height() float64 {
	var height func(bn *BNode) float64
	height = func(bn *BNode) float64 {
		if bn == nil {
			return -1
		}
		return 1 + math.Max(height(bn.Left), height(bn.Right))
	}
	return height(t.Root)
}

// Search returns the node with the key k or nil
// if no node exists with key k
func (t *BST) Search(k int) *BNode {
	curr := t.Root
	for curr != nil && curr.Key != k {
		if k < curr.Key {
			curr = curr.Left
		} else {
			curr = curr.Right
		}
	}
	return curr
}

// Delete deletes a node is a bst
func (t *BST) Delete(key int) error {
	bn := t.Search(key)
	if bn == nil {
		return fmt.Errorf("node with key %d does not exist", key)
	}
	if bn.Left == nil {
		// left node is nil (a single child)
		t.transplant(bn, bn.Right)
	} else if bn.Right == nil {
		// right node is nil (a single child)
		t.transplant(bn, bn.Left)
	} else {
		// both right and left subtrees are available
		s := treeSuccessor(bn)
		if s.Parent != bn {
			// the successor has min key at this point
			// and cannot have a left child otherwise it wont be the successor
			// so we transplant it's right child to take it's position on that subtree
			t.transplant(s, s.Right)
			s.Right = bn.Right
			s.Right.Parent = s
		}
		t.transplant(bn, s)
		s.Left = bn.Left
		s.Left.Parent = bn
	}
	return nil
}

// helper methods for bst
// ==============================================================

// isLeftChild checks if the given node
// is a left child of its parent
// the node should not be the root node
func (bn *BNode) isLeftChild() bool {
	return bn.Parent != nil && bn == bn.Parent.Left
}

// isRightChild checks if the given node
// is the right child of its parent
// the node should not be the root node
func (bn *BNode) isRightChild() bool {
	return bn.Parent != nil && bn == bn.Parent.Right
}

// isRoot checks if node is root
func isRoot(bn *BNode) bool {
	return bn.Parent == nil
}

// min from x returns the node with the min key
// rooted at x
func findMin(bn *BNode) *BNode {
	for bn.Left != nil {
		bn = bn.Left
	}
	return bn
}

// findMax is symmetric to findMin
func findMax(bn *BNode) *BNode {
	for bn.Right != nil {
		bn = bn.Right
	}
	return bn
}

// treeSuccessor of x returns the node whose key
// is the smallest key greater than x.key
// it returns nil if x.key is the greatest in the tree
// successor aims to give the next node in a bst in the form
// that is produced when an inorder tree walk is performed on the tree
// for example when an inorder tree walk is done and it produces
// a, b, c, d in sorted order - then the successor of a is b, and of b is c, etc
func treeSuccessor(bn *BNode) *BNode {
	if bn.Right == nil && bn.Left != nil { // if bn has the node with the max key return nil
		return nil
	}
	if bn.Right != nil { // if bn has a right child then the smallest
		return findMin(bn.Right)
	}
	// if bn has no right subtree, we move up the tree till we find a node
	// whose parent is not a right child of it's parent
	currParent := bn.Parent
	for currParent != nil && bn == currParent.Right {
		bn = currParent
		currParent = currParent.Parent
	}
	return currParent
}

// transplant replaces the the subtree rooted at u
// with the subtree rooted at v
// setting node u's parent pointing to v as it's new child
// and v pointing to u's parent as it new parent
func (t *BST) transplant(u, v *BNode) {
	if u.Parent == nil { // case where u is the root node
		t.Root = v
	} else if u.isLeftChild() {
		u.Parent.Left = v
	} else {
		u.Parent.Right = v
	}
	if v != nil {
		v.Parent = u.Parent
	}
}

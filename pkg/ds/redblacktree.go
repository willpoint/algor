package ds

import "errors"

type color int

const (
	black color = iota
	red
)

var (
	// ErrRotatingNode ...
	ErrRotatingNode = errors.New("errors rotating node")
)

// RBT is basically a BST with a different
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
type RBT struct {
	Root *RBNode

	// sentinel to conserve space
	Nil *RBNode
}

// NewRBT constructs a new RBT
func NewRBT() *RBT {
	n := &RBNode{}
	return &RBT{
		Nil: n,
	}
}

// RBNode is an RBT Node
type RBNode struct {
	Key    int
	Parent *RBNode
	Left   *RBNode
	Right  *RBNode

	color       color
	blackHeight int
}

// NewRBNode constructs a new RBNode
// it sets the left and right children to t.Nil
func (t *RBT) NewRBNode(key int) *RBNode {
	n := &RBNode{}
	n.Key = key
	n.Left = t.Nil
	n.Right = t.Nil
	return n
}

// left rotate pivots around the link between u and v
// and sets v to replace u as the root of u's subtree,
// sets u as v's left child and set v's left child as
// u's right child.
// leftRotate assumes u's right node is not t.Nil
func (t *RBT) leftRotate(u *RBNode) error {
	if u.Right == t.Nil {
		return ErrRotatingNode
	}
	v := u.Right
	u.Right = v.Left
	if v.Left != t.Nil {
		v.Left.Parent = u
	}

	v.Parent = u.Parent
	if u.Parent == t.Nil {
		t.Root = v
	} else if u == u.Parent.Left {
		u.Parent.Left = v
	} else {
		u.Parent.Right = v
	}
	v.Left = u
	u.Parent = v
	return nil
}

// rightRotate is the inverse of leftRotate
func (t *RBT) rightRotate(u *RBNode) error {
	if u.Left == t.Nil {
		return ErrRotatingNode
	}
	v := u.Left
	u.Left = v.Right
	if v.Right != t.Nil {
		v.Right.Parent = u
	}

	v.Parent = u.Parent
	if u.Parent == t.Nil {
		t.Root = v
	} else if u == v.Parent.Left {
		u.Parent.Left = v
	} else {
		u.Parent.Right = v
	}
	v.Right = u
	u.Parent = v
	return nil
}

// Insert ...
func (t *RBT) Insert(k int) {
	z := t.NewRBNode(k)
	y := t.Nil
	x := t.Root
	for x != nil {
		y = x
		if z.Key < x.Key {
			x = x.Left
		} else {
			x = x.Right
		}
	}
	z.Parent = y
	if y == t.Nil {
		t.Root = z
	} else if z.Key < y.Key {
		y.Left = z
	} else {
		y.Right = z
	}
	z.color = red
	InsertFixUp(t, z)
}

// InsertFixUp ...
func InsertFixUp(t *RBT, z *RBNode) {
	// since every newly inserted node is red
	// (resulting from it's two sentinels being black)
	// if its Parent is also red - a violation has occured
	// so we loop till that invariant ends
	for z.Parent.color == red {
		// check if z is a left child, if it is
		if z.Parent == z.Parent.Parent.Left {
			// set y as z's uncle
			// and check if z's uncle is red
			y := z.Parent.Parent.Right
			if y.color == red {
				// check if z's uncle (y) is red
				// and fix to meet red black tree's property
				// with a red parent and two black children
				z.Parent.color = black
				y.color = black
				z.Parent.Parent.color = red
				z = z.Parent.Parent // move z up for the next iteration
			} else {
				if z == z.Parent.Right {
					t.leftRotate(z)
				}
				z.Parent.color = black
				z.Parent.Parent.color = red
			}
		}
	}
	t.Root.color = black
}

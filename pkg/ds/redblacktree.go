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
	Ext *RBNode
}

// NewRBT constructs a new RBT
func NewRBT() *RBT {
	return &RBT{}
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
func (t *RBT) NewRBNode(key int) *RBNode {
	n := &RBNode{}
	n.Key = key
	n.Left = t.Ext
	n.Right = t.Ext
	return n
}

// left rotate pivots around the link between u and v
// and sets v to replace u as the root of u's subtree,
// sets u as v's left child and set v's left child as
// u's right child. leftRotate assumes u's right node
// is not t.ext
func (t *RBT) leftRotate(u *RBNode) error {
	if u.Right == t.Ext {
		return ErrRotatingNode
	}
	v := u.Right
	u.Right = v.Left
	if v.Left != t.Ext {
		v.Parent = u
	}

	if u.Parent == t.Ext {
		t.Root = v
	} else if u.Parent.Left == u {
		u.Parent.Left = v
	} else {
		u.Parent.Right = v
	}
	v.Left = u
	u.Parent = v
	return nil
}

// rightRotate is the inverse of leftRotate
func (t *RBT) rightRotate(v *RBNode) error {
	if v.Left == t.Ext {
		return ErrRotatingNode
	}
	u := v.Left
	u.Left = u.Right
	if u.Right != t.Ext {
		u.Parent = v
	}

	if v.Parent == t.Ext {
		t.Root = u
	} else if v.Parent.Left == v {
		v.Parent.Left = u
	} else {
		v.Parent.Right = u
	}
	v.Right = u
	u.Parent = v
	return nil
}

// Insert ...
func (t *RBT) Insert(k int) {
	z := t.NewRBNode(k)
	y := t.Ext
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
	if y == t.Ext {
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
    // while the node just inserted has a parent that is red
    for z.Parent.color == red {
        // 
        if z.Parent == z.Parent.Parent.Left
    }
}

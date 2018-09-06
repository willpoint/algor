package ds

// BNode is a binary search tree node
type BNode struct {
	Key    int
	Parent *BNode
	Left   *BNode
	Right  *BNode
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

// InOrderTreeWalk ...
func InOrderTreeWalk(t *BNode) []int {
	r := []int{}
	if t != nil {
		InOrderTreeWalk(t.Left)
		r = append(r, t.Key)
		InOrderTreeWalk(t.Right)
	}
	return r
}

// Insert ...
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

// Min returns the min key
func (t *BST) Min() int {
	curr := t.Root
	for curr.Left != nil {
		curr = curr.Left
	}
	return curr.Key
}

// Max returns the max key
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

// helper methods such as - if a node is a right or left child

// isLeftChild assumes the node checked has a parent that is not nil
// if node's parent p is nil - the node is root and the function
// returns false
func (bn *BNode) isLeftChild() bool {
	if bn.Parent == nil {
		return false
	}
	return bn.Parent.Left == bn
}

// isRightChild performs the converse of of isLeftChild
func (bn *BNode) isRightChild() bool {
	if bn.Parent == nil {
		return false
	}
	return bn.Parent.Right == bn
}

// transplant replaces a node u with a node v
// making u's parent the parent of v and making v the child of u's parent
func (t *BST) transplant(u, v *BNode) {
	if u.Parent == nil {
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

// // successor can be found using two cases
// // if there is a non empty right subtree - the node with
// // the minimum key is return
// // if there is no right subtree and the node in question u
// // has a successor v, then v is the lowest ancestor of u
// // whose left child is also an ancestor of x
// func noop() {}

// Delete a node in the bst provided as an argument (bn)
// The logic for this method is broken into 3
// if bn has no childern, then it is simply removed from the bst
// by modifying it's parent to point to nil
// if bn has a single child, the we elevate that child take bn's position
// by modifying bn's parent to point to bn's single child
// if bn has two children, then we find bn's successor u - which must be
// in the bn's right subtree and update u to take bn's position, and bn's
// left subtree becomes u's left subtree. It is important that u - the successor
// is in bn's right subtree - to maintain the binary tree property -
// u (right child) preceeds bn (parent) which preceeds v (left child)
// u < p < v --- therefore v must replace p resulting in u < v
func (t *BST) Delete(bn *BNode) {
	//
}

// Search traverses a subtree of a BST rooted at root
func Search(root *BNode, k int) *BNode {
	for root != nil && root.Key != k {
		if k < root.Key {
			root = root.Left
		} else {
			root = root.Right
		}
	}
	return root
}

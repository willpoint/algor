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

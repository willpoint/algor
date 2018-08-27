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
}

// NewBST ...
func NewBST() *BST {
	root := &BNode{}
	return &BST{root}
}

// Insert ...
func (t *BST) Insert(e int) {

	var parent *BNode

	nn := &BNode{
		Key: e,
	}

	currNode := t.Root
	for currNode != nil {
		parent = currNode
		if currNode.Key < nn.Key {
			parent = currNode.Left
		} else {
			parent = currNode.Right
		}
	}

	if parent == nil {
		t.Root = nn
	} else if parent.Key < nn.Key {
		nn.Parent = parent
		parent.Left = nn
	} else {
		nn.Parent = parent
		parent.Right = nn
	}
}

// Min returns the min key
func (t *BST) Min() int {
	curr := t.Root

	if curr.Left == nil {
		return 0
	}

	for curr.Left != nil {
		curr = curr.Left
	}
	return curr.Key
}

// Max returns the max key
func (t *BST) Max() int {
	curr := t.Root

	if curr.Right == nil {
		return 0
	}

	for curr.Right != nil {
		curr = curr.Right
	}
	return curr.Key
}

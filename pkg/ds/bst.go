package ds

// BNode is a binary search tree node
type BNode struct {
	E     string
	Left  *BNode
	Right *BNode
}

// BST is a binary search tree with root
type BST struct {
	Root *BNode
}

// NewBST ...
func NewBST(e string) *BST {
	root := &BNode{
		E: e,
	}
	return &BST{root}
}

// Insert ...
func (t *BST) Insert(e string) {

	var parent *BNode

	nn := &BNode{
		E: e,
	}

	currNode := t.Root
	for currNode != nil {
		parent = currNode
		if currNode.E < nn.E {
			parent = currNode.Left
		} else {
			parent = currNode.Right
		}
	}

	if parent == nil {
		t.Root = nn
	} else if parent.E < nn.E {
		parent.Left = nn
	} else {
		parent.Right = nn
	}
}

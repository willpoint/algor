package ds

import "testing"

func TestRBT_leftRotate(t *testing.T) {
	rbt := NewRBT()
	n1 := rbt.NewRBNode(10)
	n1.Parent = rbt.Ext // parent points to t.Ext
	n2 := rbt.NewRBNode(5)
	n3 := rbt.NewRBNode(13)
	rbt.Root = n1
	rbt.Root.Left = n2
	rbt.Root.Right = n3
	rbt.leftRotate(n1)
	// after left rotate
	// manually done - the root key becomes 13
	// node with key 5 is left of root
	// and node with key 13
	rootk := 13
	leftK := 5
	rightK := 10
	if rootk != rbt.Root.Key && leftK != rbt.Root.Left.Key && rbt.Root.Right.Key != rightK {
		t.Errorf("lefttree() expects %d, %d, %d", rootk, leftK, rightK)
	}

}

package ds

import "testing"

func TestRBT_leftRotate(t *testing.T) {
	// set up a manual rbt with a root and a left and right child
	rbt := NewRBT()
	n1 := rbt.NewRBNode(10) // root node
	n1.Parent = rbt.Nil     // parent points to t.Nil
	n2 := rbt.NewRBNode(5)
	n3 := rbt.NewRBNode(13)
	rbt.Root = n1
	rbt.Root.Left = n2
	rbt.Root.Right = n3

	// this should make the right child the new root
	// since n3 is the right child
	// the root.Key value should be equal to n3.Key after a left rotate
	if err := rbt.leftRotate(n1); err != nil {
		t.Errorf("error %#v\n", err)
	}
	if n3.Key != rbt.Root.Key {
		t.Errorf("leftRotate() expects new root key to be %d, got %d", n3.Key, rbt.Root.Key)
	}
}

func TestRBT_rightRotate(t *testing.T) {
	// set up a manual rbt with a root and a left and right child
	rbt := NewRBT()
	n1 := rbt.NewRBNode(97) // root node
	n1.Parent = rbt.Nil     // parent points to t.Nil
	n2 := rbt.NewRBNode(65)
	n3 := rbt.NewRBNode(100)
	rbt.Root = n1
	rbt.Root.Left = n2
	rbt.Root.Right = n3

	// this should make the left child the new root
	// since n2 is the left child
	// the root.Key value should be equal to n3.Key after a left rotate
	if err := rbt.rightRotate(n1); err != nil {
		t.Errorf("error %#v\n", err)
	}
	if n2.Key != rbt.Root.Key {
		t.Errorf("rightRotate() expects new root key to be %d, got %d", n2.Key, rbt.Root.Key)
	}
}

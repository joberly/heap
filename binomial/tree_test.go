package binomial

import (
	"fmt"
	"testing"
)

func intLess(a, b interface{}) bool {
	return a.(int) < b.(int)
}

func TestMerge(t *testing.T) {
	values := [][]int{
		[]int{1, 2},
		[]int{2, 1},
	}

	for _, vs := range values {
		na := newNode(vs[0])
		nb := newNode(vs[1])

		if na.t.rank() != 0 {
			t.Errorf("Tree rank incorrect: %+v\n", *na.t)
		}

		tname := fmt.Sprintf("testing %v", vs)
		t.Run(tname, func(t *testing.T) {
			checkMergedTree(t, merge(na.t, nb.t, intLess))
		})
	}
}

// Assumes tree items are 1 and 2.
func checkMergedTree(t *testing.T, mt *Tree) {
	if mt.rank() != 1 {
		t.Errorf("new Tree rank incorrect: %+v\n", mt)
	}

	if int(mt.n.Item.(int)) != 1 {
		t.Errorf("root Tree item incorrect: %+v\n", *mt)
	}

	if mt.sibling != nil {
		t.Errorf("root sibling not nil: %+v\n", *mt.sibling)
	}

	c := mt.child
	if c == nil {
		t.Fatalf("child Tree is nil\n")
		return
	}

	if c.parent != mt {
		t.Errorf("child's parent not root: %+v\n", *c)
	}

	if int(c.n.Item.(int)) != 2 {
		t.Errorf("child Tree item incorrect: %+v\n", *c)
	}

	if c.sibling != nil {
		t.Errorf("child sibling not nil: %+v\n", *c.sibling)
	}

	if c.child != nil {
		t.Errorf("child child not nil: %+v\n", *c.child)
	}
}

func TestMergeDiffRanks(t *testing.T) {
	n1 := newNode(1)
	n2 := newNode(2)
	n3 := newNode(3)

	t23 := merge(n2.t, n3.t, intLess)

	mt := merge(n1.t, t23, intLess)
	if mt != nil {
		t.Errorf("merge did not fail: %+v", *mt)
	}
}

func TestBubble(t *testing.T) {
	b := newNode(50)
	tree := merge(
		merge(b.t, newNode(20).t, intLess),
		merge(newNode(30).t, newNode(10).t, intLess),
		intLess,
	)

	// Expect tree to be as follows:
	//    10
	//  /  |
	// 20 30
	//  |
	// 50

	// Now set 50 item to 1 and see that it makes it all the way up
	// to the top of the tree, so the tree looks like this:
	//     1
	//  /  |
	// 10 30
	//  |
	// 20

	b.Item = 1
	b.bubble(intLess)

	if b.t != tree && b.t.parent != nil {
		t.Errorf("item not moved to root: %+v", *b)
	}

	if int(tree.n.Item.(int)) != 1 {
		t.Errorf("root item value incorrect: %+v", *tree)
	}
}

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
		ta := NewTree(vs[0])
		tb := NewTree(vs[1])

		if ta.Rank() != 0 {
			t.Errorf("Tree rank incorrect: %+v\n", *ta)
		}

		tname := fmt.Sprintf("testing %v", vs)
		t.Run(tname, func(t *testing.T) {
			checkMergedTree(t, Merge(ta, tb, intLess))
		})
	}
}

// Assumes tree items are 1 and 2.
func checkMergedTree(t *testing.T, mt *Tree) {
	if mt.Rank() != 1 {
		t.Errorf("new Tree rank incorrect: %+v\n", mt)
	}

	if int(mt.Item.(int)) != 1 {
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

	if int(c.Item.(int)) != 2 {
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
	t1 := NewTree(1)
	t2 := NewTree(2)
	t3 := NewTree(3)

	t23 := Merge(t2, t3, intLess)

	mt := Merge(t1, t23, intLess)
	if mt != nil {
		t.Errorf("merge did not fail: %+v", *mt)
	}
}

func TestBubble(t *testing.T) {
	b := NewTree(50)
	tree := Merge(
		Merge(b, NewTree(20), intLess),
		Merge(NewTree(30), NewTree(10), intLess),
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
	b = Bubble(b, intLess)

	if b != tree && b.parent != nil {
		t.Errorf("item not moved to root: %+v", *b)
	}

	if int(tree.Item.(int)) != 1 {
		t.Errorf("root item value incorrect: %+v", *tree)
	}
}

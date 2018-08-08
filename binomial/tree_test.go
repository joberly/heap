package binomial

import (
	"fmt"
	"testing"
)

// IntItem wraps int in a type which provides the Item interface.
type IntItem int

func (i IntItem) Less(j Item) bool {
	return int(i) < int(j.(IntItem))
}

func TestNewNilItem(t *testing.T) {
	tree := NewTree(nil)
	if tree != nil {
		t.Errorf("tree created with nil item: %+v", *tree)
	}
}

func TestMerge(t *testing.T) {
	values := [][]int{
		[]int{1, 2},
		[]int{2, 1},
	}

	for _, vs := range values {
		ta := NewTree(IntItem(vs[0]))
		tb := NewTree(IntItem(vs[1]))

		if ta.Rank() != 0 {
			t.Errorf("Tree rank incorrect: %+v\n", *ta)
		}

		tname := fmt.Sprintf("testing %v", vs)
		t.Run(tname, func(t *testing.T) {
			checkMergedTree(t, Merge(ta, tb))
		})
	}
}

// Assumes tree items are 1 and 2.
func checkMergedTree(t *testing.T, mt *Tree) {
	if mt.Rank() != 1 {
		t.Errorf("new Tree rank incorrect: %+v\n", mt)
	}

	if int(mt.Item.(IntItem)) != 1 {
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

	if int(c.Item.(IntItem)) != 2 {
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
	t1 := NewTree(IntItem(1))
	t2 := NewTree(IntItem(2))
	t3 := NewTree(IntItem(3))

	t23 := Merge(t2, t3)

	mt := Merge(t1, t23)
	if mt != nil {
		t.Errorf("merge did not fail: %+v", *mt)
	}
}

func TestBubble(t *testing.T) {
	b := NewTree(IntItem(50))
	tree := Merge(
		Merge(b, NewTree(IntItem(20))),
		Merge(NewTree(IntItem(30)), NewTree(IntItem(10))),
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

	b.Item = IntItem(1)
	b = Bubble(b)

	if b != tree && b.parent != nil {
		t.Errorf("item not moved to root: %+v", *b)
	}

	if int(tree.Item.(IntItem)) != 1 {
		t.Errorf("root item value incorrect: %+v", *tree)
	}
}
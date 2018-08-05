package binomial

// Item is the interface for data on a binomial tree and heap.
type Item interface {
	Less(Item) bool
}

// Tree is a binomial tree.
type Tree struct {
	// Item is the binomial tree Tree data.
	Item Item

	// k is the rank of the tree.
	k uint

	// Tree structure pointers.
	parent  *Tree  // direct parent
	sibling *Tree  // next sibling
	child   *Tree  // first child
}

// NewTree creates a new binomial tree with the specified Item.
func NewTree(item Item) *Tree {
	if item == nil {
		return nil
	}
	return &Tree{Item: item}
}

// Rank returns the rank of the tree.
func (t *Tree) Rank() uint {
	return t.k
}

// Merge combines two Trees of the same rank, returning the new binomial tree.
// This consumes n1 and n2 into the new tree.
func Merge(t1, t2 *Tree) *Tree {

	// It is up to the caller to understand that only Trees of the
	// same rank can be merged.
	if t1.k != t2.k {
		return nil
	}

	// Determine which Tree is the parent and which will be the child.
	tp := t1
	tc := t2
	if t2.Item.Less(t1.Item) {
		tp = t2
		tc = t1
	}

	// Save the original last child.
	c := tp.child
	// Make the new child Tree the last child.
	tp.child = tc
	// Make the new child Tree's first sibline the root's last child.
	tc.sibling = c
	// Ensure the new child Tree points to its parent.
	tc.parent = tp
	// Increase the rank of the parent now that it has a new child.
	tp.k++

	return tp
}

// Bubble moves an item up the tree if it is less than its successive parents.
// Returns the new Tree where the Item was placed.
func Bubble(t *Tree) *Tree {
	item := t.Item
	c := t
	p := t.parent
	for p != nil && item.Less(p.Item) {
		c.Item = p.Item
		p.Item = item
		c = p
		p = p.parent
	}

	return c
}

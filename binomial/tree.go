package binomial

// Node is a user data node on a binomial Tree.
type Node struct {
	Item interface{} // Consumer data

	t *Tree // Current tree for this Node
}

// Tree is a binomial tree.
type Tree struct {
	n *Node // Container for user data
	k uint  // Rank of the tree.

	// Tree structure pointers.
	parent  *Tree // direct parent
	sibling *Tree // next sibling
	child   *Tree // first child
}

// Less is a function that returns true if a is less than b.
type Less func(a, b interface{}) bool

// NewNode creates a new binomial tree with the specified item.
// It returns the Node created for the item.
func newNode(item interface{}) *Node {
	n := &Node{Item: item}
	t := &Tree{n: n}
	n.t = t
	return n
}

// Rank returns the rank of the tree.
func (t *Tree) rank() uint {
	return t.k
}

// Merge combines two Trees of the same rank, returning the new binomial tree.
// This consumes n1 and n2 into the new tree.
func merge(t1, t2 *Tree, less Less) *Tree {
	// It is up to the caller to understand that only Trees of the
	// same rank can be merged.
	if t1.rank() != t2.rank() {
		return nil
	}

	// Determine which Tree is the parent and which will be the child.
	tp := t1
	tc := t2
	if less(t2.n.Item, t1.n.Item) {
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
// The Node n now sits in the correct place in its Tree.
func (n *Node) bubble(less Less) {
	pt := n.t.parent
	for pt != nil && less(n.Item, pt.n.Item) {
		swap(n.t, pt)
		pt = n.t.parent
	}
}

// Swap exchanges the nodes between two trees.
func swap(t1 *Tree, t2 *Tree) {
	ntemp := t1.n
	t1.n = t2.n
	t2.n = ntemp
	t1.n.t = t1
	t2.n.t = t2
}

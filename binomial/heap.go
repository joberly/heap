// Package binomial contains binomial tree and heap data structures.
package binomial

import "container/list"

// Heap is a binomial heap.
type Heap struct {
	list list.List
	less Less
}

// NewHeap creates a new Heap with the given function to determine the lesser
// of two items on the heap. The Less function must handle each item added
// with Add.
func NewHeap(less Less) *Heap {
	return &Heap{less: less}
}

// Add places a new Item on the heap.
func (h *Heap) Add(i interface{}) *Node {
	// Add new item as rank 0 tree.
	n := newNode(i)
	h.combine(n.t)
	return n
}

// Combine looks through the existing trees on the heap, combining the new
// tree into another tree of the same rank. It repeats this process until
// the new combined tree has no other trees of the same rank in the list.
// The new tree is added to the front of the list.
func (h *Heap) combine(t *Tree) {
	tnew := t
	e := h.list.Front()
	for e != nil {
		tnext := e.Value.(*Tree)
		if tnew.rank() == tnext.rank() {
			h.list.Remove(e)
			tnew = merge(tnew, tnext, h.less)
			e = h.list.Front()
		} else {
			e = e.Next()
		}
	}
	h.list.PushFront(tnew)
}

// FindMin finds the minimum item on the heap.
func (h *Heap) FindMin() *Node {
	tmin, _ := h.findMin()
	if tmin == nil {
		return nil
	}
	return tmin.n
}

// FindMin finds the tree whose root has the minimum item on the heap.
func (h *Heap) findMin() (*Tree, *list.Element) {
	e := h.list.Front()
	var emin *list.Element
	var tmin *Tree
	for e != nil {
		t := e.Value.(*Tree)
		if tmin == nil || h.less(t.n.Item, tmin.n.Item) {
			emin = e
			tmin = t
		}
		e = e.Next()
	}
	return tmin, emin
}

// RemoveMin removes the minimum item from the heap.
func (h *Heap) RemoveMin() interface{} {
	// Find the minimum item and its tree's list element.
	tmin, emin := h.findMin()
	if tmin == nil {
		return nil
	}

	imin := tmin.n.Item

	// Remove the element from the list.
	h.list.Remove(emin)

	// Now combine all the children of tmin as distinct trees into the
	// existing trees on the heap's list.

	// Remove the root and get the first child tree.
	t := tmin.child
	for t != nil {
		// First save t's next sibling in the list of siblings.
		tnext := t.sibling

		// Disconnect t from its parent and siblings as its own root node.
		t.parent = nil
		t.sibling = nil

		// Combine t into other trees on the heap's list.
		h.combine(t)

		// Go to the next sibling of the original minimum item tree.
		t = tnext
	}

	return imin
}

// Update updates the item in a Heap when the value of the item changes.
// The results are undefined if Node n is not on this Heap.
func (h *Heap) Update(n *Node) {
	n.bubble(h.less)
}

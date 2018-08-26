package binomial

import (
	"math/rand"
	"testing"
	"time"
)

func TestHeapAddRemove(t *testing.T) {
	h := NewHeap(intLess)
	vs := []int{10, 20, 30, 30, 40, 40, 40, 50}

	vsinsert := make([]int, len(vs))
	copy(vsinsert, vs)

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for len(vsinsert) > 0 {
		i := r.Int() % len(vsinsert)
		v := vsinsert[i]
		n := h.Add(v)
		if n.Item.(int) != v {
			t.Errorf("Add returned bad Node: %v (expected Item %v)", *n, v)
		}
		vsinsert = append(vsinsert[:i], vsinsert[i+1:]...)
	}

	for _, v := range vs {
		wfind := h.FindMin()
		wrem := h.RemoveMin()
		if wfind.Item != wrem {
			t.Errorf("FindMin and RemoveMin not the same item %v %v", *wfind, wrem)
		}
		w := wrem.(int)
		if w != v {
			t.Errorf("unexpected item from heap %v (expected %v)", w, v)
		}
	}
}

func TestHeapUpdate(t *testing.T) {
	h := NewHeap(intLess)
	n := h.Add(50)
	h.Add(20)
	h.Add(30)
	h.Add(10)

	// Expect heap to be as follows:
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

	n.Item = 1
	h.Update(n)

	// Remove items in order
	expvs := []int{1, 10, 20, 30}
	for _, ev := range expvs {
		av := h.RemoveMin().(int)
		if av != ev {
			t.Errorf("Removed value out of order or unexpected %v (expected %v)", av, ev)
		}
	}

	// Check that no items are left
	niln := h.FindMin()
	if niln != nil {
		t.Errorf("Empty heap FindMin returned valid value %v", niln)
	}
	nilv := h.RemoveMin()
	if nilv != nil {
		t.Errorf("Empty heap RemoveMin returned valid value %v", nilv)
	}
}

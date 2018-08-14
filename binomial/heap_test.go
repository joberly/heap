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
		h.Add(v)
		vsinsert = append(vsinsert[:i], vsinsert[i+1:]...)
	}

	for _, v := range vs {
		w := int(h.RemoveMin().(int))
		if w != v {
			t.Errorf("unexpected item from heap %v (expected %v)", w, v)
		}
	}
}

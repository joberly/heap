# Package binomial

[![GoDoc](https://godoc.org/github.com/joberly/heap/binomial?status.svg)](https://godoc.org/github.com/joberly/heap/binomial)

## Tree Structure

Each box below is a Tree struct. Each Tree's _child_ pointer is its first child in a linked list of direct children. The subsequent child in the list is the current child's _sibling_. Each child has a pointer back to its _parent_.

The rank, k, of each node is the rank of the tree below that node if you consider that node a root.

This entire diagram represents a rank 2 or degree 2 binomial tree.

```
      +---------+
+----->Item: 10 |
|     +---------+
|     |k: 2     |
|     +---------+
|     |parent   |
|     |sibling  |
|  +--+child    |
|  |  +---------+
|  |
|  |  +---------+        +---------+
|  +-->Item: 20 | +------>Item: 12 <--+
|     +---------+ |      +---------+  |
|     |k: 0     | |      |k: 1     |  |
|     +---------+ |      +---------+  |
+-----+parent   | | +----+parent   |  |
|     |sibling  +-+ |    |sibling  |  |
|     |child    |   | +--+child    |  |
|     +---------+   | |  +---------+  |
|                   | |               |
+-------------------+ |               |
                      |  +---------+  |
                      +-->Item: 18 |  |
                         +---------+  |
                         |k: 0     |  |
                         +---------+  |
                         |parent   +--+
                         |sibling  |
                         |child    |
                         +---------+

```

## Heap Structure

A Heap contains a [standard library doubly linked list](https://golang.org/pkg/container/list) where the Value for each list Element is a pointer to a Tree.
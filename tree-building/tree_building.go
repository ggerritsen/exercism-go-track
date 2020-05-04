// Package tree provides functionality around storing forum posts in a tree
package tree

import (
	"fmt"
	"sort"
)

// Record is a forum post
type Record struct {
	// ID is the identifier of the Record
	ID int
	// Parent is the identifier of the parent of the Record
	Parent int
}

// Node is a way to provide a tree structure to Records
type Node struct {
	// ID is the identifier of the Node
	ID int
	// Children are the Node's child nodes
	Children []*Node
}

func (n *Node) add(r Record) {
	if n.ID == r.Parent {
		n.Children = append(n.Children, &Node{ID: r.ID})
		sort.Slice(n.Children, func(i, j int) bool {
			return n.Children[i].ID < n.Children[j].ID
		})
	}

	for _, c := range n.Children {
		c.add(r)
	}
}

// Build builds a tree of Nodes out of a set of Records
func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}

	sort.Slice(records, func(i, j int) bool {
		return records[i].ID < records[j].ID
	})

	if records[0].ID != 0 {
		return nil, fmt.Errorf("no root node: %q", records)
	}
	if records[0].Parent != 0 {
		return nil, fmt.Errorf("root node has invalid parent id: %q", records)
	}
	for i := 1; i < len(records); i++ {
		r := records[i]
		if r.ID != i {
			return nil, fmt.Errorf("non-continuous ids detected: %q", r)
		}
		if r.ID == r.Parent {
			return nil, fmt.Errorf("direct cycle detected: %q", r)
		}
		if r.Parent > r.ID {
			return nil, fmt.Errorf("invalid parent id: %q", r)
		}
		if r.ID == records[i-1].ID {
			return nil, fmt.Errorf("duplicate node detected: %q", r)
		}
	}

	root := &Node{ID: 0}

	for _, r := range records {
		if r.ID < 0 {
			return nil, fmt.Errorf("invalid id: %q", r)
		}
		if r.Parent < 0 {
			return nil, fmt.Errorf("invalid parent id: %q", r)
		}
		if r.ID == 0 {
			continue
		}

		if r.ID > 0 {
			root.add(r)
		}
	}

	return root, nil
}

package sort

import (
// "sort"
)

// type Interface = sort.Interface

type MapIntByte struct {
	Index uint
	Body  []byte
}

type MapIntInterface struct {
	Index uint
	Body  []interface{}
}

type MapSlice []*MapIntByte

func (m MapSlice) Len() int {
	return len(m)
}
func (m MapSlice) Less(i, j int) bool {
	return m[i].Index < m[j].Index
}

// Swap swaps the elements with indexes i and j.
func (m MapSlice) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

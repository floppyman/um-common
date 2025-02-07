package types

import "sort"

// IntSlice32 attaches the methods of Interface to []int32, sorting in increasing order.
type IntSlice32 []int32

func (x IntSlice32) Len() int           { return len(x) }
func (x IntSlice32) Less(i, j int) bool { return x[i] < x[j] }
func (x IntSlice32) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

// Sort is a convenience method: x.Sort() calls Sort(x).
func (x IntSlice32) Sort() { sort.Sort(x) }

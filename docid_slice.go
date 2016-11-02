// Generated by: main
// TypeWriter: slice
// Directive: +gen on DocID

package idx

import (
	"errors"
	"sort"
)

// DocIDSlice is a slice of type DocID. Use it where you would use []DocID.
type DocIDSlice []DocID

// Min returns the minimum value of DocIDSlice. In the case of multiple items being equally minimal, the first such element is returned. Returns error if no elements. See: http://clipperhouse.github.io/gen/#Min
func (rcv DocIDSlice) Min() (result DocID, err error) {
	l := len(rcv)
	if l == 0 {
		err = errors.New("cannot determine the Min of an empty slice")
		return
	}
	result = rcv[0]
	for _, v := range rcv {
		if v < result {
			result = v
		}
	}
	return
}

// Max returns the maximum value of DocIDSlice. In the case of multiple items being equally maximal, the first such element is returned. Returns error if no elements. See: http://clipperhouse.github.io/gen/#Max
func (rcv DocIDSlice) Max() (result DocID, err error) {
	l := len(rcv)
	if l == 0 {
		err = errors.New("cannot determine the Max of an empty slice")
		return
	}
	result = rcv[0]
	for _, v := range rcv {
		if v > result {
			result = v
		}
	}
	return
}

// Sort returns a new ordered DocIDSlice. See: http://clipperhouse.github.io/gen/#Sort
func (rcv DocIDSlice) Sort() DocIDSlice {
	result := make(DocIDSlice, len(rcv))
	copy(result, rcv)
	sort.Sort(result)
	return result
}

func (rcv DocIDSlice) Len() int {
	return len(rcv)
}
func (rcv DocIDSlice) Less(i, j int) bool {
	return rcv[i] < rcv[j]
}
func (rcv DocIDSlice) Swap(i, j int) {
	rcv[i], rcv[j] = rcv[j], rcv[i]
}

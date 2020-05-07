// Package sublist provides functionality around checking whether one list is a sublist of the other
package sublist

import "reflect"

// Relation represents the relationship between two lists
type Relation string

// Sublist returns the `Relation` between l1 and l2
func Sublist(l1, l2 []int) Relation {
	if reflect.DeepEqual(l1, l2) {
		return "equal"
	}

	if sublist(l1, l2) {
		return "sublist"
	}
	if sublist(l2, l1) {
		return "superlist"
	}

	return "unequal"
}

// sublist returns whether `listOne` is a sublist of `listTwo`
func sublist(listOne, listTwo []int) bool {
	if len(listOne) == 0 {
		return true
	}

	for i := range listTwo {
		if len(listTwo[i:]) < len(listOne) {
			return false
		}

		k := 0
		for j := i; j < len(listTwo) && k < len(listOne); j++ {
			if listTwo[j] != listOne[k] {
				break
			}
			k++
		}
		if k > len(listOne)-1 {
			return true
		}
	}
	return false
}

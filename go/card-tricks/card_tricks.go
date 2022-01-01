package cards

// GetItem retrieves an item from a slice at given position. The second return value indicates whether
// the given index exists in the slice or not.
func GetItem(slice []int, index int) (int, bool) {
	if index < 0 || index >= len(slice) {
		return 0, false
	} else {
		return slice[index], true
	}
}

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range the value needs to be appended.
func SetItem(slice []int, index, value int) []int {
	if index < 0 || index >= len(slice) {
		return append(slice, value)
	} else {
		slice[index] = value
		return slice
	}
}

// PrefilledSlice creates a slice of given length and prefills it with the given value.
func PrefilledSlice(value, length int) []int {
	if slice := []int{}; length <= 0 {
		return slice
	} else {
		for len(slice) < length {
			slice = append(slice, value)
		}
		return slice
	}
}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
	if index < 0 || index >= len(slice) {
		return slice
	} else {
		pre_index := slice[0:index]
		post_index := slice[index+1:]
		return append(pre_index, post_index...)
	}
}

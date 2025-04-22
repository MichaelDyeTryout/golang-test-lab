package sequence

func CountingNumbers(qty int8) (*[]int8, error) {
	// Return sequence of counting ints with size equal to qty up to math.MaxInt size.

	// We prefer returning usable, empty slice to nil.
	if qty < 1 {
		return &[]int8{}, nil
	}

	var nums []int8

	for ix := range qty {
		v := ix + 1
		nums = append(nums, v)
	}

	return &nums, nil
}

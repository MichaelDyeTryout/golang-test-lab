package sequence

func CountingNumbers(qty int8) (*[]int8, error) {
	// Return sequence of counting ints with size equal to qty up to math.MaxInt size.

	// We prefer returning usable, empty slice to nil.
	if qty < 1 {
		return &[]int8{}, nil
	}

	nums := make([]int8, qty)

	for ix := range qty {
		v := ix + 1
		nums[ix] = v
	}

	return &nums, nil
}

func Repeat(qty int8, val int) (*[]int, error) {
	// Return a slice of qty numbers, all with given val.

	// We prefer returning usable, empty slice to nil.
	if qty < 1 {
		return &[]int{}, nil
	}

	nums := make([]int, qty)

	for ix := range qty {
		nums[ix] = val
	}

	return &nums, nil
}

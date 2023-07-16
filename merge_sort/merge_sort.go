package main

func merge(left []int, right []int) []int {
	res := make([]int, len(left)+len(right))

	i := 0
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			res[i] = left[0]
			left = left[1:]
		} else {
			res[i] = right[0]
			right = right[1:]
		}
		i++
	}

	// These two for blocks are needed in case the left and right arrays are not
	// equal in length

	// I think we could optimize this because we know that the difference in
	// length between the two arrays is at most 1

	// But I'm being a bit lazy
	for j := 0; j < len(left); j++ {
		res[i] = left[j]
		i++
	}
	for j := 0; j < len(right); j++ {
		res[i] = right[j]
		i++
	}

	return res
}

func merge_sort(data []int) []int {
	data_len := len(data)

	if data_len <= 1 {
		return data
	}

	pivot := data_len / 2

	left := data[:pivot]
	right := data[pivot:]

	return merge(merge_sort(left), merge_sort(right))
}

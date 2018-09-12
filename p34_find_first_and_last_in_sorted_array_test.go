package leetcode

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	low, high := 0, len(nums)-1
	index := -1
	for low <= high {
		piv := (low + high) / 2
		if nums[piv] == target {
			index = piv
			break
		} else if target > nums[piv] {
			low = piv + 1
		} else {
			high = piv - 1
		}
	}

	if index == -1 {
		return []int{-1, -1}
	}

	// expand and get the range
	first, last := index, index
	for first > 0 && nums[first-1] == target {
		first--
	}

	for last < len(nums)-1 && nums[last+1] == target {
		last++
	}

	return []int{first, last}

}

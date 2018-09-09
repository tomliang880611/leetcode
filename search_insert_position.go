package leetcode

func searchInsert(nums []int, target int) int {

	if nums == nil || len(nums) == 0 {
		return -1
	}

	low, high := 0, len(nums)-1

	for high-low > 1 {
		lowVal, highVal := nums[low], nums[high]
		if target <= lowVal {
			return low
		}

		if target == highVal {
			return high
		}
		if target >= highVal {
			return high + 1
		}

		piv := (low + high) / 2
		if target == nums[piv] {
			return piv
		} else if target < nums[piv] {
			high = piv
		} else {
			low = piv
		}
	}
	if target == nums[high] {
		return high
	}
	if target == nums[low] {
		return low
	}
	if target > nums[high] {
		return high + 1
	}
	if target < nums[low] {
		return low
	}
	if nums[low] < target && target < nums[high] {
		return high
	}

	return -1
}

package leetcode

// this ignores the order in the resulting array
func removeElement(nums []int, val int) int {

	if nums == nil || len(nums) == 0 {
		return 0
	}

	low, high := 0, len(nums)-1
	for low < high {
		v := nums[low]
		if v == val {
			for ; low < high; high-- {
				// find the tailing non-val and swap
				if nums[high] != val {
					nums[low], nums[high] = nums[high], nums[low]
					high--
					low++
					break
				}
			}

		} else {
			low++
		}
	}

	return low + 1
}

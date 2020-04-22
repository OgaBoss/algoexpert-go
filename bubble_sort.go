package bubble_sort

func BubbleSort(nums []int) []int {
	sorted := false
	counter := 1
	for !sorted {
		sorted = true
		for i := 0; i < len(nums)-counter; i++ {
			if nums[i] > nums[i+1] {
				nums[i], nums[i+1] = nums[i+1], nums[i]
				sorted = false
			}
		}
		counter++
	}

	return nums
}

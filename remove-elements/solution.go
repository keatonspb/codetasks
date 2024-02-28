package remove_elements

func removeElement(nums []int, val int) int {
	var i int
	end := len(nums) - 1
	for i < len(nums) && i <= end {
		if nums[i] == val {
			//swap with last
			nums[i] = nums[end]
			end--
		} else {
			i++
		}
	}
	return i
}

package megre_sorted_arrays

// merge merges two sorted arrays into the first array
// nums1 has enough space to hold all elements from nums1 and nums2
// last element of nums1 should be max from all arrays
// so we can start from the end of both arrays and compare elements
// and put the max element to the end of nums1
// if nums2 has elements left we should put them to of nums1
// if nums1 has elements left we don't need to do anything
// time complexity is O(m+n)
// space complexity is O(1)
func merge(nums1 []int, m int, nums2 []int, n int) {
	i1 := m - 1
	i2 := n - 1

	c := m + n - 1 // this is cursor in nums1 to move elements from the end

	for i1 >= 0 && i2 >= 0 {
		// if last of first array is bigger than last of second array
		// this is biggest element from both arrays
		// we put it to the end of nums1
		if nums1[i1] > nums2[i2] {
			nums1[c] = nums1[i1]
			i1-- // first array element is moved, so we move cursor to the left
		} else {
			nums1[c] = nums2[i2]
			i2--
		}
		c--
	}

	// if second array has elements it means that all elements from first array are moved
	// and we need to move all elements from second array to the beginning of nums1 because
	// this elements are less than all elements from first array
	for i2 >= 0 {
		nums1[c] = nums2[i2]
		i2--
		c--
	}
}

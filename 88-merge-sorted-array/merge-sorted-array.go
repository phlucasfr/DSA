func merge(nums1 []int, m int, nums2 []int, n int)  {
    firstPointer := m - 1
	secondPointer := n - 1
	mergeIndex := m + n - 1

	for firstPointer >= 0 && secondPointer >= 0 {
		if nums1[firstPointer] > nums2[secondPointer] {
			nums1[mergeIndex] = nums1[firstPointer]
			firstPointer--
		} else {
			nums1[mergeIndex] = nums2[secondPointer]
			secondPointer--
		}
		mergeIndex--
	}

	for secondPointer >= 0 {
		nums1[mergeIndex] = nums2[secondPointer]
		secondPointer--
		mergeIndex--
	}
}
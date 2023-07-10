package sorts

// 快速排序,申请一份临时空间
func MergeSort(nums []int) []int {
	var temp []int = make([]int, len(nums)) // 申请一份临时数组
	mergeSort(nums, 0, len(nums)-1, temp)
	return nums
}

func mergeSort(nums []int, start int, end int, temp []int) {
	// 递归终止分解条件,只有1个元素则不在分解
	if start >= end {
		return
	}
	mid := start + (end-start)/2      // 取数组的start~end中间位置
	mergeSort(nums, start, mid, temp) // 拆分数组,start~mid之间
	mergeSort(nums, mid+1, end, temp) // 拆分数组,mid+1~end之间
	merge(nums, start, mid, end, temp)
}

// 合并两个有序数组
func merge(nums []int, start int, mid int, end int, temp []int) {
	var i = start
	var j = mid + 1
	var k = 0 // 临时数组的下标
	for i <= mid && j <= end {
		if nums[i] <= nums[j] { // 要用小于等于,保证稳定排序
			temp[k] = nums[i]
			i++
		} else {
			temp[k] = nums[j]
			j++
		}
		k++
	}
	for ; i <= mid; i++ { // start ~ end ,拷贝剩余部分
		temp[k] = nums[i]
		k++
	}
	for ; j <= end; j++ { // mid+1~end，拷贝剩余部分
		temp[k] = nums[j]
		k++
	}
	// 将合并后的结果拷贝到原数组中
	k = 0
	for start <= end {
		nums[start] = temp[k]
		k++
		start++
	}
}

// 归并排序,每次分解后的合并都申请一份临时空间
func MergeSortV2(nums []int) []int {

	return nums
}

// 合并两个有序数组
func mergeV2(a []int, b []int) []int {
	var result []int = make([]int, len(a)+len(b))

	return result
}

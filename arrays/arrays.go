package arrays

import (
	"fmt"
	"math"
	"sort"
)

func TwoSum(numbers []int, target int) []int {
	return twoSum(numbers, target)
}
func twoSum(numbers []int, target int) []int {
	var first, second = 0, 0
	var m = map[int]int{} // 定一个map并做空初始化
	for i, x := range numbers {
		k, ok := m[target-x]
		if ok { // 存在target-x,说明已经找到,k=上一个数的下标,i为当前数的下标
			first = k
			second = i
			break
		} else { // 不存在则先记录当前元素的下标,并放入map
			m[x] = i
		}

	}
	var ans = []int{first, second} // 定义2个元素的切片
	return ans
}

// Search 典型的二分查找
func Search(nums []int, target int) int {
	return search(nums, target)
}

func search(nums []int, target int) int {
	start, end := 0, len(nums)-1
	for start <= end {
		mid := start + (end-start)>>2 // start,end的中间位置
		if nums[mid] < target {       // 中间件值小于target,说明target在右边
			start = mid + 1
		} else if nums[mid] > target { // 中间值大于target,说明target在左边
			end = mid - 1
		} else {
			return mid
		}
	}
	return -1
}

func RemoveElement(nums []int, val int) int {
	// 用双指针,一个快指针,一个慢指针
	// 慢指针（slow)：在原数组中,始终指向的位置始终不包含val，快指针(fast)：遍历数组时,指向每一个元素的位置
	// slow指针左侧所形成的子数组范围里则都是不等于val的数
	slow, fast := 0, 0
	for fast < len(nums) {
		if nums[fast] != val { // 查找val,如果不等于val,将fast下标位置的元素放入到slow位置,然后slow到下一个位置
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	fmt.Printf("nums=%v,slow=%v\n", nums, slow)
	return slow // 由于范围长度,slow即子数组的长度,从0~slow-1都是不包含val的数
}

func SortedSquares(nums []int) []int {
	var newNums = make([]int, len(nums)) // 创建切片
	for i, n := range nums {
		newNums[i] = n * n // n^2
	}
	sort.Ints(newNums)
	fmt.Printf("Sorted: %v\n", newNums)
	return newNums
}

func SortedSquares2(nums []int) []int {
	length := len(nums)
	var newNums = make([]int, length) // 创建切片
	var k = length - 1
	for left, right := 0, length-1; left <= right; {
		x := nums[left] * nums[left]
		y := nums[right] * nums[right]
		if x > y {
			newNums[k] = x
			left++
		} else {
			newNums[k] = y
			right--
		}
		k--
	}
	fmt.Printf("Sorted: %v\n", newNums)
	return newNums
}

func MinSubArrayLen(target int, nums []int) int {
	return minSubArrayLen2(target, nums)
}

func minSubArrayLen(target int, nums []int) int {
	// 定义一个变量用来表示最小数组长度
	var minLen = math.MaxInt32
	// 定义前后指针front,back,初始化位置相同
	// front走在前面，back走在后面, 查找back与front区间满足>=target,记录区间大小,并用最小值滚动更新
	for back := 0; back < len(nums); back++ {
		sum := 0
		for front := back; front < len(nums); front++ {
			sum += nums[front]
			if sum >= target {
				minLen = minInt(minLen, front-back+1)
				break // 退出当前循环
			}
		}
	}
	// 表示没找到,没有进入sum>=target
	if minLen == math.MaxInt32 {
		return 0
	}
	return minLen
}

func minInt(x int, y int) int {
	if x <= y {
		return x
	}
	return y
}

func maxInt(x int, y int) int {
	if x <= y {
		return y
	}
	return x
}

func minSubArrayLen2(target int, nums []int) int {
	// 定义一个变量用来表示最小数组长度
	var subLen = math.MaxInt32
	// 定义start,end双指针通过组成一个滑动窗口
	var sum = 0 // 滑动窗口累加和
	for start, end := 0, 0; end < len(nums); end++ {
		sum += nums[end]
		for sum >= target && start <= end { // 不断缩小start,end窗口
			subLen = minInt(subLen, end-start+1) // 先记录满足条件的窗口大小
			sum -= nums[start]                   // 开始减start位置对应的数,剩余的累加和是否满足>=target
			start++
		}

	}
	// 表示没找到,没有进入sum>=target
	if subLen == math.MaxInt32 {
		return 0
	}
	return subLen
}

// GenerateMatrix 生成矩阵
func GenerateMatrix(n int) [][]int {
	// 创建大小为n的二维数组(切片）,golang需要显示初始化二维的对象
	var matrix = make([][]int, n) // 初始化第一维
	for k := 0; k < n; k++ {
		matrix[k] = make([]int, n)
	}
	var loop = n / 2 // 总共有几环,几层,矩阵大小为n,则只有n/2层
	var startx, starty = 0, 0
	var offset = 1 // 填充每一条边的偏移位置,每一层都少处理1个.第一层少处理1个,第二个少处理2个
	var count = 1  // 计数器,循环完成后计数到n平方
	for loop > 0 {
		i, j := startx, starty // 每层的起始位置
		// 左上到右上,i不变,执行完成后, j到达n-offset
		for j = starty; j < n-offset; j++ {
			matrix[startx][j] = count
			count++
		}
		// 从右上到右下,i增加,j不变,执行完成后,i到达n-offset
		for i = startx; i < n-offset; i++ {
			matrix[i][j] = count
			count++
		}
		// 从右下到左下,i不变,j减小,执行完成后,j到达starty
		for ; j > starty; j-- {
			matrix[i][j] = count
			count++
		}

		// 从左下到左上,i减小,j不变,执行完成后,i,j到达startx,starty
		for ; i > startx; i-- {
			matrix[i][j] = count
			count++
		}

		// 第一层填充完成，继续下一层,起始位置startx++,starty++,offset++
		startx++
		starty++
		offset++
		loop--
	}
	// 如果n是奇数,中间有个空位置,需要单独填充
	var center = n / 2
	if mod := n % 2; mod == 1 {
		matrix[center][center] = n * n
	}
	// fmt.Printf("fill matrix=%v\n", matrix)
	return matrix
}

func ThreeSum(nums []int) [][]int {
	return threeSum(nums)
}

// 查找 x+y+z=0的三个重复的数
func threeSum(nums []int) [][]int {
	// 创建存放结果的切片
	var result = make([][]int, 0) // 指定大小为1,需要初始化对应的二维
	// 边界值处理
	if nums == nil || len(nums) < 3 {
		return result
	}
	// 核心算法：排序 + 双指针
	// 1.排序,负数排在前,相同数排在一起,正数排在后
	sort.Ints(nums)
	// 2.遍历数组,锚定每一个x=nums[i],然后通过双指针向后查找y,z
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 { // x大于0,说明后面不可能存在x+y+z等于0
			return result
		}
		// 重复检查
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		// 双指针：一左一右,查找i+1~n-1直接的数
		// l为左边的,r=右边的,当l<r,依次查找y=nums[l],z=nums[r],计算s = x+y+z
		// 如果s=0,说明x,y,z满足条件,l++,r--,继续查找下一组
		// 如果s>0,说明z较大,r--左移
		// 如果s<0,说明y较小,l++右移
		l, r := i+1, len(nums)-1
		for l < r {
			s := nums[i] + nums[l] + nums[r]
			if s == 0 { // 左右两边如果存在重复都取遇到的第一个
				// 找到满足条件的三元组,放入结果中
				result = append(result, []int{nums[i], nums[l], nums[r]})
				// 找到答案后,继续移动,但是需要检查左右两边是否存在重复
				for l < r && nums[l] == nums[l+1] { // 左边存在重复,跳过
					l++
				}
				for l < r && nums[r] == nums[r-1] { // 右边存在重复,跳过
					r--
				}
				l++
				r--
			} else if s > 0 {
				r--
			} else {
				l++
			}
		}
	}
	return result
}

func NumSubarrayProductLessThanK(nums []int, k int) int {
	return numSubarrayProductLessThanK(nums, k)
}
func numSubarrayProductLessThanK(nums []int, k int) int {
	count := 0
	product := 1
	for i, j := 0, 0; j < len(nums); j++ { // 遍历每一个j,然后做乘积积累
		product *= nums[j]
		for i <= j && product >= k { // 如果大于k
			product /= nums[i]
			i++
		}
		v := j - i + 1
		count += v // 这行是关键
	}
	return count
}

// FindMaxLength 题意：找到含有相同数量的 0 和 1 的最长连续子数组
func FindMaxLength(nums []int) int {
	// 第一步：把原数组先做转换,将0转换为-1.这样问题就转换为: 寻找子数组和为0的最长子数组
	var newNums = make([]int, len(nums))
	for i, v := range nums {
		if v == 0 {
			newNums[i] = -1
		} else {
			newNums[i] = v
		}
	}
	// 第二步：计算前缀和，定义一个map,key为:下标i的前缀和,value为下标i
	var prefixSumMap = map[int]int{0: -1} // 初始化一个0:-1,表示有一个虚拟的前缀和为0,下标位置为-1
	var prefixSum = 0                     // 下标i的前缀和
	k := 0
	maxLen := 0
	for i := 0; i < len(newNums); i++ {
		// 检查是否存在一个prefixSum[j] + k = prefixSum(prefixSum[i]) 特殊处理一些边界值
		prefixSum += newNums[i]
		j, ok := prefixSumMap[prefixSum-k]
		if ok { // 如果存在prefixSum[j],其下标等于j
			maxLen = maxInt(i-j, maxLen)
		} else { // 如果不存在,则把当前i的前缀和放入map
			prefixSumMap[prefixSum] = i
		}
	}
	return maxLen
}

func SubarraySum(nums []int, k int) int {
	var count = 0
	var pre = 0                    // 前缀和
	var preMap = map[int]int{0: 1} // map的key为前j个数的前缀和,value为前缀和次数
	for i := 0; i < len(nums); i++ {
		pre += nums[i]  // pre=前i个数的前缀和
		preJ := pre - k // 定义为前j个数的前缀和
		// 查看map里是否存在一个前缀和preJ
		// 如果存在说明j到i之间满足:preJ + k = pre,说明存在一个连续子数组和=k
		// 如果不存在,则把pre放入map,value为1,表示该pre出现过1次
		if v, ok := preMap[preJ]; ok {
			count += v
		}
		preMap[pre] += 1
	}
	return count
}

// 给定一个整数数组和一个整数 k ，请找到该数组中和为 k 的连续子数组的个数
func SubarraySum2(nums []int, k int) int {
	var count = 0
	for i := 0; i < len(nums); i++ {
		sum := 0
		// 在start-end区间里查找累加和=k
		for j := i; j >= 0; j-- {
			sum += nums[j]
			if sum == k {
				count++
			}
		}
	}
	return count
}

func RunningSum(nums []int) []int {
	return runningSum1(nums)
}

// 暴力解法
func runningSum1(nums []int) []int {
	runningSum := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		sum := 0
		for j := i; j >= 0; j-- {
			sum += nums[j]
		}
		runningSum[i] = sum
	}
	return runningSum
}

func RunningSum2(nums []int) []int {
	return runningSum2(nums)
}

// 前缀和
func runningSum2(nums []int) []int {
	runningSum := make([]int, len(nums))
	prefixSum := 0
	for i := 0; i < len(nums); i++ {
		iSum := prefixSum + nums[i]
		runningSum[i] = iSum
		prefixSum = iSum
	}
	return runningSum
}

type NumArray struct {
	prefixSum []int
}

func Constructor(nums []int) NumArray {
	prefixSum := make([]int, len(nums)) // 局部变量小写
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			prefixSum[i] = nums[i]
		} else {
			prefixSum[i] = prefixSum[i-1] + nums[i]
		}
	}
	return NumArray{prefixSum}
}

// SumRange 求[left,right]的数字之后,包含left,即求prefixSum[right]-prefix[left-1]这个区间的数字之和
func (this *NumArray) SumRange(left int, right int) int {
	// 边界处理
	if left == 0 {
		return this.prefixSum[right]
	}
	return this.prefixSum[right] - this.prefixSum[left-1]
}

func FindMiddleIndex(nums []int) int {
	return findMiddleIndex(nums)
}

func findMiddleIndex(nums []int) int {
	// 1. 先计算前缀和
	prefixSum := make([]int, len(nums)) // 局部变量小写
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			prefixSum[i] = nums[i]
		} else {
			prefixSum[i] = prefixSum[i-1] + nums[i]
		}
	}
	// 得到的前缀和示例
	// 0,1,2,.....maxIndex
	middleIndex := -1
	minIndex := 0
	maxIndex := len(nums) - 1
	// 2.边界检查,检查middleIndex是否为最左边或最右边
	// 如果prefixSum[maxIndex]-prefix[0]=0，说明在最左边,middleIndex=0
	if prefixSum[maxIndex]-prefixSum[0] == 0 {
		middleIndex = minIndex
		return middleIndex
	}
	// 如果prefixSum[maxIndex-1]=0说明middleIndex在最右边,middleIndex=maxIndex
	if prefixSum[maxIndex-1] == 0 {
		middleIndex = maxIndex
	}
	// 3.依次检查是否存在一个k=middleIndex满足：prefixSum[middleIndex-1]=prefixSum[maxIndex]-prefixSum[middleIndex]
	for k := 1; k < maxIndex; k++ {
		if prefixSum[k-1] == prefixSum[maxIndex]-prefixSum[k] {
			middleIndex = k
			break
		}
	}
	if middleIndex == -1 {
		return middleIndex
	}
	if middleIndex == maxIndex {
		return middleIndex
	}
	return middleIndex
}

func NumberOfSubarrays(nums []int, k int) int {
	// 1.转换数组,将奇数转换为1，偶数转换为0
	newNums := make([]int, len(nums))
	for i, v := range nums {
		if v%2 == 0 {
			newNums[i] = 0
		} else {
			newNums[i] = 1
		}
	}

	// 2.计算前缀和,并统计子数组数量
	count := 0
	prefixSum := 0                    // prefixSum始终为prefixSum[i]的前缀和
	prefixSumMap := map[int]int{0: 1} // 定一个前缀和计数器,初始化一个空的前缀和,其值为0,数量为1
	for i := 0; i < len(newNums); i++ {
		prefixSum += newNums[i]            // prefixSum 即 prefixSum[i]
		v, ok := prefixSumMap[prefixSum-k] // 检查是否存在一个prefixSum[j]满足：prefixSum[j] + k = prefixSum[i]
		if ok {                            // 如果存在,value为该前缀和的数量
			count += v
		}
		prefixSumMap[prefixSum] += 1
	}
	return count
}

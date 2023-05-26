package hashs

import (
	"sort"
)

func IsAnagram(s string, t string) bool {
	return isAnagram2(s, t)
}

func isAnagram1(s string, t string) bool {
	// 边界
	if len(s) != len(t) || s == "" || t == "" {
		return false
	}
	// 把字符串s的所有字符都放入哈希表中，value为字符的次数
	charMap := make(map[string]int, len(s))
	for _, r := range s {
		cnt, ok := charMap[string(r)]
		if ok {
			cnt++
		} else {
			cnt = 0
		}
		charMap[string(r)] = cnt
	}
	for _, r := range t {
		cnt, ok := charMap[string(r)]
		if !ok {
			return false
		}
		if cnt < 0 {
			return false
		}
		cnt--
		charMap[string(r)] = cnt
	}
	return true
}

// 用数组来模拟哈希表，因为题目中只有小写字母
// 时间复杂度：O(s) + O(t), 空间复杂度：O(1)
func isAnagram2(s string, t string) bool {
	// 边界
	if s == "" || t == "" || len(s) != len(t) {
		return false
	}
	// 把字符串s的所有字符都放入数组中，value为字符的次数
	charArray := make([]rune, 26)
	for _, r := range s {
		charArray[r-'a']++
	}
	for _, r := range t {
		charArray[r-'a']--
	}
	// 再次检查字符数组中各个字符的计数,如果小于0说明两个字符串不是变位词
	for _, ch := range charArray {
		if ch < 0 {
			return false
		}
	}
	return true
}

func Intersection(nums1 []int, nums2 []int) []int {
	numMap := make(map[int]bool, len(nums1))
	var distinctMap = make(map[int]int)
	for _, n := range nums1 {
		numMap[n] = true
	}
	for _, n := range nums2 {
		v, ok := numMap[n]
		if ok && v {
			distinctMap[n] = 1
		}
	}
	arr := make([]int, 0)
	for k, _ := range distinctMap {
		arr = append(arr, k)
	}
	return arr
}

func IsHappy(n int) bool {
	sumMap := make(map[int]bool, 0)
	i := findSum(n)
	for i != 1 {
		v := sumMap[i]
		if v {
			return false
		}
		sumMap[i] = true
		i = findSum(i)
	}
	return true
}

// 计算每个位置的平方sum
// 找到n的余数：68 = 60 + 8
func findSum(n int) int {
	sum := 0
	for n > 0 {
		t := n % 10
		sum += t * t
		n /= 10
	}
	return sum
}

func FourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	return fourSumCount(nums1, nums2, nums3, nums4)
}

// 用哈希表来优化常规的暴力查找,空间换时间
func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	var k = 0
	var fourSumMap = map[int]int{}
	for _, x := range nums1 {
		for _, y := range nums2 {
			twoSum := x + y
			v, ok := fourSumMap[twoSum]
			if ok {
				v++
			} else {
				v = 1
			}
			fourSumMap[twoSum] = v
		}
	}
	count := 0
	// 在map中查找两数之和=k-twoSum
	for _, x := range nums3 {
		for _, y := range nums4 {
			twoSum := x + y
			v, ok := fourSumMap[k-twoSum]
			if ok {
				count += v
			}
		}
	}
	return count
}

func CanConstruct(ransomNote string, magazine string) bool {
	magazinemap := make(map[rune]bool, len(magazine))
	for _, r := range magazine {
		magazinemap[r] = true
	}
	ransommap := make(map[rune]bool, len(ransomNote))
	for _, r := range ransomNote {
		ransommap[r] = true
	}
	if len(ransomNote) >= len(magazine) {
		for k := range magazinemap {
			if !ransommap[k] {
				return false
			}
			delete(magazinemap, k)
		}
	} else {
		for k := range ransommap {
			if !magazinemap[k] {
				return false
			}
			delete(magazinemap, k)
		}
	}

	return true
}

func ThreeSum(nums []int) [][]int {
	result := make([][]int, 0)
	if len(nums) == 0 {
		return result
	}
	// 初始化result
	// 通过排序+双指针来查找三数之和
	// 第一步：排序
	sort.Ints(nums)
	// 第二步：遍历每一个数组元素i,x=nums[i], 然后从left=i+1~num[len-1]之间查找2个数满足: x+y+z=0
	size := len(nums)
	for i := 0; i < size; i++ {
		// 排序后,如果当前x>0,后面不可能出现x+y+z=0
		if nums[i] > 0 {
			return result
		}
		// x去重,如果前一个与当前x相同,前一个x如果满足,则当前x页满足
		// 整体去重逻辑,如果遇到的第一个数满足则加入候选,下一个重复则跳过
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}
		x := nums[i]
		left := i + 1
		right := len(nums) - 1
		// 因为找2个数,left<right即可
		for left < right {
			y, z := nums[left], nums[right]
			if x+y+z == 0 { // x+y+z=0,说明找到了
				result = append(result, []int{x, y, z})
				// y去重检查,当前与下一个相同,跳过
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				// z去重检查,当前与下一个相同,跳过
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if x+y+z > 0 { // x+y+z>0,说明z偏大,z左移
				right--
			} else { // x+y+z<0,说明y偏小,y右移动
				left++
			}
		}
	}
	return result
}

func FourSum(nums []int, target int) [][]int {
	result := make([][]int, 0)
	if len(nums) == 0 {
		return result
	}
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		// 第一层剪枝优化
		//if nums[i] > target && nums[i] >= 0 {
		//	break
		//}
		// a去重
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		a := nums[i]
		for j := i + 1; j < len(nums); j++ {
			// 第二层剪枝优化
			//if nums[i]+nums[j] > target && nums[i]+nums[j] >= 0 {
			//	break
			//}
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			b := nums[j]
			left, right := j+1, len(nums)-1
			for left < right {
				c, d := nums[left], nums[right]
				if a+b+c+d == target { // 说明找到了
					result = append(result, []int{a, b, c, d})
					// 去重检查,当前与下一个相同,跳过
					for left < right && nums[left] == nums[left+1] {
						left++
					}
					for left < right && nums[right] == nums[right-1] {
						right--
					}
					left++
					right--
				} else if a+b+c+d > target { // 说明z偏大,z左移
					right--
				} else { // 说明y偏小,y右移动
					left++
				}
			}
		}

	}
	return result
}

// 判断 s2 是否包含 s1 的某个变位词。
func CheckInclusion(s1 string, s2 string) bool {
	// 边界
	if len(s1) > len(s2) {
		return false
	}
	// 创建一个简单的数组用来模拟哈希表
	var charmap = [26]int{}
	// 基于s1的长度遍历字符串
	for i := 0; i < len(s1); i++ {
		charmap[s1[i]-'a']++
		charmap[s2[i]-'a']--
	}
	if allZero(charmap) {
		return true
	}
	// 基于s1的长度n定位一个滑动窗口
	// right为下一个滑入的字符,left为下一个滑出的字符
	for right := len(s1); right < len(s2); right++ {
		left := right - len(s1)
		charmap[s2[right]-'a']--
		charmap[s2[left]-'a']++
		if allZero(charmap) {
			return true
		}
	}
	return false
}

func allZero(charmap [26]int) bool {
	return charmap == [26]int{} // 直接判断数组为0
}

// 用2个数组分别来计数s1,s2的字符个数，如果最后2个数组字符个数相同，则认为存在变位词
func CheckInclusion2(s1 string, s2 string) bool {
	// 边界
	if len(s1) > len(s2) {
		return false
	}
	charCnt1 := [26]int{}
	charCnt2 := [26]int{}
	// 基于s1的长度遍历字符串
	for i := 0; i < len(s1); i++ {
		charCnt1[s1[i]-'a']++
		charCnt2[s2[i]-'a']++
	}
	if charCnt1 == charCnt2 {
		return true
	}
	// 构建一个窗口大小为s1的长度n,新进入一个字符,其计数++,出去一个字符其计数--
	n := len(s1)
	for i := n; i < len(s2); i++ {
		charCnt2[s2[i]-'a']++
		charCnt2[s2[i-n]-'a']--
		if charCnt1 == charCnt2 {
			return true
		}
	}
	return false
}

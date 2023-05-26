package main

import (
	"fmt"
	"myalgo/alonestrings"
	"myalgo/arrays"
	"myalgo/hashs"
	"myalgo/integers"
)

func main() {
	fmt.Println("div =", integers.Divide(15, 2))
	fmt.Println("div =", integers.Divide(7, -3))
	fmt.Println("div =", integers.Divide(-10999, -3))
	fmt.Println("div =", integers.Divide(2147483647, 2))
	fmt.Println("addBinary = ", integers.AddBinary("11", "11"))
	fmt.Println("addBinary = ", integers.AddBinary("11", "10"))
	fmt.Println("SingleNumber = ", integers.SingleNumber([]int{2, 2, 3, 2}))
	fmt.Println("SingleNumber = ", integers.SingleNumber([]int{0, 1, 0, 1, 0, 1, 100}))
	fmt.Println("SingleNumber = ", integers.SingleNumber([]int{-2, -2, 1, 1, 4, 1, 4, 4, -4, -2}))
	fmt.Println("TwoSum = ", arrays.TwoSum([]int{1, 2, 4, 6, 10}, 8))
	fmt.Println("TwoSum = ", arrays.TwoSum([]int{2, 3, 4}, 6))
	fmt.Println("Search = ", arrays.Search([]int{-1, 0, 3, 5, 9, 12}, 2))
	fmt.Println("Search = ", arrays.Search([]int{-1, 0, 3, 5, 9, 12}, 9))
	fmt.Println("RemoveElement = ", arrays.RemoveElement([]int{3, 2, 2, 3}, 3))
	fmt.Println("RemoveElement = ", arrays.RemoveElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2))
	fmt.Println("SortedSquares = ", arrays.SortedSquares([]int{-4, -1, 0, 3, 10}))
	fmt.Println("SortedSquares = ", arrays.SortedSquares2([]int{-4, -1, 0, 3, 10}))
	fmt.Println("SortedSquares = ", arrays.SortedSquares2([]int{7, -3, 2, 3, 11}))
	fmt.Println("MinSubArrayLen = ", arrays.MinSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}))
	fmt.Println("MinSubArrayLen = ", arrays.MinSubArrayLen(4, []int{1, 4, 4}))
	fmt.Println("MinSubArrayLen = ", arrays.MinSubArrayLen(11, []int{1, 1, 1, 1, 1, 1, 1, 1}))
	fmt.Println("GenerateMatrix = ", arrays.GenerateMatrix(3))
	fmt.Println("GenerateMatrix = ", arrays.GenerateMatrix(4))
	fmt.Println("GenerateMatrix = ", arrays.GenerateMatrix(1))
	fmt.Println("ThreeSum = ", arrays.ThreeSum([]int{-1, 0, 1, 2, -1}))
	fmt.Println("ThreeSum = ", arrays.ThreeSum([]int{0, 1, 1}))
	fmt.Println("ThreeSum = ", arrays.ThreeSum([]int{0, 0, 0}))
	fmt.Println("NumSubarrayProductLessThanK = ", arrays.NumSubarrayProductLessThanK([]int{10, 5, 2, 6}, 100))
	fmt.Println("SubarraySum2 = ", arrays.SubarraySum2([]int{1, 1, 1, 1}, 2))
	fmt.Println("SubarraySum = ", arrays.SubarraySum([]int{1, 1, 1, 1}, 2))
	fmt.Println("RunningSum = ", arrays.RunningSum([]int{1, 1, 1, 1}))
	fmt.Println("RunningSum = ", arrays.RunningSum([]int{1, 2, 3, 4}))
	fmt.Println("RunningSum2 = ", arrays.RunningSum2([]int{1, 2, 3, 4}))
	var a = arrays.Constructor([]int{-2, 0, 3, -5, 2, -1})
	fmt.Printf("constructor=%v\n", a)
	fmt.Printf("constructor.SumRange=%v\n", a.SumRange(0, 2))
	fmt.Printf("constructor.SumRange=%v\n", a.SumRange(2, 5))
	fmt.Printf("constructor.SumRange=%v\n", a.SumRange(0, 5))
	fmt.Printf("FindMiddleIndex=%v\n", arrays.FindMiddleIndex([]int{2, 3, -1, 8, 4}))
	fmt.Printf("FindMiddleIndex=%v\n", arrays.FindMiddleIndex([]int{1, -1, 4}))
	fmt.Printf("FindMaxLength=%v\n", arrays.FindMaxLength([]int{0, 1}))
	fmt.Printf("FindMaxLength=%v\n", arrays.FindMaxLength([]int{0, 1, 0}))
	fmt.Printf("NumberOfSubarrays=%v\n", arrays.NumberOfSubarrays([]int{1, 1, 2, 1, 1}, 3))
	fmt.Printf("NumberOfSubarrays=%v\n", arrays.NumberOfSubarrays([]int{2, 4, 6}, 1))
	fmt.Printf("IsAnagram=%v\n", hashs.IsAnagram("anagram", "nagaram"))
	fmt.Printf("IsAnagram=%v\n", hashs.IsAnagram("ab", "a"))
	fmt.Printf("Intersection=%v\n", hashs.Intersection([]int{1, 2, 2, 1}, []int{2, 2}))
	//fmt.Printf("IsHappy=%v\n", hashs.IsHappy(19))
	fmt.Printf("IsHappy=%v\n", hashs.IsHappy(2))
	fmt.Printf("CanConstruct=%v\n", hashs.CanConstruct("aa", "aab"))
	fmt.Printf("CanConstruct=%v\n", hashs.CanConstruct("aa", "ab"))

	fmt.Printf("FourSumCount=%v\n",
		hashs.FourSumCount(
			[]int{1, 2},
			[]int{-2, -1},
			[]int{-1, 2},
			[]int{0, 2},
		))
	fmt.Printf("ThreeSum=%v\n", hashs.ThreeSum([]int{-1, 0, 1, 2, -1, -4}))
	fmt.Printf("FourSum=%v\n", hashs.FourSum([]int{1, 0, -1, 0, -2, 2}, 0))
	fmt.Printf("CheckInclusion=%v\n", hashs.CheckInclusion("ab", "eidbaooo"))
	fmt.Printf("CheckInclusion=%v\n", hashs.CheckInclusion("ab", "baa"))
	fmt.Printf("CheckInclusion=%v\n", hashs.CheckInclusion2("ab", "baa"))
	fmt.Printf("IsPalindrome=%v\n", alonestrings.IsPalindrome("A man, a plan, a canal: Panama"))
	fmt.Printf("IsPalindrome=%v\n", alonestrings.IsPalindrome("race a car"))
	fmt.Printf("ReverseString=%v\n", alonestrings.ReverseString([]byte{'h', 'e', 'l', 'l', 'o'}))
	fmt.Printf("ReverseStr=%v\n", alonestrings.ReverseStr("abcdefg", 2))
}
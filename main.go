package main

import (
	"container/list"
	"fmt"
	"myalgo/alonestrings"
	"myalgo/arrays"
	"myalgo/binarytrees"
	"myalgo/hashs"
	"myalgo/heaps"
	"myalgo/integers"
	"myalgo/sorts"
	"myalgo/stackqueues"
)

func main() {
	// Create a new list and put some numbers in it.
	l := list.New()
	e4 := l.PushBack(4)
	e1 := l.PushFront(1)
	l.InsertBefore(3, e4)
	l.InsertAfter(2, e1)
	// Iterate through list and print its contents.
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

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
	fmt.Printf("ReplaceSpace=%v\n", alonestrings.ReplaceSpace("We are happy."))
	fmt.Printf("ReplaceSpace2=%v\n", alonestrings.ReplaceSpace2("We are happy."))
	// fmt.Printf("ReverseWords=%v\n", alonestrings.ReverseWords("We are happy."))
	fmt.Printf("ReverseWords=%v\n", alonestrings.ReverseWords("  hello   world  "))
	fmt.Printf("ReverseLeftWords=%v\n", alonestrings.ReverseLeftWords("abcdefg", 2))
	fmt.Printf("ReverseLeftWords=%v\n", alonestrings.ReverseLeftWords2("lrloseumgh", 6))
	fmt.Printf("ReverseLeftWords3=%v\n", alonestrings.ReverseLeftWords3("lrloseumgh", 6))
	//fmt.Printf("BuildKmpNext=%v\n", alonestrings.BuildKmpNext("abcd"))
	//fmt.Printf("StrStr_KMP=%v\n", alonestrings.StrStr_KMP("abcd", "ab"))
	//fmt.Printf("BuildKmpNext=%v\n", alonestrings.BuildKmpNext("aabaaf"))
	fmt.Printf("StrStr_KMP=%v\n", alonestrings.StrStr_KMP("aabaabaaf", "aabaaf"))
	fmt.Printf("StrStr_KMP=%v\n", alonestrings.StrStr_KMP("xyz", "z"))
	//fmt.Printf("BuildPrefixString=%v\n", alonestrings.BuildPrefixString("abcd"))
	//fmt.Printf("RepeatedSubstringPattern=%v\n", alonestrings.RepeatedSubstringPattern("abab"))
	//fmt.Printf("BuildKmpNext=%v\n", alonestrings.BuildKmpNext("aaaaaa"))
	fmt.Printf("String=%v,BuildKmpNext=%v\n", "abcabcabcabc", alonestrings.BuildKmpNext("abcabcabcabc"))
	fmt.Printf("String=%v,BuildKmpNext=%v\n", "abcabc", alonestrings.BuildKmpNext("abcabc"))
	fmt.Printf("BuildPrefixString=%v\n", alonestrings.BuildPrefixString("abab"))
	fmt.Printf("RepeatedSubstringPattern=%v\n", alonestrings.RepeatedSubstringPattern("abcabcabcabc"))
	fmt.Printf("IsValid=%v\n", stackqueues.IsValid("[{()}][]"))
	fmt.Printf("IsValid=%v\n", stackqueues.IsValid2("[{()}][]"))
	fmt.Printf("RemoveDuplicates=%v\n", stackqueues.RemoveDuplicates("abbaca"))
	fmt.Printf("EvalRPN=%v\n", stackqueues.EvalRPN([]string{"4", "13", "5", "/", "+"}))
	fmt.Printf("EvalRPN=%v\n", stackqueues.EvalRPN([]string{"4", "3", "-"}))
	fmt.Printf("EvalRPN=%v\n", stackqueues.EvalRPN([]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}))
	fmt.Printf("MaxSlidingWindow=%v\n", stackqueues.MaxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
	fmt.Printf("MaxSlidingWindow=%v\n", stackqueues.MaxSlidingWindow2([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
	fmt.Printf("MaxSlidingWindow=%v\n", stackqueues.MaxSlidingWindow([]int{7, 2, 4}, 2))
	fmt.Printf("MaxSlidingWindow=%v\n", stackqueues.MaxSlidingWindow2([]int{7, 2, 4}, 2))
	fmt.Printf("HeapDemo=%v\n", heaps.Demo())

	fmt.Printf("PreorderTraversal=%v\n", binarytrees.PreorderTraversal(createTree()))
	fmt.Printf("PreorderTraversalByLoop=%v\n", binarytrees.PreorderTraversalByLoop(createTree()))
	fmt.Printf("InorderTraversalByRecursive=%v\n", binarytrees.InorderTraversalByRecursive(createTree()))
	fmt.Printf("InorderTraversalByLoop=%v\n", binarytrees.InorderTraversalByLoop(createTree()))
	fmt.Printf("PostorderTraversalByRecursive=%v\n", binarytrees.PostorderTraversalByRecursive(createTree()))
	fmt.Printf("PostorderTraversalByLoop=%v\n", binarytrees.PostorderTraversalByLoop(createTree()))
	fmt.Printf("LevelOrder=%v\n", binarytrees.LevelOrder(createLevelTree()))
	fmt.Printf("LevelOrderByArray=%v\n", binarytrees.LevelOrderByArray(createLevelTree()))
	fmt.Printf("MaxDepth=%v\n", binarytrees.MaxDepth(createLevelTree()))
	fmt.Printf("LevelOrderBottom=%v\n", binarytrees.LevelOrderBottom(createLevelTree()))
	fmt.Printf("RightSideView=%v\n", binarytrees.RightSideView(createLevelTree()))
	fmt.Printf("AverageOfLevels=%v\n", binarytrees.AverageOfLevels(createLevelTree()))
	fmt.Printf("MinDepth=%v\n", binarytrees.MinDepth(createLevelTree()))
	fmt.Printf("InvertTree=%v\n", binarytrees.InvertTree(createRevertTree()))
	fmt.Printf("IsSymmetric=%v\n", binarytrees.IsSymmetric(createRevertTree()))
	fmt.Printf("IsSymmetricV3=%v\n", binarytrees.IsSymmetricV3(createRevertTree()))
	fmt.Printf("BinaryTreePathsV1=%v\n", binarytrees.BinaryTreePathsV1(createRevertTree()))
	fmt.Printf("BinaryTreePathsV1=%v\n", binarytrees.BinaryTreePathsV1(createLevelTree()))
	fmt.Printf("FindBottomLeftValueV1=%v\n", binarytrees.FindBottomLeftValueV1(createLevelTree()))
	fmt.Printf("HasPathSumV1=%v\n", binarytrees.HasPathSumV1(createLevelTree(), 30))
	fmt.Printf("PathSumV1=%v\n", binarytrees.PathSumV1(createLevelTree(), 30))
	fmt.Printf("PathSumV2=%v\n", binarytrees.PathSumV2(createLevelTree(), 30))
	fmt.Printf("SearchBST=%v\n", binarytrees.SearchBSTV1(createSearchTree(), 2))

	fmt.Printf("FindNum_Soluction1=%v\n", arrays.FindNum_Soluction1([]int{1, 6, 9, 2, 1, 5, 15, 20}))
	fmt.Printf("FindNum_Soluction2=%v\n", arrays.FindNum_Soluction2([]int{1, 6, 9, 2, 1, 5, 15, 20}))
	fmt.Printf("FindNumV2=%v\n", arrays.FindNumV2([]int{10, 8, 9, 7, 6, 5, 1, 3, 2}))
	fmt.Printf("MergeSort=%v\n", sorts.MergeSort([]int{8, 4, 5, 7, 1, 3, 6, 2}))
	fmt.Printf("MergeSort=%v\n", sorts.MergeSort([]int{10, 8, 9, 7, 6, 5, 1, 3, 2}))

}

func createTree() *binarytrees.TreeNode {
	c := &binarytrees.TreeNode{Val: 3}
	b := &binarytrees.TreeNode{Val: 2, Left: c, Right: nil}
	a := &binarytrees.TreeNode{Val: 1, Left: nil, Right: b}
	return a
}

func createLevelTree() *binarytrees.TreeNode {
	a := &binarytrees.TreeNode{Val: 10,
		Left: &binarytrees.TreeNode{Val: 9},
		Right: &binarytrees.TreeNode{Val: 20,
			Left:  &binarytrees.TreeNode{Val: 15},
			Right: &binarytrees.TreeNode{Val: 7}}}
	return a
}

func createRevertTree() *binarytrees.TreeNode {
	a := &binarytrees.TreeNode{Val: 1,
		Left:  &binarytrees.TreeNode{Val: 2},
		Right: &binarytrees.TreeNode{Val: 2}}
	return a
}

func createSearchTree() *binarytrees.TreeNode {
	a := &binarytrees.TreeNode{Val: 4,
		Left: &binarytrees.TreeNode{Val: 2,
			Left:  &binarytrees.TreeNode{Val: 1},
			Right: &binarytrees.TreeNode{Val: 3}},
		Right: &binarytrees.TreeNode{Val: 7},
	}
	return a
}

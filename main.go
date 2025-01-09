package main

import (
	"DSA/array"
	"fmt"
)

func main() {
	fmt.Println("Pivot Index: ", array.PivotIndex([]int{1, 7, 3, 6, 5, 6}))

	fmt.Println("Remove Element To Make Strictly Increasing: ", array.CanBeIncreasing([]int{1, 2, 10, 5, 7}))

	fmt.Println("Remove Element: ", array.RemoveElement([]int{3, 2, 2, 3}, 3))

	fmt.Println("Remove Duplicate Sorted Array: ", array.RemoveDuplicates([]int{1, 1, 2}))

	fmt.Println("Two Sum: ", array.TwoSum([]int{-1, -2, -3, -4, -5}, -8))

	fmt.Println("Concatenation array: ", array.GetConcatenation([]int{1, 3, 2, 1}))

	fmt.Println("Contains duplicate: ", array.ContainsDuplicate([]int{1, 2, 3}))

	fmt.Println("Is Anagram: ", array.IsAnagram("rat", "tar"))

	fmt.Println("Group Anagrams: ", array.GroupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))

	fmt.Println("Top K Frequent Elements: ", array.TopKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))

	ce := array.CustomEncoder{Delimiter: '#'}
	encodedStr := ce.Encode([]string{"String with new\nline", "Another\nLine", "And\nOne\nMore"})
	fmt.Println("Custom Encoder: ", encodedStr)
	decodedStr := ce.Decode(encodedStr)
	fmt.Println("Custom Decoder: ", decodedStr)

	fmt.Println("Product Array Except Self: ", array.ProductExceptSelf([]int{1, 2, 3, 4}))

	fmt.Println("Is Valid Sodoku: ", array.IsValidSudoku([][]byte{
		{'1', '2', '.', '.', '3', '.', '.', '.', '.'},
		{'4', '.', '.', '5', '.', '.', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '.', '3'},
		{'5', '.', '.', '.', '6', '.', '.', '.', '4'},
		{'.', '.', '.', '8', '.', '3', '.', '.', '5'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '.', '.', '.', '.', '.', '2', '.', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '8'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'}}))

	fmt.Println("Count Vowel Strings in Ranges: ", array.VowelStrings([]string{"aba", "bcb", "ece", "aa", "e"}, [][]int{{0, 2}, {1, 4}, {1, 1}}))

	fmt.Println("Number of Ways to Split Array: ", array.WaysToSplitArray([]int{10, 4, -8, 7}))

	fmt.Println("Unique Length-3 Palindromic Subsequences: ", array.CountPalindromicSubsequence("aabca"))

	fmt.Println("Shifting Letters II: ", array.ShiftingLetters("iaozjb",
		[][]int{{0, 4, 0}, {0, 2, 1}, {1, 3, 1}, {0, 4, 1}, {4, 4, 1}, {2, 3, 0}, {5, 5, 0}, {3, 3, 0}, {2, 3, 0}, {5, 5, 1}, {5, 5, 1}, {5, 5, 0}}))

	fmt.Println("Count Prefix and Suffix Pairs I: ", array.CountPrefixSuffixPairs([]string{"a", "aba", "ababa", "aa"}))
}

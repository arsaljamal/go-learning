package main

import "fmt"


func lengthOfLongestSubString(s string) int {
	// Map to store the last seen index of each character.
	seen := make(map[byte]int)
	// left pointer of the sliding window.
	left := 0
	//max length of the substring found so far.
	maxLength := 0

	// Iterate over the string with the right pointer.
	for right := 0; right < len(s); right++ {
		// If the character has been seen and is within the current window.
		if idx, ok := seen[s[right]]; ok && idx >= left {
			// Move the left pointer to one position right of the last seen index.
			left = idx + 1
		}
		// Update the last seen index of the character.
		seen[s[right]] = right
		// Calculate the current window size and update maxLength if it's larger.
		if right-left+1 > maxLength {
			maxLength = right -left + 1
		}
	}
	return maxLength
	//Imagine walking through a string with two fingers (left and right).
	// You move right to expand the window.
	// If you see a duplicate, slide left past the last occurrence
	// of that duplicate. The largest window you see = answer.
}

func main() {
	fmt.Println(lengthOfLongestSubString("abcabcbb")) // Output: 3
	fmt.Println(lengthOfLongestSubString("bbbbb"))   // Output: 1
	fmt.Println(lengthOfLongestSubString("pwwkew"))   // Output: 3
}
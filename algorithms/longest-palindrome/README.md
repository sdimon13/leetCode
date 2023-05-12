# 409. Longest Palindrome
Level: `Easy`

https://leetcode.com/problems/longest-palindrome/

---

# Description

Given a string `s` which consists of lowercase or uppercase letters, return the length of the **longest palindrome** that can be built with those letters.

Letters are **case sensitive**, for example, `"Aa"` is not considered a palindrome here.

## Example 1:

    Input: s = "abccccdd"
    Output: 7
    Explanation: One longest palindrome that can be built is "dccaccd", whose length is 7.

## Example 2:

    Input: s = "a"
    Output: 1
    Explanation: The longest palindrome that can be built is "a", whose length is 1.

## Constraints:

 - `1 <= s.length <= 2000`

 - `s` consists of lowercase **and/or** uppercase English letters only.

---

# Solution

## Intuition
The problem requires us to determine the length of the longest palindrome that can be built from the characters of a given string. A palindrome is a sequence that reads the same backwards as forwards. In a palindrome, all characters appear an even number of times, except for at most one character that can appear an odd number of times (this character would be in the middle of the palindrome). Therefore, to solve this problem, we need to count the characters in the string and calculate how many of them could be used to construct a palindrome.

## Approach
We use a hash map (in Go, a **map[rune]bool**) to count the characters in the string. For each character, we check if it's already in the hash map. If it is, we have found a pair of this character, so we increment the maximum length of the palindrome by 2 and remove the character from the hash map. If it's not, we add the character to the hash map.

After going through all the characters, we check if there are any characters left in the hash map. If there are, it means we have characters that appear an odd number of times in the string. We can take one of these characters and put it in the middle of the palindrome, so we increment the maximum length by 1.

## Complexity
- Time complexity:
  The time complexity is O(n), where n is the length of the string. We need to go through all the characters in the string once.

- Space complexity:
  The space complexity is O(1), because the hash map stores at most all unique characters from the string, and the number of unique characters can be at most 128 (for ASCII), which is a constant.

---

# Code
Here are the solution and test files:
- [Solution Code](./solution.go)
- [Test Code](./solution_test.go)
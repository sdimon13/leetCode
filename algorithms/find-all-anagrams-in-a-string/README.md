# 438. Find All Anagrams in a String

Level: `Medium`

https://leetcode.com/problems/find-all-anagrams-in-a-string/

---

# Description

Given two strings `s` and `p`, return an array of all the start indices of `p`'s anagrams in `s`. You may return the answer in **any order**.

An **Anagram** is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.

## Example 1:

    Input: s = "cbaebabacd", p = "abc"
    Output: [0,6]
    Explanation:
    The substring with start index = 0 is "cba", which is an anagram of "abc".
    The substring with start index = 6 is "bac", which is an anagram of "abc".

## Example 2:

    Input: s = "abab", p = "ab"
    Output: [0,1,2]
    Explanation:
    The substring with start index = 0 is "ab", which is an anagram of "ab".
    The substring with start index = 1 is "ba", which is an anagram of "ab".
    The substring with start index = 2 is "ab", which is an anagram of "ab".

## Constraints:

- `1 <= s.length, p.length <= 3 * 10`<sup>`4`</sup>
- `s` and `p` consist of lowercase English letters.

---

# Solution

## Intuition
Upon first glance at this problem, the technique of "sliding window" in combination with "character frequency counting" seems like a promising approach. The primary challenge is finding all anagrams of a string **p** within a string **s**, and the key insight here is recognizing that two strings are anagrams of each other if their character frequencies are the same.

## Approach
1. Create two frequency count arrays of size 26 (for each alphabet letter), one to track the frequency of characters in the target string (**p**), and another to track the frequency in the current sliding window in **s**.
2. Traverse the string **p** and increment the corresponding counter in the first array for each character.
3. Begin scanning the string **s** with a sliding window of the size of **p**. Increment the corresponding counter in the second array for each character.
4. Compare the two frequency count arrays. If they are the same, we have found an anagram, and add the starting index of this window to the result list.
5. After each step, move the window to the right. Decrease the counter for the character that is moving out of the window, and increase the counter for the character that is entering the window.
6. Continue steps 4 and 5 until the window reaches the end of string **s**.
7. At the end of the algorithm, the result list will contain the starting indices of all the found anagrams.

## Complexity
- Time complexity:
  The time complexity is O(n) because we are passing through the string **s** only once, where **n** is the length of the string **s**.

- Space complexity:
  The space complexity is O(1) because the size of the frequency arrays is fixed (26 for the number of lowercase English letters), and the result list's size does not exceed the length of the string **s**.

---

# Code
Here are the solution and test files:
- [Solution Code](./solution.go)
- [Test Code](./solution_test.go)
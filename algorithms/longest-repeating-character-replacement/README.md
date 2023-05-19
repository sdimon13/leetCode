# 424. Longest Repeating Character Replacement

Level: `Medium`

https://leetcode.com/problems/longest-repeating-character-replacement/

---

# Description

You are given a string `s` and an integer `k`. You can choose any character of the string and change it to any other uppercase English character. You can perform this operation at most `k` times.

Return the length of the longest substring containing the same letter you can get after performing the above operations.

## Example 1:

    Input: s = "ABAB", k = 2
    Output: 4
    Explanation: Replace the two 'A's with two 'B's or vice versa.

## Example 2:

    Input: s = "AABABBA", k = 1
    Output: 4
    Explanation: Replace the one 'A' in the middle with 'B' and form "AABBBBA".
    The substring "BBBB" has the longest repeating letters, which is 4.
    There may exists other ways to achive this answer too.

## Constraints:

- `1 <= s.length <= 10`<sup>`5`</sup>
- `s` consists of only uppercase English letters.
- 0 <= k <= s.length

---

# Solution

## Intuition
Upon first seeing this problem, it seems we are dealing with substring issues. The hint given is that we need to find the longest repeating character by replacement. It appears to be a sliding window problem because we want to keep the maximum length of the substring that could be converted to a string of repeating characters by replacing 'k' characters.

## Approach
Our approach to this problem will be using a sliding window concept and an array to keep track of character frequencies.

1. We initialize an array **counts** of length 26 (as there are 26 English alphabets), which will store the count of each character in the current window. The array indices correspond to the alphabets (A-Z), with the values at these indices being the counts of each character in the current window.

2. We then initialize variables **maxCount**, **start**, and **maxLength**, which represent the count of the most common character in the window, the starting index of the window, and the maximum length of such a window respectively.

3. In a for loop, we traverse every character in the string s. On each step, the count of the respective character in **counts** is incremented, and **maxCount** is updated.

4. We then check if the current window length **end-start+1** exceeds **maxCount + k** (the maximum possible window length with 'k' replacements). If it does, we decrement the count for the character at the **start** index in **counts** and shift the window's start to the right.

5. On each loop iteration, **maxLength** is updated, maintaining the maximum length of such a window encountered so far.

6. Finally, the function returns **maxLength**, which is the length of the longest substring that can be made into a repeating character substring by replacing 'k' characters.

## Complexity
- Time complexity:
  The time complexity is O(n), where n is the length of the string. This is because we are processing each character exactly once.

- Space complexity:
  The space complexity is O(1), as we are using a constant space to store the counts of each character (26 English alphabets), regardless of the input string size.

---

# Code
Here are the solution and test files:
- [Solution Code](./solution.go)
- [Test Code](./solution_test.go)
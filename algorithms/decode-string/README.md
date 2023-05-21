# 394. Decode String

Level: `Medium`

https://leetcode.com/problems/decode-string/

---

# Description

Given an encoded string, return its decoded string.

The encoding rule is: `k[encoded_string]`, where the `encoded_string` inside the square brackets is being repeated exactly `k` times. Note that `k` is guaranteed to be a positive integer.

You may assume that the input string is always valid; there are no extra white spaces, square brackets are well-formed, etc. Furthermore, you may assume that the original data does not contain any digits and that digits are only for those repeat numbers, `k`. For example, there will not be input like `3a` or `2[4]`.

The test cases are generated so that the length of the output will never exceed `10`<sup>`5`</sup>.

## Example 1:

    Input: s = "3[a]2[bc]"
    Output: "aaabcbc"

## Example 2:

    Input: s = "3[a2[c]]"
    Output: "accaccacc"

## Example 3:

    Input: s = "2[abc]3[cd]ef"
    Output: "abcabccdcdcdef"

## Constraints:

- `1 <= s.length <= 30`
- `s` consists of lowercase English letters, digits, and square brackets `'[]'`.
- `s` is guaranteed to be **a valid** input.
- All the integers in `s` are in the range `[1, 300]`.

---

# Solution

## Intuition
Upon reading the problem, it becomes clear that we need to parse the input string and decode it according to its embedded instructions, which involve repeating certain substrings. The concept of a "stack" comes to mind, as it is often used in situations where you need to keep track of nested structures or sequences.

## Approach
To solve this problem, we will use a stack to keep track of the current substring and its associated repeat count whenever we encounter an opening bracket '['. We will also maintain a **currentSubstring** and **repeatCount** to represent the currently processing substring and its repeat count.

Our approach is as follows:

- Parse the string from left to right.

- When we encounter a digit, we consider it as part of the **repeatCount** for the upcoming substring.

- When we encounter an opening bracket '[', we push the current **currentSubstring** and **repeatCount** to the stack, and reset them to process the new substring.

- When we encounter a closing bracket ']', we know that we have finished processing the current substring. We pop the topmost element from the stack, and append the **currentSubstring** repeated **repeatCount** times to the popped string.

- When we encounter a letter, we append it to **currentSubstring**.

## Complexity
- Time complexity:
  The time complexity is O(n), where n is the length of the input string. This is because we are processing each character in the string once.

- Space complexity:
  The space complexity is O(n) as well. In the worst case, every character could be part of the stack. For example, in the case of an input string like "100[a]", we would end up with a stack of size 100.

---

# Code
Here are the solution and test files:
- [Solution Code](./solution.go)
- [Test Code](./solution_test.go)
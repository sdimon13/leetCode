# 844. Backspace String Compare

Level: `Easy`

https://leetcode.com/problems/backspace-string-compare/

---

# Description

Given two strings `s` and `t`, return `true` if they are equal when both are typed into empty text editors. `'#'` means a backspace character.

Note that after backspacing an empty text, the text will continue empty.

## Example 1:

    Input: s = "ab#c", t = "ad#c"
    Output: true
    Explanation: Both s and t become "ac".

## Example 2:

    Input: s = "ab##", t = "c#d#"
    Output: true
    Explanation: Both s and t become "".

## Example 3:

    Input: s = "a#c", t = "b"
    Output: false
    Explanation: s becomes "c" while t becomes "b".

## Constraints:

- `1 <= s.length, t.length <= 200`
- `s` and `t` only contain lowercase letters and `'#'` characters.

---

# Solution

## Intuition
This problem can be approached by utilizing two pointers technique. We can start from the end of both strings and move our pointers backwards whenever we encounter a non-backspace character. If we encounter a backspace character '#', we need to skip the following character(s).

## Approach
1. Initialize two pointers **i** and **j** at the end of strings **s** and **t**, respectively.
2. Iterate while either **i** or **j** is not less than zero (i.e., we still have characters to examine).
   - Find the next non-backspace character in **s** by using a helper function **getNextIndex** and update **i**.
   - Do the same for **t** and update **j**.
   - If the non-backspace characters of **s** and **t** are different, return **false** because the strings are not the same.
   - If one of the strings is finished (i.e., the index is less than 0) and the other is not, return **false** because the lengths of the resulting strings would be different.
3. If none of the checks failed during the iteration, return true because the strings are the same after applying backspaces.

The **getNextIndex** helper function moves the given index to the next non-backspace character and returns the updated index. It counts the number of consecutive backspaces and skips the same number of non-backspace characters.
## Complexity
- Time complexity:
  The time complexity is O(n), where n is the total length of the two strings. In the worst-case scenario, we need to iterate through every character in both strings once.

- Space complexity:
  The space complexity is O(1), as we're using a constant amount of space to store our pointers and counters, and we're not creating any new data structures proportionate to the input size.

---

# Code
Here are the solution and test files:
- [Solution Code](./solution.go)
- [Test Code](./solution_test.go)
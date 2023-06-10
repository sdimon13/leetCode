# 1768. Merge Strings Alternately

Level: `Easy`

https://leetcode.com/problems/merge-strings-alternately/

---

# Description

You are given two strings `word1` and `word2`. Merge the strings by adding letters in alternating order, starting with
`word1`. If a string is longer than the other, append the additional letters onto the end of the merged string.

Return the merged string.

## Example 1:

    Input: word1 = "abc", word2 = "pqr"
    Output: "apbqcr"
    Explanation: The merged string will be merged as so:
    word1:  a   b   c
    word2:    p   q   r
    merged: a p b q c r

## Example 2:

    Input: word1 = "ab", word2 = "pqrs"
    Output: "apbqrs"
    Explanation: Notice that as word2 is longer, "rs" is appended to the end.
    word1:  a   b
    word2:    p   q   r   s
    merged: a p b q   r   s

## Example 3:

    Input: word1 = "abcd", word2 = "pq"
    Output: "apbqcd"
    Explanation: Notice that as word1 is longer, "cd" is appended to the end.
    word1:  a   b   c   d
    word2:    p   q
    merged: a p b q c   d

## Constraints:

- 1 <= word1.length, word2.length <= 100
- word1 and word2 consist of lowercase English letters.

---

# Solution

## Intuition

This problem is relatively straightforward. We are given two strings, and our task is to merge them alternatively. If
one string is longer than the other, we append the remaining letters from the longer string to the end of the merged
string.

## Approach

We will use two pointers **i** and **j** to navigate through **word1** and **word2** respectively. We keep adding one
character from
each string to the result until one of the pointers reaches the end of its string. At this point, we simply append the
remaining characters from the longer string to our result.

This algorithm works because it follows the exact procedure outlined in the problem: merge the strings by adding letters
in alternating order, and if a string is longer than the other, append the additional letters onto the end of the merged
string.

## Complexity

- Time complexity:
  The time complexity is O(n), where n is the total length of the two strings. This is because in the worst case
  scenario, we have to go through every character in both strings.

- Space complexity:
  The space complexity is O(n), where n is the total length of the two strings. This is because we are storing the
  result in a new string. In Go, strings are immutable, which means that every time we append a character, a new string
  gets created.

---

# Code

Here are the solution and test files:

- [Solution Code](./solution.go)
- [Test Code](./solution_test.go)
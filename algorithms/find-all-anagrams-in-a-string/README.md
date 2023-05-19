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
The problem asks for the calculation of the nth Fibonacci number. The Fibonacci sequence is a series of numbers where each number is the sum of the two preceding ones, usually starting with 0 and 1. That is, 0, 1, 1, 2, 3, 5, 8, 13, 21, 34, and so forth.

## Approach
An iterative approach is taken here to solve this problem efficiently. The idea is to use two variables to keep track of the last two numbers in the Fibonacci sequence. Start from the base case where the first number is 0 and the second number is 1. Then, iterate from 2 to n, in each iteration, calculate the next Fibonacci number as the sum of the last two numbers, and update the last two numbers accordingly.

## Complexity
- Time complexity:
  The time complexity is O(n) because we have to iterate through all numbers up to n.

- Space complexity:
  The space complexity is O(1) because we are only using two variables to keep track of the last two Fibonacci numbers, regardless of the size of n.

---

# Code
Here are the solution and test files:
- [Solution Code](./solution.go)
- [Test Code](./solution_test.go)
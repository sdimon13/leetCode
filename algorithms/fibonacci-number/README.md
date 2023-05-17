# 509. Fibonacci Number

Level: `Easy`

https://leetcode.com/problems/fibonacci-number/

---

# Description

The **Fibonacci numbers**, commonly denoted `F(n)` form a sequence, called the **Fibonacci sequence**, such that each number is the sum of the two preceding ones, starting from 0 and 1. That is,

    F(0) = 0, F(1) = 1
    F(n) = F(n - 1) + F(n - 2), for n > 1.
Given `n`, calculate `F(n)`.

## Example 1:

    Input: n = 2
    Output: 1
    Explanation: F(2) = F(1) + F(0) = 1 + 0 = 1.

## Example 2:

    Input: n = 3
    Output: 2
    Explanation: F(3) = F(2) + F(1) = 1 + 1 = 2.

## Example 3:

    Input: n = 4
    Output: 3
    Explanation: F(4) = F(3) + F(2) = 2 + 1 = 3.

## Constraints:

- `0 <= n <= 30`

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
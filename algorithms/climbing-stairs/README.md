# 70. Climbing Stairs

Level: `Easy`

https://leetcode.com/problems/climbing-stairs/

---

# Description

You are climbing a staircase. It takes `n` steps to reach the top.

Each time you can either climb `1` or `2` steps. In how many distinct ways can you climb to the top?

## Example 1:

    Input: n = 2
    Output: 2
    Explanation: There are two ways to climb to the top.
    1. 1 step + 1 step
    2. 2 steps

## Example 2:

    Input: n = 3
    Output: 3
    Explanation: There are three ways to climb to the top.
    1. 1 step + 1 step + 1 step
    2. 1 step + 2 steps
    3. 2 steps + 1 step

## Constraints:

- `1 <= n <= 45`

---

# Solution

## Intuition
This problem is a classic case of dynamic programming. The problem is to find out the number of ways to climb a staircase, given that at a time we can either take 1 step or 2 steps.

## Approach
The number of ways to reach the **ith** step is the sum of ways to reach the **(i-1)th** step and the **(i-2)th** step, because we can arrive at the **ith** step by taking a single step from **(i-1)** or by taking two steps from **(i-2)**.

This behavior is exactly like the Fibonacci sequence, where each number (starting from the third) is the sum of the two preceding ones.

We use two variables **a** and **b** to keep track of the last two Fibonacci numbers starting from **a=1** and **b=2**. Then we simply loop from **2** to **n**, updating **a** and **b** to the next number in the Fibonacci sequence at each step.

Finally, **b** becomes the **nth** number in the Fibonacci sequence, which is the answer to our problem.

## Complexity
- Time complexity:
  The time complexity of this algorithm is O(n), as we are essentially calculating the nth Fibonacci number.

- Space complexity:
  The space complexity is O(1), as we are only using two variables a and b to keep track of the last two Fibonacci numbers.

---

# Code
Here are the solution and test files:
- [Solution Code](./solution.go)
- [Test Code](./solution_test.go)
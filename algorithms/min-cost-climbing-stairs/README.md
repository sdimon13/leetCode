# 746. Min Cost Climbing Stairs

Level: `Easy`

https://leetcode.com/problems/min-cost-climbing-stairs/

---

# Description

You are given an integer array `cost` where `cost[i]` is the cost of `i`<sup>`th`</sup> step on a staircase. Once you pay the cost, you can either climb one or two steps.

You can either start from the step with index `0`, or the step with index `1`.

Return the minimum cost to reach the top of the floor.

## Example 1:

    Input: cost = [10,15,20]
    Output: 15
    Explanation: You will start at index 1.
    - Pay 15 and climb two steps to reach the top.
      The total cost is 15.

## Example 2:

    Input: cost = [1,100,1,1,1,100,1,1,100,1]
    Output: 6
    Explanation: You will start at index 0.
    - Pay 1 and climb two steps to reach index 2.
      - Pay 1 and climb two steps to reach index 4.
      - Pay 1 and climb two steps to reach index 6.
      - Pay 1 and climb one step to reach index 7.
      - Pay 1 and climb two steps to reach index 9.
      - Pay 1 and climb one step to reach the top.
        The total cost is 6.

## Constraints:

- `2 <= cost.length <= 1000`
- `0 <= cost[i] <= 999`

---

# Solution

## Intuition
When we first look at this problem, we can see that the cost of reaching each step depends on the cost of reaching the previous steps. This means that we can't determine the cost of reaching a step without knowing the costs of reaching previous steps. This problem is a perfect candidate for a dynamic programming solution, where we use previously calculated results to calculate the next ones.

## Approach
Our approach to solving this problem is to use dynamic programming to build up a **dp** array where **dp[i]** represents the minimum cost to reach step **i**. We initialize the first two elements of **dp** with the first two elements of **cost** because we can start from either step.

Then, for each subsequent step, we calculate the minimum cost to reach it by considering the cost of reaching the previous step and the cost of reaching the step before that, adding the cost of the current step. This is done using **dp[i] = min(dp[i-1], dp[i-2]) + cost[i]**.

Finally, we return the minimum cost to reach the top of the stairs, which is the smaller of the costs to reach the last step and the second last step, i.e., **min(dp[n-1], dp[n-2])**.

## Complexity
- Time complexity:
  O(n), where n is the number of stairs. We need to iterate through the entire cost array once.

- Space complexity:
  O(n), where n is the number of stairs. We need to create a dp array of the same length as the cost array.

---

# Code
Here are the solution and test files:
- [Solution Code](./solution.go)
- [Test Code](./solution_test.go)
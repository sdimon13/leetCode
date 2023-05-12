# 121. Best Time to Buy and Sell Stock
Level: `Easy`

https://leetcode.com/problems/best-time-to-buy-and-sell-stock/

---

# Description

You are given an array `prices` where `prices[i]` is the price of a given stock on the `i`<sup>`th`</sup> day.

You want to maximize your profit by choosing a **single day** to buy one stock and choosing a **different day in the future** to sell that stock.

Return the _maximum profit you can achieve from this transaction_. If you cannot achieve any profit, return `0`.

## Example 1:

  Input: prices = [7,1,5,3,6,4]
  Output: 5
  Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
  Note that buying on day 2 and selling on day 1 is not allowed because you must buy before you sell.

## Example 2:

  Input: prices = [7,6,4,3,1]
  Output: 0
  Explanation: In this case, no transactions are done and the max profit = 0.

## Constraints:

- `1 <= prices.length <= 10`<sup>`5`</sup>

- `0 <= prices[i] <= 104`

---

# Solution

## Intuition
The problem can be solved by finding the maximum difference between any two numbers in the array such that the smaller number comes before the larger one. This means we need to find a "valley" followed by a "peak".

## Approach
The approach is to keep track of the smallest price seen so far (minPrice) and the maximum profit we can get (maxProfit). We iterate through the prices array, updating the minPrice if the current price is lower and updating the maxProfit if the difference between the current price and minPrice is greater than the current maxProfit.

## Complexity
- Time complexity:
  The time complexity is O(n), where n is the number of days (length of the prices array). This is because we only have a single loop that goes through the prices array.

- Space complexity:
  The space complexity is O(1), because we only use two variables (minPrice and maxProfit), and this does not change with the size of the input array.

---

# Code
Here are the solution and test files:
- [Solution Code](./solution.go)
- [Test Code](./solution_test.go)
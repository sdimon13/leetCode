# 1. Two Sum

Level: `Easy`

https://leetcode.com/problems/two-sum/

---

# Description

Given an array of integers `nums` and an integer `target`, return indices of the two numbers such that they add up to `target`.

You may assume that each input would have `exactly one solution`, and you may not use the same element twice.

You can return the answer in any order.

## Example 1:

    Input: nums = [2,7,11,15], target = 9
    Output: [0,1]
    Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].

## Example 2:

    Input: nums = [3,2,4], target = 6
    Output: [1,2]

## Example 3:

    Input: nums = [3,3], target = 6
    Output: [0,1]

## Constraints:

- `2 <= nums.length <= 10`<sup>`4`</sup>
- `-10`<sup>`9`</sup> `<= nums[i] <= 10`<sup>`9`</sup>
- `-10`<sup>`9`</sup>` <= target <= 10`<sup>`9`</sup>
- **Only one valid answer exists.**

---

# Solution

## Intuition
The problem requires us to find two numbers in the array that add up to a given target number. A brute force approach would be to iterate through each number and then iterate through the rest of the array looking for a complement number (the difference between the target and the current number). However, this would lead to a quadratic time complexity.

## Approach
A more efficient way to solve this problem is by using a hash map (or a dictionary). We iterate through the array once, using the hash map to store numbers as keys and their indices as values. For each number we encounter, we check if its complement is already in the hash map. If the complement is present, it means we have found two numbers that add up to the target, and we return their indices. If the complement is not present, we add the current number and its index to the hash map and move on to the next number.

## Complexity
- Time complexity:
  O(n). We perform a single pass through the array, so the time complexity is linear.

- Space complexity:
  O(n). In the worst case, we might end up storing every number in the hash map.

---

# Code
Here are the solution and test files:
- [Solution Code](./solution.go)
- [Test Code](./solution_test.go)
# 704. Binary Search
Level: `Easy`

https://leetcode.com/problems/binary-search/

---

# Description

Given an array of integers `nums` which is sorted in ascending order, and an integer `target`, write a function to search `target` in `nums`. If `target` exists, then return its index. Otherwise, return `-1`.

You must write an algorithm with `O(log n)` runtime complexity.

## Example 1:

    Input: nums = [-1,0,3,5,9,12], target = 9
    Output: 4
    Explanation: 9 exists in nums and its index is 4

## Example 2:

    Input: nums = [-1,0,3,5,9,12], target = 2
    Output: -1
    Explanation: 2 does not exist in nums so return -1

## Constraints:

 - `1 <= nums.length <= 10`<sup>`4`</sup>
 - `-10`<sup>4</sup> `< nums[i], target < 104`
 - All the integers in `nums` are **unique**.
 - `nums` is sorted in ascending order.

---

# Solution

## Intuition
The task asks to implement binary search in a sorted array. Binary search is a standard algorithm that is used to find the position of a target value within a sorted array. The approach of binary search is to reduce the search space by half at each step.

## Approach
We start with defining the search boundaries, **left** and **right**, at the start and end of the array, respectively. Then, in a loop, we check the middle element between **left** and **right**. If this element equals the **target**, we found the desired element and return its index. If the element is less than the **target**, we shift the left boundary (**left**) after the middle. If the element is greater than the **target**, we shift the right boundary (**right**) before the middle. If none of the values equal the **target**, we return **-1**.

## Complexity
- Time complexity: The time complexity of binary search is O(log n). With each comparison, we cut the search space in half. Thus, in the worst case, we need log n comparisons where n is the number of elements in the array.

- Space complexity: The space complexity is O(1) as we are not using any extra space that scales with input size.

---

# Code
Here are the solution and test files:
- [Solution Code](./solution.go)
- [Test Code](./solution_test.go)
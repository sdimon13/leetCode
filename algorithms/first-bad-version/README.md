# 278. First Bad Version
Level: `Easy`

https://leetcode.com/problems/first-bad-version/

---

# Description

You are a product manager and currently leading a team to develop a new product. Unfortunately, the latest version of your product fails the quality check. Since each version is developed based on the previous version, all the versions after a bad version are also bad.

Suppose you have `n` versions `[1, 2, ..., n]` and you want to find out the first bad one, which causes all the following ones to be bad.

You are given an API `bool isBadVersion(version)` which returns whether `version` is bad. Implement a function to find the first bad version. You should minimize the number of calls to the API.

## Example 1:

    Input: n = 5, bad = 4
    Output: 4
    Explanation:
    call isBadVersion(3) -> false
    call isBadVersion(5) -> true
    call isBadVersion(4) -> true
    Then 4 is the first bad version.

## Example 2:

    Input: n = 1, bad = 1
    Output: 1

## Constraints:

 - `1 <= bad <= n <= 2`<sup>`31`</sup>`- 1`

---

# Solution

## Intuition
The problem is asking for the first bad version in a sequence of versions. The basic intuition behind this problem is to use a binary search to find the first bad version. We are given a helper method **isBadVersion** which can check whether a version is bad or not.

## Approach
1. Initialize two pointers: left at 1 and right at n.
2. Perform a binary search. In each iteration, we calculate the middle element mid and check if it's a bad version using the isBadVersion method.
3. If the mid is a bad version, we know all versions after it are also bad, so we set right to mid. If it's not a bad version, we know all versions before it are also good, so we set left to mid + 1
4. We keep shrinking the range until left is equal to right, at which point we've found the first bad version.

## Complexity
- Time complexity:
  O(logn), because in each step we reduce the number of versions we are looking at by half.

- Space complexity:
  O(1), because we only use a constant amount of space to store our variables.

---

# Code
Here are the solution and test files:
- [Solution Code](./solution.go)
- [Test Code](./solution_test.go)
# 62. Unique Paths

Level: `Medium`

https://leetcode.com/problems/unique-paths/

---

# Description

There is a robot on an `m x n` grid. The robot is initially located at the **top-left corner** (i.e., `grid[0][0]`). The robot tries to move to the **bottom-right corner** (i.e., `grid[m - 1][n - 1]`). The robot can only move either down or right at any point in time.

Given the two integers `m` and `n`, return the number of possible unique paths that the robot can take to reach the bottom-right corner.

The test cases are generated so that the answer will be less than or equal to `2 * 10`<sup>`9`</sup>.

## Example 1:

![Example 1](https://assets.leetcode.com/uploads/2018/10/22/robot_maze.png)

    Input: m = 3, n = 7
    Output: 28

## Example 2:

    Input: m = 3, n = 2
    Output: 3
    Explanation: From the top-left corner, there are a total of 3 ways to reach the bottom-right corner:
    1. Right -> Down -> Down
    2. Down -> Down -> Right
    3. Down -> Right -> Down

## Constraints:

- `1 <= m, n <= 100`

---

# Solution

## Intuition
The problem can be solved using dynamic programming. Observe that each cell can be reached only from the cell above or the cell on the left. Therefore, the number of unique paths to a specific cell equals the sum of the unique paths to the cell above it and the cell to its left.

## Approach
Initialize a 2D array of size m*n (the size of the grid) where each cell represents the number of unique paths to reach that cell. Fill the first row and the first column with 1s because there's only one way to reach any cell in the first row or column (by continuously moving right or down). Then, iterate through the rest of the cells and set its value to be the sum of the value of the cell above it and the value of the cell to its left.

## Complexity
- Time complexity:
  The time complexity is O(m*n) because we have to go through every cell in the grid once.

- Space complexity:
  The space complexity is also O(m*n) because we create a 2D array to hold the number of unique paths to each cell.

---

# Code
Here are the solution and test files:
- [Solution Code](./solution.go)
- [Test Code](./solution_test.go)
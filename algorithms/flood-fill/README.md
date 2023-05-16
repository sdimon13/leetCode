# 733. Flood Fill
Level: `Easy`

https://leetcode.com/problems/flood-fill

---

# Description

An image is represented by an `m x n` integer grid `image` where `image[i][j]` represents the pixel value of the image.

You are also given three integers `sr`, `sc`, and `color`. You should perform a **flood fill** on the image starting from the pixel `image[sr][sc]`.

To perform a **flood fill**, consider the starting pixel, plus any pixels connected **4-directionally** to the starting pixel of the same color as the starting pixel, plus any pixels connected **4-directionally** to those pixels (also with the same color), and so on. Replace the color of all of the aforementioned pixels with `color`.

Return the modified image after performing the flood fill.

## Example 1:

![Example 1](https://assets.leetcode.com/uploads/2021/06/01/flood1-grid.jpg)

    Input: image = [[1,1,1],[1,1,0],[1,0,1]], sr = 1, sc = 1, color = 2
    Output: [[2,2,2],[2,2,0],[2,0,1]]
    Explanation: From the center of the image with position (sr, sc) = (1, 1) (i.e., the red pixel), all pixels connected by a path of the same color as the starting pixel (i.e., the blue pixels) are colored with the new color.
    Note the bottom corner is not colored 2, because it is not 4-directionally connected to the starting pixel.

## Example 2:

    Input: image = [[0,0,0],[0,0,0]], sr = 0, sc = 0, color = 0
    Output: [[0,0,0],[0,0,0]]
    Explanation: The starting pixel is already colored 0, so no changes are made to the image.

## Constraints:

- `m == image.length`
- `n == image[i].length`
- `1 <= m, n <= 50`
- `0 <= image[i][j], color < 2`<sup>`16`</sup>
- `0 <= sr < m`
- `0 <= sc < n`

---

# Solution

## Intuition
The problem is a classic "flood fill" problem that can be solved using Depth-First Search (DFS) or Breadth-First Search (BFS). The goal is to start from a pixel (represented by a cell in the matrix), and to change its color and the color of all the adjacent pixels of the same color to a new color.

## Approach
The approach to solving this problem involves using the Depth-First Search (DFS) algorithm. The DFS algorithm starts at the root (selecting some arbitrary node as the root in the case of a graph) and explores as far as possible along each branch before backtracking.

In this case, we start at the given pixel, and then continue to the adjacent pixels (up, down, left, and right), but only if they have the same color as the starting pixel and are within the boundaries of the image. If the starting pixel already has the new color, we don't need to do anything.

## Complexity
- Time complexity:
  The time complexity is O(n), where n is the number of pixels in the image. This is because in the worst case, we might need to visit all the pixels.

- Space complexity:
  The space complexity is also O(n), because in the worst case, if all pixels in the image have the same color, the depth of the recursion could go up to n.

---

# Code
Here are the solution and test files:
- [Solution Code](./solution.go)
- [Test Code](./solution_test.go)
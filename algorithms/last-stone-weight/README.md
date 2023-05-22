# 1046. Last Stone Weight

Level: `Easy`

https://leetcode.com/problems/last-stone-weight/

---

# Description

You are given an array of integers `stones` where `stones[i]` is the weight of the `i`<sup>`th`</sup> stone.

We are playing a game with the stones. On each turn, we choose the **heaviest two stones** and smash them together. Suppose the heaviest two stones have weights `x` and `y` with `x <= y`. The result of this smash is:

- If `x == y`, both stones are destroyed, and
- If `x != y`, the stone of weight x is destroyed, and the stone of weight y has new weight y - x.
At the end of the game, there is **at most one** stone left.

Return the weight of the last remaining stone. If there are no stones left, return `0`.

## Example 1:

    Input: stones = [2,7,4,1,8,1]
    Output: 1
    Explanation:
    We combine 7 and 8 to get 1 so the array converts to [2,4,1,1,1] then,
    we combine 2 and 4 to get 2 so the array converts to [2,1,1,1] then,
    we combine 2 and 1 to get 1 so the array converts to [1,1,1] then,
    we combine 1 and 1 to get 0 so the array converts to [1] then that's the value of the last stone.

## Example 2:

    Input: stones = [1]
    Output: 1

## Constraints:

- `1 <= stones.length <= 30`
- `1 <= stones[i] <= 1000`.

---

# Solution

## Intuition
In this problem, we are given the weights of stones, and we need to smash them following a specific rule until we are left with one or no stone. My first thought is to always smash the heaviest two stones together. To efficiently do this, we could use a max heap data structure, which can give us the heaviest stone in constant time.

## Approach
The approach is as follows:

1. We first create a MaxHeap data structure and add all the stone weights to it.
2. We keep removing the top two stones (which are the heaviest) from the max heap and smash them together.
3. If they are not of the same weight, we push the difference of their weights back to the heap.
4. We repeat this process until there is one or no stone left in the heap.
5. If there's one stone left, we return its weight; otherwise, we return 0 if no stones are left.

## Complexity
- Time complexity:
  The time complexity of this approach is O(n log n), where n is the number of stones. We are performing a heapify operation which takes O(n) time, and then we are performing heap pop and push operations inside a loop which run log n times for each stone, making the time complexity O(n log n).

- Space complexity:
  The space complexity is O(n), where n is the number of stones, as we are storing all the stones in the heap.

---

# Code
Here are the solution and test files:
- [Solution Code](./solution.go)
- [Test Code](./solution_test.go)
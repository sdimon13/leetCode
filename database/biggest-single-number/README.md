# 619. Biggest Single Number

Level: `Easy`

https://leetcode.com/problems/biggest-single-number/

---

# Description

Table: `MyNumbers`

    +-------------+------+
    | Column Name | Type |
    +-------------+------+
    | num         | int  |
    +-------------+------+
    There is no primary key for this table. It may contain duplicates.
    Each row of this table contains an integer.

A **single number** is a number that appeared only once in the MyNumbers table.

Write an SQL query to report the largest **single number**. If there is no **single number**, report null.

The query result format is in the following example.

## Example 1:

    Input:
    MyNumbers table:
    +-----+
    | num |
    +-----+
    | 8   |
    | 8   |
    | 3   |
    | 3   |
    | 1   |
    | 4   |
    | 5   |
    | 6   |
    +-----+
    Output:
    +-----+
    | num |
    +-----+
    | 6   |
    +-----+
    Explanation: The single numbers are 1, 4, 5, and 6.
    Since 6 is the largest single number, we return it.

## Example 2:

    Input:
    MyNumbers table:
    +-----+
    | num |
    +-----+
    | 8   |
    | 8   |
    | 7   |
    | 7   |
    | 3   |
    | 3   |
    | 3   |
    +-----+
    Output:
    +------+
    | num  |
    +------+
    | null |
    +------+
    Explanation: There are no single numbers in the input table so we return null.

---

# Solution

## Intuition

Given a table MyNumbers, we want to find the largest number that appears only once in the table. This problem involves
two key steps: first, we need to identify the numbers that appear only once in the table; second, we need to return the
largest among these numbers.

## Approach

To solve this problem, we can use a subquery. In the subquery, we group the table by the num column and select the
groups having a count of 1, which gives us the single numbers. Then, in the outer query, we select the maximum of these
single numbers.

## Complexity

- Time complexity:
  The time complexity is O(n), where n is the number of rows in the table. This is because we need to examine each row
  in order to count the occurrences of each number.

- Space complexity:
  The space complexity is O(m), where m is the number of unique numbers in the table. This is because we need to store
  the count of each unique number.

---

# Code

Here are the solution and test files:

- [Solution Code](./solution.sql)
- [Test Code](./solution_test.go)
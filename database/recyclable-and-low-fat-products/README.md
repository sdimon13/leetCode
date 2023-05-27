# 1757. Recyclable and Low Fat Products

Level: `Easy`

https://leetcode.com/problems/delete-duplicate-emails/

---

# Description

Table: `Products`

    +-------------+---------+
    | Column Name | Type    |
    +-------------+---------+
    | product_id  | int     |
    | low_fats    | enum    |
    | recyclable  | enum    |
    +-------------+---------+
    product_id is the primary key for this table.
    low_fats is an ENUM of type ('Y', 'N') where 'Y' means this product is low fat and 'N' means it is not.
    recyclable is an ENUM of types ('Y', 'N') where 'Y' means this product is recyclable and 'N' means it is not.


Write an SQL query to find the ids of products that are both low fat and recyclable.

Return the result table in **any order**.

The query result format is in the following example.

## Example 1:

    Input:
    Products table:
    +-------------+----------+------------+
    | product_id  | low_fats | recyclable |
    +-------------+----------+------------+
    | 0           | Y        | N          |
    | 1           | Y        | Y          |
    | 2           | N        | Y          |
    | 3           | Y        | Y          |
    | 4           | N        | N          |
    +-------------+----------+------------+
    Output:
    +-------------+
    | product_id  |
    +-------------+
    | 1           |
    | 3           |
    +-------------+
    Explanation: Only products 1 and 3 are both low fat and recyclable.

---

# Solution

## Intuition
The problem is a simple database query exercise where we need to select certain records based on specified conditions. Specifically, we are asked to find products that are both low fat and recyclable.

## Approach
To solve this problem, we will use SQL's **SELECT** command along with a **WHERE** clause to specify the conditions that the selected records must meet. In this case, we need to find products where **low_fats** is 'Y' (indicating that the product is low fat) and **recyclable** is 'Y' (indicating that the product is recyclable).

## Complexity
- Time complexity:
  The time complexity of this SQL query is O(n), where n is the total number of products in the table. This is because the query needs to check each product in the table to see if it meets the specified conditions.

- Space complexity:
  The space complexity is O(1) because we are not using any extra space that scales with the size of the input. The memory used to store the result set is not typically counted towards space complexity.

---

# Code
Here are the solution and test files:
- [Solution Code](./solution.sql)
- [Test Code](./solution_test.go)
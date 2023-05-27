# 175. Combine Two Tables

Level: `Easy`

https://leetcode.com/problems/combine-two-tables/

---

# Description

Table: `Person`

    +-------------+---------+
    | Column Name | Type    |
    +-------------+---------+
    | personId    | int     |
    | lastName    | varchar |
    | firstName   | varchar |
    +-------------+---------+
    personId is the primary key column for this table.
    This table contains information about the ID of some persons and their first and last names.

Table: `Address`

    +-------------+---------+
    | Column Name | Type    |
    +-------------+---------+
    | addressId   | int     |
    | personId    | int     |
    | city        | varchar |
    | state       | varchar |
    +-------------+---------+
    addressId is the primary key column for this table.
    Each row of this table contains information about the city and state of one person with ID = PersonId.

Write an SQL query to report the first name, last name, city, and state of each person in the `Person` table. If the
address of a `personId` is not present in the `Address` table, report `null` instead.

Return the result table in **any order**.

The query result format is in the following example.

## Example 1:

    Input:
    Person table:
    +----------+----------+-----------+
    | personId | lastName | firstName |
    +----------+----------+-----------+
    | 1        | Wang     | Allen     |
    | 2        | Alice    | Bob       |
    +----------+----------+-----------+
    Address table:
    +-----------+----------+---------------+------------+
    | addressId | personId | city          | state      |
    +-----------+----------+---------------+------------+
    | 1         | 2        | New York City | New York   |
    | 2         | 3        | Leetcode      | California |
    +-----------+----------+---------------+------------+
    Output:
    +-----------+----------+---------------+----------+
    | firstName | lastName | city          | state    |
    +-----------+----------+---------------+----------+
    | Allen     | Wang     | Null          | Null     |
    | Bob       | Alice    | New York City | New York |
    +-----------+----------+---------------+----------+
    Explanation:
    There is no address in the address table for the personId = 1 so we return null in their city and state.
    addressId = 1 contains information about the address of personId = 2.

---

# Solution

## Intuition

We need to report the first name, last name, city, and state of each person in the Person table. But the catch is that,
the address of a person might not be present in the Address table, in that case, we need to report null. This kind of
problem hints towards using a SQL JOIN, specifically a LEFT JOIN. We want to include all the people from the Person
table, and if they have corresponding address in the Address table we pick that, otherwise we pick null.

## Approach

We use a LEFT JOIN operation to solve this problem. The LEFT JOIN keyword returns all records from the left table (
Person), and the matched records from the right table (Address). The result is NULL from the right side (Address table),
if there is no match. This will ensure we get all persons irrespective of whether their address exists or not.

## Complexity

- Time complexity:
  The time complexity is O(n), where n is the number of persons in the Person table. This is because we need to scan
  through each person once.

- Space complexity:
  The space complexity is O(1), because we are not using any extra space that scales with the input size. We are only
  returning the final result set, which does not count towards space complexity.

---

# Code

Here are the solution and test files:

- [Solution Code](./solution.sql)
- [Test Code](./solution_test.go)
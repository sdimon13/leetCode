# 1693. Daily Leads and Partners

Level: `Easy`

https://leetcode.com/problems/daily-leads-and-partners/

---

# Description

Table: `DailySales`

    +-------------+---------+
    | Column Name | Type    |
    +-------------+---------+
    | date_id     | date    |
    | make_name   | varchar |
    | lead_id     | int     |
    | partner_id  | int     |
    +-------------+---------+
    This table does not have a primary key.
    This table contains the date and the name of the product sold and the IDs of the lead and partner it was sold to.
    The name consists of only lowercase English letters.

Write an SQL query that will, for each `date_id` and `make_name`, return the number of **distinct** lead_id's and *
*distinct** partner_id's.

Return the result table in **any order**.

The query result format is in the following example.

## Example 1:

    Input:
    DailySales table:
    +-----------+-----------+---------+------------+
    | date_id   | make_name | lead_id | partner_id |
    +-----------+-----------+---------+------------+
    | 2020-12-8 | toyota    | 0       | 1          |
    | 2020-12-8 | toyota    | 1       | 0          |
    | 2020-12-8 | toyota    | 1       | 2          |
    | 2020-12-7 | toyota    | 0       | 2          |
    | 2020-12-7 | toyota    | 0       | 1          |
    | 2020-12-8 | honda     | 1       | 2          |
    | 2020-12-8 | honda     | 2       | 1          |
    | 2020-12-7 | honda     | 0       | 1          |
    | 2020-12-7 | honda     | 1       | 2          |
    | 2020-12-7 | honda     | 2       | 1          |
    +-----------+-----------+---------+------------+
    Output:
    +-----------+-----------+--------------+-----------------+
    | date_id   | make_name | unique_leads | unique_partners |
    +-----------+-----------+--------------+-----------------+
    | 2020-12-8 | toyota    | 2            | 3               |
    | 2020-12-7 | toyota    | 1            | 2               |
    | 2020-12-8 | honda     | 2            | 2               |
    | 2020-12-7 | honda     | 3            | 2               |
    +-----------+-----------+--------------+-----------------+
    Explanation:
    For 2020-12-8, toyota gets leads = [0, 1] and partners = [0, 1, 2] while honda gets leads = [1, 2] and partners = [1, 2].
    For 2020-12-7, toyota gets leads = [0] and partners = [1, 2] while honda gets leads = [0, 1, 2] and partners = [1, 2].

---

# Solution

## Intuition

We can solve this problem by using the SQL GROUP BY statement, which groups rows that have the same values in specified
columns into aggregated data. This can be used in conjunction with COUNT DISTINCT to get unique counts for both *
*lead_id** and **partner_id**.

## Approach

1. Group the data by **date_id** and **make_name**. These are the two criteria we are interested in.
2. For each of these groups, calculate the count of unique **lead_ids** and **partner_ids** using the **COUNT(DISTINCT
   column_name)** function. This gives us the number of distinct **lead_ids** and **partner_ids** for each group.
3. Assign the counts from step 2 to new columns named **unique_leads** and **unique_partners** respectively.
4. Return the grouped data.

## Complexity

- Time complexity:
  The time complexity is O(n), where n is the number of rows in the table. The reason is that we are performing the
  operations on each row of the table once.

- Space complexity:
  The space complexity is O(1), because we are not using any additional data structures that scale with the size of the
  input.

---

# Code

Here are the solution and test files:

- [Solution Code](./solution.sql)
- [Test Code](./solution_test.go)
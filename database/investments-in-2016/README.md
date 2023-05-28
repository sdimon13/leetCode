# 585. Investments in 2016

Level: `Medium`

https://leetcode.com/problems/investments-in-2016/

---

# Description

Table: `Insurance`

    +-------------+-------+
    | Column Name | Type  |
    +-------------+-------+
    | pid         | int   |
    | tiv_2015    | float |
    | tiv_2016    | float |
    | lat         | float |
    | lon         | float |
    +-------------+-------+
    pid is the primary key column for this table.
    Each row of this table contains information about one policy where:
    pid is the policyholder's policy ID.
    tiv_2015 is the total investment value in 2015 and tiv_2016 is the total investment value in 2016.
    lat is the latitude of the policy holder's city. It's guaranteed that lat is not NULL.
    lon is the longitude of the policy holder's city. It's guaranteed that lon is not NULL.

Write an SQL query to report the sum of all total investment values in 2016 tiv_2016, for all policyholders who:

- have the same `tiv_2015` value as one or more other policyholders, and
- are not located in the same city like any other policyholder (i.e., the (`lat, lon`) attribute pairs must be unique).

Round `tiv_2016` to **two decimal places**.

The query result format is in the following example.

## Example 1:

    Input:
    Insurance table:
    +-----+----------+----------+-----+-----+
    | pid | tiv_2015 | tiv_2016 | lat | lon |
    +-----+----------+----------+-----+-----+
    | 1   | 10       | 5        | 10  | 10  |
    | 2   | 20       | 20       | 20  | 20  |
    | 3   | 10       | 30       | 20  | 20  |
    | 4   | 10       | 40       | 40  | 40  |
    +-----+----------+----------+-----+-----+
    Output:
    +----------+
    | tiv_2016 |
    +----------+
    | 45.00    |
    +----------+
    Explanation:
    The first record in the table, like the last record, meets both of the two criteria.
    The tiv_2015 value 10 is the same as the third and fourth records, and its location is unique.
    
    The second record does not meet any of the two criteria. Its tiv_2015 is not like any other policyholders and its location is the same as the third record, which makes the third record fail, too.
    So, the result is the sum of tiv_2016 of the first and last record, which is 45.

---

# Solution

## Intuition

This problem can be solved by leveraging the SQL's subqueries and aggregation functions to filter out the rows that
satisfy the conditions. The conditions are:

1. The **tiv_2015** of the insurance policy is the same as one or more other insurance policies.
2. The latitude and longitude of the insurance policy is not the same as any other insurance policy.

Once we filter out these rows, we sum up the **tiv_2016** values and round it to two decimal places.

## Approach

We utilize the SQL **IN** clause to filter the rows that have the same **tiv_2015** as one or more other rows. To
achieve this, we create a subquery that groups the table by **tiv_2015** and uses **HAVING COUNT(*) > 1** to filter the
groups that appear more than once.

To filter the rows that have unique latitude and longitude, we use the SQL **NOT EXISTS** clause to check whether there
exists a row with the same latitude and longitude but a different **pid**.

Finally, we calculate the sum of **tiv_2016** for all remaining rows using **SUM()**, and round the result to two
decimal places with **ROUND()**.

## Complexity

- Time complexity:
  O(n<sup>2</sup>) - The query involves subqueries that scan the entire table for each row, resulting in a quadratic
  time complexity.

- Space complexity:
  O(n) - The space complexity is linear because it needs to store the entire table in memory.

---

# Code

Here are the solution and test files:

- [Solution Code](./solution.sql)
- [Test Code](./solution_test.go)
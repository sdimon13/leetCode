# 595. Big Countries

Level: `Easy`

https://leetcode.com/problems/big-countries/

---

# Description

Table: `World`

    +-------------+---------+
    | Column Name | Type    |
    +-------------+---------+
    | name        | varchar |
    | continent   | varchar |
    | area        | int     |
    | population  | int     |
    | gdp         | bigint  |
    +-------------+---------+
    name is the primary key column for this table.
    Each row of this table gives information about the name of a country, the continent to which it belongs, its area, the population, and its GDP value.


A country is **big** if:

- it has an area of at least three million (i.e., `3000000 km`<sup>`2`</sup>), or
- it has a population of at least twenty-five million (i.e., `25000000`).
Write an SQL query to report the name, population, and area of the **big countries**.

Return the result table in **any order**.

The query result format is in the following example.

## Example 1:

    Input:
    World table:
    +-------------+-----------+---------+------------+--------------+
    | name        | continent | area    | population | gdp          |
    +-------------+-----------+---------+------------+--------------+
    | Afghanistan | Asia      | 652230  | 25500100   | 20343000000  |
    | Albania     | Europe    | 28748   | 2831741    | 12960000000  |
    | Algeria     | Africa    | 2381741 | 37100000   | 188681000000 |
    | Andorra     | Europe    | 468     | 78115      | 3712000000   |
    | Angola      | Africa    | 1246700 | 20609294   | 100990000000 |
    +-------------+-----------+---------+------------+--------------+
    Output:
    +-------------+------------+---------+
    | name        | population | area    |
    +-------------+------------+---------+
    | Afghanistan | 25500100   | 652230  |
    | Algeria     | 37100000   | 2381741 |
    +-------------+------------+---------+

---

# Solution

## Intuition
Given the SQL schema of the table **World**, we need to find countries that are big either by population or by area. Here 'big' is defined by the conditions: population is at least 25 million or area is at least 3 million square kilometers.

## Approach
We can solve this problem with a single SQL SELECT statement. By using a WHERE clause, we can filter out the rows that do not meet the criteria. The OR operator in the WHERE clause ensures that if either condition is met (population is at least 25 million or area is at least 3 million), the country will be selected.

## Complexity
- Time complexity:
  The time complexity of this solution is O(n), where n is the total number of rows in the table. This is because the database needs to traverse all rows in the table to check the conditions in the WHERE clause.

- Space complexity:
  The space complexity is O(1), since no additional space is required to solve this problem. The output space does not count towards space complexity.

---

# Code
Here are the solution and test files:
- [Solution Code](./solution.sql)
- [Test Code](./solution_test.go)
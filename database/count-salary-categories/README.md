# 1907. Count Salary Categories

Level: `Medium`

https://leetcode.com/problems/count-salary-categories/

---

# Description

Table: `Accounts`

    +-------------+------+
    | Column Name | Type |
    +-------------+------+
    | account_id  | int  |
    | income      | int  |
    +-------------+------+
    account_id is the primary key for this table.
    Each row contains information about the monthly income for one bank account.

Write an SQL query to report the number of bank accounts of each salary category. The salary categories are:

`"Low Salary"`: All the salaries **strictly less** than `$20000`.
`"Average Salary"`: All the salaries in the **inclusive** range `[$20000, $50000]`.
`"High Salary"`: All the salaries **strictly greater** than `$50000`.
The result table **must** contain all three categories. If there are no accounts in a category, then report `0`.

Return the result table in **any order**.

The query result format is in the following example.

## Example 1:

    Input:
    Accounts table:
    +------------+--------+
    | account_id | income |
    +------------+--------+
    | 3          | 108939 |
    | 2          | 12747  |
    | 8          | 87709  |
    | 6          | 91796  |
    +------------+--------+
    Output:
    +----------------+----------------+
    | category       | accounts_count |
    +----------------+----------------+
    | Low Salary     | 1              |
    | Average Salary | 0              |
    | High Salary    | 3              |
    +----------------+----------------+
    Explanation:
    Low Salary: Account 2.
    Average Salary: No accounts.
    High Salary: Accounts 3, 6, and 8.

---

# Solution

## Intuition

The task requires categorizing incomes into three categories: 'Low Salary', 'Average Salary', and 'High Salary', with a
certain income range defining each category. Then, the number of accounts falling into each category needs to be
counted. SQL provides functions to select and count data based on conditions, making it an excellent fit for this
problem.

## Approach

We can use SQL's **SELECT** statement with the **COUNT(*)** function to count the number of accounts whose income falls
into a specific category. The categories are specified as ranges with SQL's WHERE clause. Each category is selected
separately and then combined into a single result using the **UNION** operator, ensuring that all categories are always
present in the result, even if the count is zero. Each **SELECT** statement is associated with a string that specifies
the category.

The categories and their ranges are as follows:

- "Low Salary": All the salaries strictly less than $20000.

- "Average Salary": All the salaries in the inclusive range [$20000, $50000].

- "High Salary": All the salaries strictly greater than $50000.

## Complexity

- Time complexity:
  The time complexity is O(n) where n is the number of rows in the Accounts table. Each row must be examined to
  determine which category it belongs to.

- Space complexity:
  The space complexity is O(1) because the query only returns a fixed number of rows (one for each category), regardless
  of the size of the Accounts table.

---

# Code

Here are the solution and test files:

- [Solution Code](./solution.sql)
- [Test Code](./solution_test.go)
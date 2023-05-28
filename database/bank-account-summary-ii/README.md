# 1587. Bank Account Summary II

Level: `Easy`

https://leetcode.com/problems/bank-account-summary-ii/

---

# Description

Table: `Users`

    +--------------+---------+
    | Column Name  | Type    |
    +--------------+---------+
    | account      | int     |
    | name         | varchar |
    +--------------+---------+
    account is the primary key for this table.
    Each row of this table contains the account number of each user in the bank.
    There will be no two users having the same name in the table.

Table: `Transactions`

    +---------------+---------+
    | Column Name   | Type    |
    +---------------+---------+
    | trans_id      | int     |
    | account       | int     |
    | amount        | int     |
    | transacted_on | date    |
    +---------------+---------+
    trans_id is the primary key for this table.
    Each row of this table contains all changes made to all accounts.
    amount is positive if the user received money and negative if they transferred money.
    All accounts start with a balance of 0.

Write an SQL query to report the name and balance of users with a balance higher than `10000`. The balance of an account
is equal to the sum of the amounts of all transactions involving that account.

Return the result table in **any order**.

The query result format is in the following example.

## Example 1:

    Input:
    Users table:
    +------------+--------------+
    | account    | name         |
    +------------+--------------+
    | 900001     | Alice        |
    | 900002     | Bob          |
    | 900003     | Charlie      |
    +------------+--------------+
    Transactions table:
    +------------+------------+------------+---------------+
    | trans_id   | account    | amount     | transacted_on |
    +------------+------------+------------+---------------+
    | 1          | 900001     | 7000       |  2020-08-01   |
    | 2          | 900001     | 7000       |  2020-09-01   |
    | 3          | 900001     | -3000      |  2020-09-02   |
    | 4          | 900002     | 1000       |  2020-09-12   |
    | 5          | 900003     | 6000       |  2020-08-07   |
    | 6          | 900003     | 6000       |  2020-09-07   |
    | 7          | 900003     | -4000      |  2020-09-11   |
    +------------+------------+------------+---------------+
    Output:
    +------------+------------+
    | name       | balance    |
    +------------+------------+
    | Alice      | 11000      |
    +------------+------------+
    Explanation:
    Alice's balance is (7000 + 7000 - 3000) = 11000.
    Bob's balance is 1000.
    Charlie's balance is (6000 + 6000 - 4000) = 8000.

---

# Solution

## Intuition

The problem is asking to find the users who have a balance higher than 10,000. The balance of an account is the sum of
all transactions of the account. Therefore, we need to sum up all transactions for each account from the Transactions
table, and then join this with the Users table to get the name of the users.

## Approach

The task can be solved using SQL query. Our first step would be to join the Users and Transactions table using a LEFT
JOIN operation. The key for join is 'account' as it is common in both tables. Then we will use GROUP BY operation on '
account' and 'name' to group transactions for each user. SUM function will be used to calculate the sum of 'amount' for
each grouped user which will give us the balance. The HAVING clause is then used to filter out those users whose balance
is greater than 10,000.

## Complexity

- Time complexity:
  The time complexity would be O(n), where n is the number of transactions. This is because we are going through every
  transaction once.

- Space complexity:
  The space complexity would be O(m), where m is the number of users. This is because we are storing the sum of
  transactions for each user.

---

# Code

Here are the solution and test files:

- [Solution Code](./solution.sql)
- [Test Code](./solution_test.go)
# 1484. Group Sold Products By The Date

Level: `Easy`

https://leetcode.com/problems/fix-names-in-a-table/

---

# Description

Table `Activities`:

    +-------------+---------+
    | Column Name | Type    |
    +-------------+---------+
    | sell_date   | date    |
    | product     | varchar |
    +-------------+---------+
    There is no primary key for this table, it may contain duplicates.
    Each row of this table contains the product name and the date it was sold in a market.

Write an SQL query to find for each date the number of different products sold and their names.

The sold products names for each date should be sorted lexicographically.

Return the result table ordered by `sell_date`.

The query result format is in the following example.

## Example 1:

    Input:
    Activities table:
    +------------+------------+
    | sell_date  | product     |
    +------------+------------+
    | 2020-05-30 | Headphone  |
    | 2020-06-01 | Pencil     |
    | 2020-06-02 | Mask       |
    | 2020-05-30 | Basketball |
    | 2020-06-01 | Bible      |
    | 2020-06-02 | Mask       |
    | 2020-05-30 | T-Shirt    |
    +------------+------------+
    Output:
    +------------+----------+------------------------------+
    | sell_date  | num_sold | products                     |
    +------------+----------+------------------------------+
    | 2020-05-30 | 3        | Basketball,Headphone,T-shirt |
    | 2020-06-01 | 2        | Bible,Pencil                 |
    | 2020-06-02 | 1        | Mask                         |
    +------------+----------+------------------------------+
    Explanation:
    For 2020-05-30, Sold items were (Headphone, Basketball, T-shirt), we sort them lexicographically and separate them by a comma.
    For 2020-06-01, Sold items were (Pencil, Bible), we sort them lexicographically and separate them by a comma.
    For 2020-06-02, the Sold item is (Mask), we just return it.

---

# Solution

## Intuition

The problem is straightforward as it's mainly about string manipulation. The requirement is to format user names in a
specific way - only the first character should be in uppercase and the rest should be in lowercase.

## Approach

We can leverage built-in SQL functions to solve this problem. Specifically, **UPPER()**, **LOWER()**, **SUBSTR()**, and
**CONCAT()** functions will be handy.

The **SUBSTR()** function is used to extract a part of the string. We will extract the first character and the rest of
the string separately.

We will use **UPPER()** on the first character to make sure it's in uppercase, and **LOWER()** on the rest of the string
to convert it to lowercase.

The **CONCAT()** function will help us combine the first character (which is now in uppercase) and the rest of the
string (which is now in lowercase) together.

Finally, we sort the result by **user_id**.

## Complexity

- Time complexity:
  The time complexity of this SQL query is O(n), where n is the total number of users in the table. This is because the
  query needs to process each user's name to make sure it meets the specified conditions.

- Space complexity:
  The space complexity is O(1). We are not using any extra space that scales with the size of the input. The memory used
  to store the result set is not typically counted towards space complexity.

---

# Code

Here are the solution and test files:

- [Solution Code](./solution.sql)
- [Test Code](./solution_test.go)
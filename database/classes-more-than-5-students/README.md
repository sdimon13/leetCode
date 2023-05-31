# 596. Classes More Than 5 Students

Level: `Easy`

https://leetcode.com/problems/classes-more-than-5-students/

---

# Description

Table: `Courses`

    +-------------+---------+
    | Column Name | Type    |
    +-------------+---------+
    | student     | varchar |
    | class       | varchar |
    +-------------+---------+
    (student, class) is the primary key column for this table.
    Each row of this table indicates the name of a student and the class in which they are enrolled.

Write an SQL query to report all the classes that have **at least five students**.

Return the result table in **any order**.

The query result format is in the following example.

## Example 1:

    Input:
    Courses table:
    +---------+----------+
    | student | class    |
    +---------+----------+
    | A       | Math     |
    | B       | English  |
    | C       | Math     |
    | D       | Biology  |
    | E       | Math     |
    | F       | Computer |
    | G       | Math     |
    | H       | Math     |
    | I       | Math     |
    +---------+----------+
    Output:
    +---------+
    | class   |
    +---------+
    | Math    |
    +---------+
    Explanation:
    - Math has 6 students, so we include it.
    - English has 1 student, so we do not include it.
    - Biology has 1 student, so we do not include it.
    - Computer has 1 student, so we do not include it.

---

# Solution

## Intuition

The problem is asking for all classes that have at least five students. Given the table structure and the requirement,
it seems like a straightforward application of SQL's aggregation capabilities. We can group by the class and count the
number of distinct students in each class. We then need to filter these groups to include only those where the count is
at least five.

## Approach

1. We start by selecting the class column from the Courses table since that's the information we need in our result.
2. We then group the rows by the class column. This will create separate groups of rows for each unique class in the
   table.
3. After grouping, we apply the HAVING clause to filter the groups. The HAVING clause works similarly to the WHERE
   clause but it operates on groupings rather than individual rows. In this case, we use the COUNT function to count the
   number of distinct students in each group. We use the DISTINCT keyword to ensure that we're not counting the same
   student more than once if they're listed in the table multiple times for the same class.
4. Finally, we include only those groups where the count of distinct students is at least five.

## Complexity

- Time complexity:
  Since we're performing a group operation on the table, the time complexity is O(n), where n is the number of rows in
  the Courses table. This is because we have to examine each row to assign it to a group and then count the distinct
  students in each group.

- Space complexity:
  The space complexity is also O(n), because in the worst case, we might have to store information for each row when
  performing the grouping. In practice, SQL databases have optimizations that can reduce the space requirement.

---

# Code

Here are the solution and test files:

- [Solution Code](./solution.sql)
- [Test Code](./solution_test.go)
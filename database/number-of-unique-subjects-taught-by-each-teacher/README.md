# 2356. Number of Unique Subjects Taught by Each Teacher

Level: `Easy`

https://leetcode.com/problems/number-of-unique-subjects-taught-by-each-teacher/

---

# Description

Table: `Teacher`

    +-------------+------+
    | Column Name | Type |
    +-------------+------+
    | teacher_id  | int  |
    | subject_id  | int  |
    | dept_id     | int  |
    +-------------+------+
    (subject_id, dept_id) is the primary key for this table.
    Each row in this table indicates that the teacher with teacher_id teaches the subject subject_id in the department dept_id.

Write an SQL query to report the number of unique subjects each teacher teaches in the university.

Return the result table in **any order**.

The query result format is shown in the following example.

## Example 1:

    Input:
    Teacher table:
    +------------+------------+---------+
    | teacher_id | subject_id | dept_id |
    +------------+------------+---------+
    | 1          | 2          | 3       |
    | 1          | 2          | 4       |
    | 1          | 3          | 3       |
    | 2          | 1          | 1       |
    | 2          | 2          | 1       |
    | 2          | 3          | 1       |
    | 2          | 4          | 1       |
    +------------+------------+---------+
    Output:  
    +------------+-----+
    | teacher_id | cnt |
    +------------+-----+
    | 1          | 2   |
    | 2          | 4   |
    +------------+-----+
    Explanation:
    Teacher 1:
    - They teach subject 2 in departments 3 and 4.
    - They teach subject 3 in department 3.
      Teacher 2:
    - They teach subject 1 in department 1.
    - They teach subject 2 in department 1.
    - They teach subject 3 in department 1.
    - They teach subject 4 in department 1.

---

# Solution

## Intuition

The task requires finding the number of unique subjects taught by each teacher. SQL is a suitable language for this task
because it allows for efficient querying of structured data. The strategy would be to group the records by teacher_id
and count the number of distinct subject_id associated with each teacher_id. This count would give the number of unique
subjects taught by each teacher.

## Approach

- Use a **LEFT JOIN** on the tables **Visits** and **Transactions** on **visit_id**. This gives us all visits, including
  those that didn't result in a transaction (which are represented as NULL in the joined table).
- Use a **WHERE** clause to filter out the visits that didn't have a corresponding transaction (i.e., **transaction_id**
  is NULL).
- Use **GROUP BY** to group these visits by **customer_id**.
- Use **COUNT** to get the number of such visits for each customer.
- Select **customer_id** and the count as **count_no_trans**.

## Complexity

- Time complexity:
  The time complexity is O(n), where n is the number of rows in the table. This is because the GROUP BY clause needs to
  scan each row once.

- Space complexity:
  The space complexity is O(m), where m is the number of unique teacher_id in the table. This is the space needed to
  store the intermediate group-by result.

---

# Code

Here are the solution and test files:

- [Solution Code](./solution.sql)
- [Test Code](./solution_test.go)
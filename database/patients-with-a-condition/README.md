# 1527. Patients With a Condition

Level: `Easy`

https://leetcode.com/problems/patients-with-a-condition/

---

# Description

Table: `Patients`

    +--------------+---------+
    | Column Name  | Type    |
    +--------------+---------+
    | patient_id   | int     |
    | patient_name | varchar |
    | conditions   | varchar |
    +--------------+---------+
    patient_id is the primary key for this table.
    'conditions' contains 0 or more code separated by spaces.
    This table contains information of the patients in the hospital.

Write an SQL query to report the patient_id, patient_name and conditions of the patients who have Type I Diabetes. Type
I Diabetes always starts with DIAB1 prefix.

Return the result table in **any order**.

The query result format is in the following example.

## Example 1:

    Input:
    Patients table:
    +------------+--------------+--------------+
    | patient_id | patient_name | conditions   |
    +------------+--------------+--------------+
    | 1          | Daniel       | YFEV COUGH   |
    | 2          | Alice        |              |
    | 3          | Bob          | DIAB100 MYOP |
    | 4          | George       | ACNE DIAB100 |
    | 5          | Alain        | DIAB201      |
    +------------+--------------+--------------+
    Output:
    +------------+--------------+--------------+
    | patient_id | patient_name | conditions   |
    +------------+--------------+--------------+
    | 3          | Bob          | DIAB100 MYOP |
    | 4          | George       | ACNE DIAB100 |
    +------------+--------------+--------------+
    Explanation: Bob and George both have a condition that starts with DIAB1.

---

# Solution

## Intuition

Given the problem, we understand that we need to select all patients whose 'conditions' field contains the prefix '
DIAB1'. This essentially boils down to a simple string matching problem within SQL.

## Approach

To approach this problem, we can use the SQL LIKE operator, which allows us to perform pattern matching in strings. We
will look for the pattern 'DIAB1' within the 'conditions' field. There are two cases to consider here:

1. 'DIAB1' appears at the start of the string. In this case, we will use the pattern 'DIAB1%'.
2. 'DIAB1' appears in the middle or at the end of the string. In this case, we will use the pattern '% DIAB1%'. The
   space before 'DIAB1' ensures that 'DIAB1' is a separate code, not a part of another code.

We use the OR operator to combine these two cases in our WHERE clause.

## Complexity

- Time complexity:
  The time complexity for this query is O(n), where n is the number of rows in the table. This is because the database
  needs to scan each row and perform the pattern matching operation.

- Space complexity:
  The result set will contain as many rows as there are unique sell dates. In the worst case scenario, each row in the
  table has a unique sell date, so the space complexity is also O(n).

---

# Code

Here are the solution and test files:

- [Solution Code](./solution.sql)
- [Test Code](./solution_test.go)
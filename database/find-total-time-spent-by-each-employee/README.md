# 1741. Find Total Time Spent by Each Employee

Level: `Easy`

https://leetcode.com/problems/find-total-time-spent-by-each-employee/

---

# Description

Table: `Employees`

    +-------------+------+
    | Column Name | Type |
    +-------------+------+
    | emp_id      | int  |
    | event_day   | date |
    | in_time     | int  |
    | out_time    | int  |
    +-------------+------+
    (emp_id, event_day, in_time) is the primary key of this table.
    The table shows the employees' entries and exits in an office.
    event_day is the day at which this event happened, in_time is the minute at which the employee entered the office, and out_time is the minute at which they left the office.
    in_time and out_time are between 1 and 1440.
    It is guaranteed that no two events on the same day intersect in time, and in_time < out_time.

Write an SQL query to calculate the total time **in minutes** spent by each employee on each day at the office. Note
that
within one day, an employee can enter and leave more than once. The time spent in the office for a single entry is
`out_time - in_time`.

Return the result table in **any order**.

The query result format is in the following example.

## Example 1:

    Input:
    Employees table:
    +--------+------------+---------+----------+
    | emp_id | event_day  | in_time | out_time |
    +--------+------------+---------+----------+
    | 1      | 2020-11-28 | 4       | 32       |
    | 1      | 2020-11-28 | 55      | 200      |
    | 1      | 2020-12-03 | 1       | 42       |
    | 2      | 2020-11-28 | 3       | 33       |
    | 2      | 2020-12-09 | 47      | 74       |
    +--------+------------+---------+----------+
    Output:
    +------------+--------+------------+
    | day        | emp_id | total_time |
    +------------+--------+------------+
    | 2020-11-28 | 1      | 173        |
    | 2020-11-28 | 2      | 30         |
    | 2020-12-03 | 1      | 41         |
    | 2020-12-09 | 2      | 27         |
    +------------+--------+------------+
    Explanation:
    Employee 1 has three events: two on day 2020-11-28 with a total of (32 - 4) + (200 - 55) = 173, and one on day 2020-12-03 with a total of (42 - 1) = 41.
    Employee 2 has two events: one on day 2020-11-28 with a total of (33 - 3) = 30, and one on day 2020-12-09 with a total of (74 - 47) = 27.

---

# Solution

## Intuition

We need to find out the total time spent by each employee on each day at the office. The total time for each entry is
calculated by subtracting the in_time from out_time. We can use the aggregation function SUM to get the total time for
each day per employee.

## Approach

We will use the SQL GROUP BY statement to group the data by the day and the employee. This way, we get individual
records for each day per employee. Once we have this, we will apply the SUM function on the (out_time - in_time)
calculation to get the total time spent by an employee on a particular day.

## Complexity

- Time complexity:
  The time complexity is O(n), where n is the number of entries in the Employees table, because in the worst case scenario we have to go through all the entries.

- Space complexity:
  The space complexity is O(1), because we only use a constant amount of space to store our SQL query, and the database handles the rest.

---

# Code

Here are the solution and test files:

- [Solution Code](./solution.sql)
- [Test Code](./solution_test.go)
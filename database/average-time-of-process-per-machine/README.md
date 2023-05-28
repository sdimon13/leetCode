# 1661. Average Time of Process per Machine

Level: `Easy`

https://leetcode.com/problems/average-time-of-process-per-machine/

---

# Description

Table: `Activity`

    +----------------+---------+
    | Column Name    | Type    |
    +----------------+---------+
    | machine_id     | int     |
    | process_id     | int     |
    | activity_type  | enum    |
    | timestamp      | float   |
    +----------------+---------+
    The table shows the user activities for a factory website.
    (machine_id, process_id, activity_type) is the primary key of this table.
    machine_id is the ID of a machine.
    process_id is the ID of a process running on the machine with ID machine_id.
    activity_type is an ENUM of type ('start', 'end').
    timestamp is a float representing the current time in seconds.
    'start' means the machine starts the process at the given timestamp and 'end' means the machine ends the process at the given timestamp.
    The 'start' timestamp will always be before the 'end' timestamp for every (machine_id, process_id) pair.

There is a factory website that has several machines each running the **same number of processes**. Write an SQL query
to
find the **average time** each machine takes to complete a process.

The time to complete a process is the `'end' timestamp` minus the `'start' timestamp`. The average time is calculated by
the
total time to complete every process on the machine divided by the number of processes that were run.

The resulting table should have the `machine_id` along with the **average time** as `processing_time`, which should be *
*rounded
to 3 decimal places**.

Return the result table in **any order**.

The query result format is in the following example.

## Example 1:

    Input:
    Activity table:
    +------------+------------+---------------+-----------+
    | machine_id | process_id | activity_type | timestamp |
    +------------+------------+---------------+-----------+
    | 0          | 0          | start         | 0.712     |
    | 0          | 0          | end           | 1.520     |
    | 0          | 1          | start         | 3.140     |
    | 0          | 1          | end           | 4.120     |
    | 1          | 0          | start         | 0.550     |
    | 1          | 0          | end           | 1.550     |
    | 1          | 1          | start         | 0.430     |
    | 1          | 1          | end           | 1.420     |
    | 2          | 0          | start         | 4.100     |
    | 2          | 0          | end           | 4.512     |
    | 2          | 1          | start         | 2.500     |
    | 2          | 1          | end           | 5.000     |
    +------------+------------+---------------+-----------+
    Output:
    +------------+-----------------+
    | machine_id | processing_time |
    +------------+-----------------+
    | 0          | 0.894           |
    | 1          | 0.995           |
    | 2          | 1.456           |
    +------------+-----------------+
    Explanation:
    There are 3 machines running 2 processes each.
    Machine 0's average time is ((1.520 - 0.712) + (4.120 - 3.140)) / 2 = 0.894
    Machine 1's average time is ((1.550 - 0.550) + (1.420 - 0.430)) / 2 = 0.995
    Machine 2's average time is ((4.512 - 4.100) + (5.000 - 2.500)) / 2 = 1.456

---

# Solution

## Intuition

The problem requires us to calculate the average time taken by each machine to process each process. We can solve this
problem by using SQL joins and group by operations. We need to subtract the 'start' timestamp from the 'end' timestamp
for each process to get the processing time and then calculate the average processing time for each machine.

## Approach

1. To achieve this, we can first join the 'Activity' table with itself, based on matching 'machine_id' and 'process_id'
   columns, and where one activity is 'start' and the other is 'end'.
2. This will allow us to have both the 'start' and 'end' timestamps for each process in one row.
3. Then we can calculate the processing time for each process by subtracting the 'start' timestamp from the 'end'
   timestamp.
4. After that, we can group the rows by 'machine_id' and calculate the average processing time for each machine.

## Complexity

- Time complexity:
  The time complexity of this approach depends on the SQL engine, but it can generally be assumed to be O(n^2), where n
  is the number of records in the 'Activity' table. This is because we are performing a self join on the 'Activity'
  table.

- Space complexity:
  The space complexity also depends on the SQL engine and can be assumed to be O(n), where n is the number of records in
  the 'Activity' table.

---

# Code

Here are the solution and test files:

- [Solution Code](./solution.sql)
- [Test Code](./solution_test.go)
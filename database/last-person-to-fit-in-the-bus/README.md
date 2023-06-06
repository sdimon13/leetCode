# 1204. Last Person to Fit in the Bus

Level: `Medium`

https://leetcode.com/problems/last-person-to-fit-in-the-bus/

---

# Description

Table: `Queue`

    +-------------+---------+
    | Column Name | Type    |
    +-------------+---------+
    | person_id   | int     |
    | person_name | varchar |
    | weight      | int     |
    | turn        | int     |
    +-------------+---------+
    person_id is the primary key column for this table.
    This table has the information about all people waiting for a bus.
    The person_id and turn columns will contain all numbers from 1 to n, where n is the number of rows in the table.
    turn determines the order of which the people will board the bus, where turn=1 denotes the first person to board and turn=n denotes the last person to board.
    weight is the weight of the person in kilograms.

There is a queue of people waiting to board a bus. However, the bus has a weight limit of `1000` **kilograms**, so there
may be some people who cannot board.

Write an SQL query to find the `person_name` of the **last person** that can fit on the bus without exceeding the weight
limit. The test cases are generated such that the first person does not exceed the weight limit.

The query result format is in the following example.

## Example 1:

    Input:
    Queue table:
    +-----------+-------------+--------+------+
    | person_id | person_name | weight | turn |
    +-----------+-------------+--------+------+
    | 5         | Alice       | 250    | 1    |
    | 4         | Bob         | 175    | 5    |
    | 3         | Alex        | 350    | 2    |
    | 6         | John Cena   | 400    | 3    |
    | 1         | Winston     | 500    | 6    |
    | 2         | Marie       | 200    | 4    |
    +-----------+-------------+--------+------+
    Output:
    +-------------+
    | person_name |
    +-------------+
    | John Cena   |
    +-------------+
    Explanation: The folowing table is ordered by the turn for simplicity.
    +------+----+-----------+--------+--------------+
    | Turn | ID | Name      | Weight | Total Weight |
    +------+----+-----------+--------+--------------+
    | 1    | 5  | Alice     | 250    | 250          |
    | 2    | 3  | Alex      | 350    | 600          |
    | 3    | 6  | John Cena | 400    | 1000         | (last person to board)
    | 4    | 2  | Marie     | 200    | 1200         | (cannot board)
    | 5    | 4  | Bob       | 175    | ___          |
    | 6    | 1  | Winston   | 500    | ___          |
    +------+----+-----------+--------+--------------+

---

# Solution

## Intuition

This problem can be solved by using a cumulative sum function. In SQL, this can be achieved by using the **SUM**
function with the **OVER** clause, which allows for windows of rows to be defined for the computation of a function's
result. By ordering the sum over the 'turn' column, we can calculate a running total of the weight as people board the
bus.

## Approach

We need to return the name of the last person who can board the bus without exceeding the weight limit of 1000 kg. We
start by computing the running total of the weight in the order of people's turns to board the bus. This running total
will represent the total weight of the bus after each person boards.

Next, from this intermediate table, we filter out the rows where the total weight exceeds the limit of 1000 kg, as these
represent the people who couldn't board the bus. The remaining rows will have the people who could board the bus, with
their total weights in ascending order of their turns to board.

Finally, to get the name of the last person who boarded, we order the remaining rows in descending order of their total
weights and pick the first row. This will give us the person who boarded last without causing the total weight to exceed
the limit.

## Complexity

- Time complexity:
  The time complexity is O(nlogn) due to the sorting operation involved in ordering the rows by turn and then by total weight.

- Space complexity:
  The space complexity is O(n) because an intermediate table with n rows is generated.

---

# Code

Here are the solution and test files:

- [Solution Code](./solution.sql)
- [Test Code](./solution_test.go)
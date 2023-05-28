# 1280. Students and Examinations

Level: `Easy`

https://leetcode.com/problems/students-and-examinations/

---

# Description

Table: `Students`

    +---------------+---------+
    | Column Name   | Type    |
    +---------------+---------+
    | student_id    | int     |
    | student_name  | varchar |
    +---------------+---------+
    student_id is the primary key for this table.
    Each row of this table contains the ID and the name of one student in the school.

Table: `Subjects`

    +--------------+---------+
    | Column Name  | Type    |
    +--------------+---------+
    | subject_name | varchar |
    +--------------+---------+
    subject_name is the primary key for this table.
    Each row of this table contains the name of one subject in the school.

Table: `Examinations`

    +--------------+---------+
    | Column Name  | Type    |
    +--------------+---------+
    | student_id   | int     |
    | subject_name | varchar |
    +--------------+---------+
    There is no primary key for this table. It may contain duplicates.
    Each student from the Students table takes every course from the Subjects table.
    Each row of this table indicates that a student with ID student_id attended the exam of subject_name.

Write an SQL query to find the number of times each student attended each exam.

Return the result table ordered by `student_id` and `subject_name`.

The query result format is in the following example.

## Example 1:

    Input:
    Students table:
    +------------+--------------+
    | student_id | student_name |
    +------------+--------------+
    | 1          | Alice        |
    | 2          | Bob          |
    | 13         | John         |
    | 6          | Alex         |
    +------------+--------------+
    Subjects table:
    +--------------+
    | subject_name |
    +--------------+
    | Math         |
    | Physics      |
    | Programming  |
    +--------------+
    Examinations table:
    +------------+--------------+
    | student_id | subject_name |
    +------------+--------------+
    | 1          | Math         |
    | 1          | Physics      |
    | 1          | Programming  |
    | 2          | Programming  |
    | 1          | Physics      |
    | 1          | Math         |
    | 13         | Math         |
    | 13         | Programming  |
    | 13         | Physics      |
    | 2          | Math         |
    | 1          | Math         |
    +------------+--------------+
    Output:
    +------------+--------------+--------------+----------------+
    | student_id | student_name | subject_name | attended_exams |
    +------------+--------------+--------------+----------------+
    | 1          | Alice        | Math         | 3              |
    | 1          | Alice        | Physics      | 2              |
    | 1          | Alice        | Programming  | 1              |
    | 2          | Bob          | Math         | 1              |
    | 2          | Bob          | Physics      | 0              |
    | 2          | Bob          | Programming  | 1              |
    | 6          | Alex         | Math         | 0              |
    | 6          | Alex         | Physics      | 0              |
    | 6          | Alex         | Programming  | 0              |
    | 13         | John         | Math         | 1              |
    | 13         | John         | Physics      | 1              |
    | 13         | John         | Programming  | 1              |
    +------------+--------------+--------------+----------------+
    Explanation:
    The result table should contain all students and all subjects.
    Alice attended the Math exam 3 times, the Physics exam 2 times, and the Programming exam 1 time.
    Bob attended the Math exam 1 time, the Programming exam 1 time, and did not attend the Physics exam.
    Alex did not attend any exams.
    John attended the Math exam 1 time, the Physics exam 1 time, and the Programming exam 1 time.

---

# Solution

## Intuition

In this problem, we are asked to find the number of times each student attended each exam. This information can be
obtained by joining the tables Students, Subjects, and Examinations. The Students and Subjects tables contain the unique
id and name for each student and subject respectively. The Examinations table contains records of each student's
participation in each subject's exam. Since a student can attend an exam multiple times, the Examinations table can have
duplicate entries.

## Approach

Our approach will be to form a combined table where each row represents a student and a subject. To do this, we will
perform a **CROSS JOIN** of the Students and Subjects tables. This operation will generate a Cartesian product of the
tables, meaning it will pair each student with each subject.

Afterward, we will **LEFT JOIN** this table with the Examinations table on the conditions that the student_id and
subject_name match in both tables. The use of a **LEFT JOIN** ensures that we retain all combinations of students and
subjects, even if they do not have a corresponding record in the Examinations table.

Finally, we will use the **COUNT** function to determine the number of times each student took an exam for each subject.
Note that this count will include duplicates from the Examinations table, which represent multiple exam attendances. We
also make sure to group by student_id and subject_name to get the count for each distinct student-subject pair. The
result is ordered by student_id and subject_name as per the problem statement.

## Complexity

- Time complexity:
  The time complexity of this query is primarily determined by the JOIN operations. If there are n students, m subjects,
  and p examination records, the time complexity is approximately O(n∗m+n∗m∗p).

- Space complexity:
  The space complexity is determined by the space needed to store the output. In the worst case, we need space for n*m
  records, so the space complexity is O(n∗m).

---

# Code

Here are the solution and test files:

- [Solution Code](./solution.sql)
- [Test Code](./solution_test.go)
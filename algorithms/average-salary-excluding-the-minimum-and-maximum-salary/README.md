# 1491. Average Salary Excluding the Minimum and Maximum Salary
Level: `Easy`

https://leetcode.com/problems/average-salary-excluding-the-minimum-and-maximum-salary/

---

# Description

You are given an array of **unique** integers `salary` where `salary[i]` is the salary of the `i`<sup>`th`</sup> employee.

Return the average salary of employees excluding the minimum and maximum salary. Answers within `10`<sup>`-5`</sup> of the actual answer will be accepted.

## Example 1:

    Input: salary = [4000,3000,1000,2000]
    Output: 2500.00000
    Explanation: Minimum salary and maximum salary are 1000 and 4000 respectively.
    Average salary excluding minimum and maximum salary is (2000+3000) / 2 = 2500

## Example 2:

    Input: salary = [1000,2000,3000]
    Output: 2000.00000
    Explanation: Minimum salary and maximum salary are 1000 and 3000 respectively.
    Average salary excluding minimum and maximum salary is (2000) / 1 = 2000

## Constraints:

 - `3 <= salary.length <= 100`

 - `1000 <= salary[i] <= 10`<sup>`6`</sup>

 - All the integers of `salary` are **unique**.

---

# Solution

## Intuition
Our initial thought on how to solve this problem is to compute the sum of all salaries while also keeping track of the minimum and maximum salaries. Once we have these values, we can subtract the minimum and maximum salaries from the total sum and divide by the number of salaries minus two (since we've excluded the minimum and maximum salaries). This approach doesn't require any additional libraries and is efficient in terms of time complexity.

## Approach
Our approach involves a single pass through the array of salaries. For each salary:

1. We check if it's less than our current minimum salary, and if so, we update our minimum salary.
2. We check if it's more than our current maximum salary, and if so, we update our maximum salary.
3. We add the salary to our running total sum.

After going through all salaries, we subtract the minimum and maximum salaries from the total sum and divide by the number of salaries minus two.

## Complexity
- Time complexity:
  O(n), where n is the length of the salary array. We are only performing a single pass through the array.

- Space complexity:
  O(1), where we only need a constant amount of space to store the sum, minimum salary, and maximum salary. The space requirement does not grow with the size of the input.

---

# Code
Here are the solution and test files:
- [Solution Code](./solution.go)
- [Test Code](./solution_test.go)
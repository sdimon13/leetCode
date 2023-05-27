# [LeetCode Solutions](https://github.com/sdimon13/leetCode)

## Description

This repository contains my solutions to problems from [LeetCode](https://leetcode.com/). I am using Mysql for solving
the
problems, and each solution includes comments explaining my approach.

To verify the correctness of database solutions, I provide a Docker container with MySQL. It will allow users to run
database queries with the same conditions as on LeetCode.

## How to Use

1. You can browse the solutions for learning purposes or as inspiration for your own solutions. Each solution includes a
   link to the corresponding problem on LeetCode.

2. Follow these steps to test your database solutions:

- Install Docker on your machine. You can follow the installation steps from the
  official [Docker documentation](https://docs.docker.com/get-docker/).

- Run the Docker container with MySQL by using the command below in your terminal:

`docker run --name test-mysql -e MYSQL_ROOT_PASSWORD=my-secret-pw -e MYSQL_DATABASE=testdb -p 3306:3306 -d mysql:latest`

- Now you can use any MySQL client to connect to the database using the host localhost, the port 3306, and the root
  password my-secret-pw.

- You can run queries present in the solutions directly on this database.

- **Remember**: After finishing your work, stop the Docker container by using the command below in your terminal:

`docker stop test-mysql`

- If you need to start the container again, use the following command:

`docker start test-mysql`

## Contributing

If you would like to contribute to this repository, please first open an issue or join the discussion of existing ones.

## Database Solutions

|  #   | Title                                                                                                                   |                      Solution                       | Difficulty |
|:----:|:------------------------------------------------------------------------------------------------------------------------|:---------------------------------------------------:|:----------:|
| 0176 | [Second Highest Salary](https://leetcode.com/problems/second-highest-salary/)                                           |           [MySql](second-highest-salary)            |   Medium   |
| 0181 | [Employees Earning More Than Their Managers](https://leetcode.com/problems/employees-earning-more-than-their-managers/) | [MySql](employees-earning-more-than-their-managers) |    Easy    |
| 0182 | [Duplicate Emails](https://leetcode.com/problems/duplicate-emails/)                                                     |              [MySql](duplicate-emails)              |    Easy    |
| 0183 | [Customers Who Never Order](https://leetcode.com/problems/customers-who-never-order/)                                   |         [MySql](customers-who-never-order)          |    Easy    |
| 0184 | [Department Highest Salary](https://leetcode.com/problems/department-highest-salary/)                                   |         [MySql](department-highest-salary)          |   Medium   |
| 0185 | [Department Top Three Salaries](https://leetcode.com/problems/department-top-three-salaries/)                           |       [MySql](department-top-three-salaries)        |    Hard    |
| 0196 | [Delete Duplicate Emails](https://leetcode.com/problems/delete-duplicate-emails/)                                       |          [MySql](delete-duplicate-emails)           |    Easy    |
| 0584 | [Find Customer Referee](https://leetcode.com/problems/find-customer-referee/)                                           |           [MySql](find-customer-referee)            |    Easy    |
| 0595 | [Big Countries](https://leetcode.com/problems/big-countries/)                                                           |               [MySql](big-countries)                |    Easy    |
| 0627 | [Swap Salary](https://leetcode.com/problems/swap-salary/)                                                               |                [MySql](swap-salary)                 |    Easy    |
| 1667 | [Fix Names in a Table](https://leetcode.com/problems/fix-names-in-a-table/)                                             |            [MySql](fix-names-in-a-table)            |    Easy    |
| 1757 | [Recyclable and Low Fat Products](https://leetcode.com/problems/recyclable-and-low-fat-products/)                       |      [MySql](recyclable-and-low-fat-products)       |    Easy    |
| 1873 | [Calculate Special Bonus](https://leetcode.com/problems/calculate-special-bonus/)                                       |          [MySql](calculate-special-bonus)           |    Easy    |

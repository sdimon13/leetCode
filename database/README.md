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

|  #   | Title                                                                                                                                               |                             Solution                              | Difficulty |
|:----:|:----------------------------------------------------------------------------------------------------------------------------------------------------|:-----------------------------------------------------------------:|:----------:|
| 0175 | [Combine Two Tables](https://leetcode.com/problems/combine-two-tables/)                                                                             |                    [MySql](combine-two-tables)                    |    Easy    |
| 0176 | [Second Highest Salary](https://leetcode.com/problems/second-highest-salary/)                                                                       |                  [MySql](second-highest-salary)                   |   Medium   |
| 0181 | [Employees Earning More Than Their Managers](https://leetcode.com/problems/employees-earning-more-than-their-managers/)                             |        [MySql](employees-earning-more-than-their-managers)        |    Easy    |
| 0182 | [Duplicate Emails](https://leetcode.com/problems/duplicate-emails/)                                                                                 |                     [MySql](duplicate-emails)                     |    Easy    |
| 0183 | [Customers Who Never Order](https://leetcode.com/problems/customers-who-never-order/)                                                               |                [MySql](customers-who-never-order)                 |    Easy    |
| 0184 | [Department Highest Salary](https://leetcode.com/problems/department-highest-salary/)                                                               |                [MySql](department-highest-salary)                 |   Medium   |
| 0185 | [Department Top Three Salaries](https://leetcode.com/problems/department-top-three-salaries/)                                                       |              [MySql](department-top-three-salaries)               |    Hard    |
| 0196 | [Delete Duplicate Emails](https://leetcode.com/problems/delete-duplicate-emails/)                                                                   |                 [MySql](delete-duplicate-emails)                  |    Easy    |
| 0197 | [Rising Temperature](https://leetcode.com/problems/rising-temperature/)                                                                             |                    [MySql](rising-temperature)                    |    Easy    |
| 0511 | [Game Play Analysis I](https://leetcode.com/problems/game-play-analysis-i/)                                                                         |                   [MySql](game-play-analysis-i)                   |    Easy    |
| 0584 | [Find Customer Referee](https://leetcode.com/problems/find-customer-referee/)                                                                       |                  [MySql](find-customer-referee)                   |    Easy    |
| 0586 | [Customer Placing the Largest Number of Orders](https://leetcode.com/problems/customer-placing-the-largest-number-of-orders/)                       |      [MySql](customer-placing-the-largest-number-of-orders)       |    Easy    |
| 0595 | [Big Countries](https://leetcode.com/problems/big-countries/)                                                                                       |                      [MySql](big-countries)                       |    Easy    |
| 0607 | [Sales Person](https://leetcode.com/problems/sales-person/)                                                                                         |                       [MySql](sales-person)                       |    Easy    |
| 0608 | [Tree Node](https://leetcode.com/problems/tree-node/)                                                                                               |                        [MySql](tree-node)                         |   Medium   |
| 0620 | [Not Boring Movies](https://leetcode.com/problems/not-boring-movies/)                                                                               |                    [MySql](not-boring-movies)                     |    Easy    |
| 0627 | [Swap Salary](https://leetcode.com/problems/swap-salary/)                                                                                           |                       [MySql](swap-salary)                        |    Easy    |
| 1050 | [Actors and Directors Who Cooperated At Least Three Times](https://leetcode.com/problems/actors-and-directors-who-cooperated-at-least-three-times/) | [MySql](actors-and-directors-who-cooperated-at-least-three-times) |    Easy    |
| 1084 | [Sales Analysis III](https://leetcode.com/problems/sales-analysis-iii/)                                                                             |                    [MySql](sales-analysis-iii)                    |    Easy    |
| 1141 | [User Activity for the Past 30 Days I](https://leetcode.com/problems/user-activity-for-the-past-30-days-i/)                                         |           [MySql](user-activity-for-the-past-30-days-i)           |    Easy    |
| 1148 | [Article Views I](https://leetcode.com/problems/article-views-i/)                                                                                   |                     [MySql](article-views-i)                      |    Easy    |
| 1158 | [Market Analysis I](https://leetcode.com/problems/market-analysis-i/)                                                                               |                    [MySql](market-analysis-i)                     |   Medium   |
| 1164 | [Product Price at a Given Date](https://leetcode.com/problems/product-price-at-a-given-date/)                                                       |              [MySql](product-price-at-a-given-date)               |   Medium   |
| 1280 | [Students and Examinations](https://leetcode.com/problems/students-and-examinations/)                                                               |                [MySql](students-and-examinations)                 |    Easy    |
| 1327 | [List the Products Ordered in a Period](https://leetcode.com/problems/list-the-products-ordered-in-a-period/)                                       |          [MySql](list-the-products-ordered-in-a-period)           |    Easy    |
| 1378 | [Replace Employee ID With The Unique Identifier](https://leetcode.com/problems/replace-employee-id-with-the-unique-identifier/)                     |      [MySql](replace-employee-id-with-the-unique-identifier)      |    Easy    |
| 1393 | [Capital Gain/Loss](https://leetcode.com/problems/capital-gainloss/)                                                                                |                     [MySql](capital-gainloss)                     |   Medium   |
| 1407 | [Top Travellers](https://leetcode.com/problems/top-travellers/)                                                                                     |                      [MySql](top-travellers)                      |    Easy    |
| 1484 | [Group Sold Products By The Date](https://leetcode.com/problems/group-sold-products-by-the-date/)                                                   |             [MySql](group-sold-products-by-the-date)              |    Easy    |
| 1527 | [Patients With a Condition](https://leetcode.com/problems/patients-with-a-condition/)                                                               |                [MySql](patients-with-a-condition)                 |    Easy    |
| 1581 | [Customer Who Visited but Did Not Make Any Transactions](https://leetcode.com/problems/customer-who-visited-but-did-not-make-any-transactions/)     |  [MySql](customer-who-visited-but-did-not-make-any-transactions)  |    Easy    |
| 1587 | [Bank Account Summary II](https://leetcode.com/problems/bank-account-summary-ii/)                                                                   |                 [MySql](bank-account-summary-ii)                  |    Easy    |
| 1661 | [Average Time of Process per Machine](https://leetcode.com/problems/average-time-of-process-per-machine/)                                           |           [MySql](average-time-of-process-per-machine)            |    Easy    |
| 1667 | [Fix Names in a Table](https://leetcode.com/problems/fix-names-in-a-table/)                                                                         |                   [MySql](fix-names-in-a-table)                   |    Easy    |
| 1683 | [Invalid Tweets](https://leetcode.com/problems/invalid-tweets/)                                                                                     |                      [MySql](invalid-tweets)                      |    Easy    |
| 1693 | [Daily Leads and Partners](https://leetcode.com/problems/daily-leads-and-partners/)                                                                 |                 [MySql](daily-leads-and-partners)                 |    Easy    |
| 1729 | [Find Followers Count](https://leetcode.com/problems/find-followers-count/)                                                                         |                   [MySql](find-followers-count)                   |    Easy    |
| 1741 | [Find Total Time Spent by Each Employee](https://leetcode.com/problems/find-total-time-spent-by-each-employee/)                                     |          [MySql](find-total-time-spent-by-each-employee)          |    Easy    |
| 1757 | [Recyclable and Low Fat Products](https://leetcode.com/problems/recyclable-and-low-fat-products/)                                                   |             [MySql](recyclable-and-low-fat-products)              |    Easy    |
| 1795 | [Rearrange Products Table](https://leetcode.com/problems/rearrange-products-table/)                                                                 |                 [MySql](rearrange-products-table)                 |    Easy    |
| 1873 | [Calculate Special Bonus](https://leetcode.com/problems/calculate-special-bonus/)                                                                   |                 [MySql](calculate-special-bonus)                  |    Easy    |
| 1890 | [The Latest Login in 2020](https://leetcode.com/problems/the-latest-login-in-2020/)                                                                 |                 [MySql](the-latest-login-in-2020)                 |    Easy    |
| 1965 | [Employees With Missing Information](https://leetcode.com/problems/employees-with-missing-information/)                                             |            [MySql](employees-with-missing-information)            |    Easy    |
| 1978 | [Employees Whose Manager Left the Company](https://leetcode.com/problems/employees-whose-manager-left-the-company/)                                 |         [MySql](employees-whose-manager-left-the-company)         |    Easy    |

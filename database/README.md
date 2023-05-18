# [LeetCode Solutions](https://github.com/sdimon13/leetCode)

## Description

This repository contains my solutions to problems from [LeetCode](https://leetcode.com/). I am using Mysql for solving the
problems, and each solution includes comments explaining my approach.

To verify the correctness of database solutions, I provide a Docker container with MySQL. It will allow users to run database queries with the same conditions as on LeetCode.

## How to Use

1. You can browse the solutions for learning purposes or as inspiration for your own solutions. Each solution includes a link to the corresponding problem on LeetCode.

2. Follow these steps to test your database solutions:

- Install Docker on your machine. You can follow the installation steps from the official [Docker documentation](https://docs.docker.com/get-docker/).

- Run the Docker container with MySQL by using the command below in your terminal:

`docker run --name test-mysql -e MYSQL_ROOT_PASSWORD=my-secret-pw -e MYSQL_DATABASE=testdb -p 3306:3306 -d mysql:latest`

- Now you can use any MySQL client to connect to the database using the host localhost, the port 3306, and the root password my-secret-pw.

- You can run queries present in the solutions directly on this database.

- **Remember**: After finishing your work, stop the Docker container by using the command below in your terminal:

`docker stop test-mysql`

- If you need to start the container again, use the following command:

`docker start test-mysql`

## Contributing

If you would like to contribute to this repository, please first open an issue or join the discussion of existing ones.

## Database Solutions

|  #   | Title                                                                             |            Solution            | Difficulty |
|:----:|:----------------------------------------------------------------------------------|:------------------------------:|:----------:|
| 0196 | [Delete Duplicate Emails](https://leetcode.com/problems/delete-duplicate-emails/) | [GO](delete-duplicate-emails)  |    Easy    |

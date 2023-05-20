select e.name as Employee
from Employee as e
         left join Employee as m on e.managerId = m.id
where e.salary > m.salary
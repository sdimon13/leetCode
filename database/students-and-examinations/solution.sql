SELECT st.student_id, st.student_name, s.subject_name, COUNT(e.subject_name) AS attended_exams
FROM Students AS st
         JOIN Subjects AS s
         LEFT JOIN Examinations AS e
                   ON st.student_id = e.student_id
                       AND s.subject_name = e.subject_name
GROUP BY st.student_id, s.subject_name
ORDER BY student_id ASC, subject_name ASC
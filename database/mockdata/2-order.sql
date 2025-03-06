INSERT INTO public.order (id, board_id, name, due_date, quantity, updated_by, updated_dt, created_by, created_dt)
OVERRIDING SYSTEM VALUE
VALUES 
(1, 1, 'Task 1', '2025-03-10', 5, NULL, NULL, 'admin', '2025-03-05 09:00:00'),
(2, 1, 'Task 2', '2025-03-12', 3, NULL, NULL, 'admin', '2025-03-05 09:05:00'),
(3, 2, 'Task 3', '2025-03-15', 8, NULL, NULL, 'admin', '2025-03-05 09:10:00'),
(4, 3, 'Task 4', '2025-03-18', 6, NULL, NULL, 'admin', '2025-03-05 09:15:00'),
(5, 2, 'Task 5', '2025-03-20', 4, NULL, NULL, 'admin', '2025-03-05 09:20:00'),
(6, 3, 'Task 6', '2025-03-25', 2, NULL, NULL, 'admin', '2025-03-05 09:25:00'),
(7, 1, 'Task 7', '2025-03-28', 7, NULL, NULL, 'admin', '2025-03-05 09:30:00'),
(8, 4, 'Task 8', '2025-03-30', 10, NULL, NULL, 'admin', '2025-03-05 09:35:00'),
(9, 4, 'Task 9', '2025-04-02', 5, NULL, NULL, 'admin', '2025-03-05 09:40:00'),
(10, 2, 'Task 10', '2025-04-05', 3, NULL, NULL, 'admin', '2025-03-05 09:45:00');

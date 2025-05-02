-- Write your migrate up statements here

INSERT INTO public.board 
(
    board_id
    ,name
    ,updated_by
    ,updated_dt
    ,created_by
    ,created_dt
)
OVERRIDING SYSTEM VALUE
VALUES 
(1, 'Board 1', NULL, NULL, 'admin', '2025-03-05 09:00:00'),
(2, 'Board 2', NULL, NULL, 'admin', '2025-03-05 09:05:00'),
(3, 'Board 3', NULL, NULL, 'admin', '2025-03-05 09:10:00'),
(4, 'Board 4', NULL, NULL, 'admin', '2025-03-05 09:15:00'),
(5, 'Board 5', NULL, NULL, 'admin', '2025-03-05 09:20:00'),
(6, 'Board 6', NULL, NULL, 'admin', '2025-03-05 09:25:00'),
(7, 'Board 7', NULL, NULL, 'admin', '2025-03-05 09:30:00'),
(8, 'Board 8', NULL, NULL, 'admin', '2025-03-05 09:35:00'),
(9, 'Board 9', NULL, NULL, 'admin', '2025-03-05 09:40:00'),
(10, 'Board 10', NULL, NULL, 'admin', '2025-03-05 09:45:00'),
(11, 'Board 11', NULL, NULL, 'admin', '2025-03-05 09:50:00'),
(12, 'Board 12', NULL, NULL, 'admin', '2025-03-05 09:55:00'),
(13, 'Board 13', NULL, NULL, 'admin', '2025-03-05 10:00:00'),
(14, 'Board 14', NULL, NULL, 'admin', '2025-03-05 10:05:00'),
(15, 'Board 15', NULL, NULL, 'admin', '2025-03-05 10:10:00'),
(16, 'Board 16', NULL, NULL, 'admin', '2025-03-05 10:15:00'),
(17, 'Board 17', NULL, NULL, 'admin', '2025-03-05 10:20:00'),
(18, 'Board 18', NULL, NULL, 'admin', '2025-03-05 10:25:00'),
(19, 'Board 19', NULL, NULL, 'admin', '2025-03-05 10:30:00'),
(20, 'Board 20', NULL, NULL, 'admin', '2025-03-05 10:35:0020');

INSERT INTO public.job 
(
    job_id
    ,name
    ,updated_by
    ,updated_dt
    ,created_by
    ,created_dt
)
OVERRIDING SYSTEM VALUE
VALUES 
(1, 'Job 1', NULL, NULL, 'admin', '2025-03-05 09:00:00'),
(2, 'Job 2', NULL, NULL, 'admin', '2025-03-05 09:05:00'),
(3, 'Job 3', NULL, NULL, 'admin', '2025-03-05 09:10:00'),
(4, 'Job 4', NULL, NULL, 'admin', '2025-03-05 09:15:00'),
(5, 'Job 5', NULL, NULL, 'admin', '2025-03-05 09:20:00'),
(6, 'Job 6', NULL, NULL, 'admin', '2025-03-05 09:25:00'),
(7, 'Job 7', NULL, NULL, 'admin', '2025-03-05 09:30:00'),
(8, 'Job 8', NULL, NULL, 'admin', '2025-03-05 09:35:00'),
(9, 'Job 9', NULL, NULL, 'admin', '2025-03-05 09:40:00'),
(10, 'Job 10', NULL, NULL, 'admin', '2025-03-05 09:45:00'),
(11, 'Job 11', NULL, NULL, 'admin', '2025-03-05 09:50:00'),
(12, 'Job 12', NULL, NULL, 'admin', '2025-03-05 09:55:00'),
(13, 'Job 13', NULL, NULL, 'admin', '2025-03-05 10:00:00'),
(14, 'Job 14', NULL, NULL, 'admin', '2025-03-05 10:05:00'),
(15, 'Job 15', NULL, NULL, 'admin', '2025-03-05 10:10:00'),
(16, 'Job 16', NULL, NULL, 'admin', '2025-03-05 10:15:00'),
(17, 'Job 17', NULL, NULL, 'admin', '2025-03-05 10:20:00'),
(18, 'Job 18', NULL, NULL, 'admin', '2025-03-05 10:25:00'),
(19, 'Job 19', NULL, NULL, 'admin', '2025-03-05 10:30:00'),
(20, 'Job 20', NULL, NULL, 'admin', '2025-03-05 10:35:00');

INSERT INTO public.line 
(
    line_id
    ,name
    ,updated_by
    ,updated_dt
    ,created_by
    ,created_dt
)
OVERRIDING SYSTEM VALUE
VALUES 
(1, 'Line 1', NULL, NULL, 'admin', '2025-03-05 09:00:00'),
(2, 'Line 2', NULL, NULL, 'admin', '2025-03-05 09:05:00'),
(3, 'Line 3', NULL, NULL, 'admin', '2025-03-05 09:10:00'),
(4, 'Line 4', NULL, NULL, 'admin', '2025-03-05 09:15:00'),
(5, 'Line 5', NULL, NULL, 'admin', '2025-03-05 09:20:00'),
(6, 'Line 6', NULL, NULL, 'admin', '2025-03-05 09:25:00');

INSERT INTO public.package 
(
    package_id
    ,name
    ,updated_by
    ,updated_dt
    ,created_by
    ,created_dt
)
OVERRIDING SYSTEM VALUE
VALUES
(1, 'Package A', NULL, NULL, 'admin', '2025-03-05 09:00:00'),
(2, 'Package B', NULL, NULL, 'admin', '2025-03-05 09:05:00'),
(3, 'Package C', NULL, NULL, 'admin', '2025-03-05 09:10:00'),
(4, 'Package D', NULL, NULL, 'admin', '2025-03-05 09:15:00'),
(5, 'Package E', NULL, NULL, 'admin', '2025-03-05 09:20:00'),
(6, 'Package F', NULL, NULL, 'admin', '2025-03-05 09:25:00'),
(7, 'Package G', NULL, NULL, 'admin', '2025-03-05 09:30:00'),
(8, 'Package H', NULL, NULL, 'admin', '2025-03-05 09:35:00'),
(9, 'Package I', NULL, NULL, 'admin', '2025-03-05 09:40:00'),
(10, 'Package J', NULL, NULL, 'admin', '2025-03-05 09:45:00'),
(11, 'Package K', NULL, NULL, 'admin', '2025-03-05 09:50:00'),
(12, 'Package L', NULL, NULL, 'admin', '2025-03-05 09:55:00'),
(13, 'Package M', NULL, NULL, 'admin', '2025-03-05 10:00:00'),
(14, 'Package N', NULL, NULL, 'admin', '2025-03-05 10:05:00'),
(15, 'Package O', NULL, NULL, 'admin', '2025-03-05 10:10:00'),
(16, 'Package P', NULL, NULL, 'admin', '2025-03-05 10:15:00'),
(17, 'Package Q', NULL, NULL, 'admin', '2025-03-05 10:20:00'),
(18, 'Package R', NULL, NULL, 'admin', '2025-03-05 10:25:00'),
(19, 'Package S', NULL, NULL, 'admin', '2025-03-05 10:30:00'),
(20, 'Package T', NULL, NULL, 'admin', '2025-03-05 10:35:00');

INSERT INTO public.auth
(
    auth_id
    ,oauth 
    ,password
    ,email
    ,updated_by
    ,updated_dt
    ,created_by
    ,created_dt
)
OVERRIDING SYSTEM VALUE
VALUES
(1,  false, 'password', 'user1@example.com', NULL, NULL, 'admin', '2025-03-05 09:00:00'),
(2,  false, 'password', 'user2@example.com', NULL, NULL, 'admin', '2025-03-05 09:05:00'),
(3,  false, 'password', 'user3@example.com', NULL, NULL, 'admin', '2025-03-05 09:10:00'),
(4,  false, 'password', 'user4@example.com', NULL, NULL, 'admin', '2025-03-05 09:15:00'),
(5,  false, 'password', 'user5@example.com', NULL, NULL, 'admin', '2025-03-05 09:20:00'),
(6,  false, 'password', 'user6@example.com', NULL, NULL, 'admin', '2025-03-05 09:25:00'),
(7,  false, 'password', 'user7@example.com', NULL, NULL, 'admin', '2025-03-05 09:30:00'),
(8,  false, 'password', 'user8@example.com', NULL, NULL, 'admin', '2025-03-05 09:35:00'),
(9,  false, 'password', 'user9@example.com', NULL, NULL, 'admin', '2025-03-05 09:40:00'),
(10, false, 'password', 'user10@example.com', NULL, NULL, 'admin', '2025-03-05 09:45:00'),
(11, false, 'password', 'user11@example.com', NULL, NULL, 'admin', '2025-03-05 09:50:00'),
(12, false, 'password', 'user12@example.com', NULL, NULL, 'admin', '2025-03-05 09:55:00'),
(13, false, 'password', 'user13@example.com', NULL, NULL, 'admin', '2025-03-05 10:00:00'),
(14, false, 'password', 'user14@example.com', NULL, NULL, 'admin', '2025-03-05 10:05:00'),
(15, false, 'password', 'user15@example.com', NULL, NULL, 'admin', '2025-03-05 10:10:00'),
(16, false, 'password', 'user16@example.com', NULL, NULL, 'admin', '2025-03-05 10:15:00'),
(17, false, 'password', 'user17@example.com', NULL, NULL, 'admin', '2025-03-05 10:20:00'),
(18, false, 'password', 'user18@example.com', NULL, NULL, 'admin', '2025-03-05 10:25:00'),
(19, false, 'password', 'user19@example.com', NULL, NULL, 'admin', '2025-03-05 10:30:00'),
(20, false, 'password', 'user20@example.com', NULL, NULL, 'admin', '2025-03-05 10:35:00');

INSERT INTO public.user
(
    user_id
    ,auth_id
    ,username
    ,updated_by
    ,updated_dt
    ,created_by
    ,created_dt
)
OVERRIDING SYSTEM VALUE
VALUES
(1, 1,  'user1', NULL, NULL, 'admin', '2025-03-05 09:00:00'),
(2, 2,  'user2', NULL, NULL, 'admin', '2025-03-05 09:05:00'),
(3, 3,  'user3', NULL, NULL, 'admin', '2025-03-05 09:10:00'),
(4, 4,  'user4', NULL, NULL, 'admin', '2025-03-05 09:15:00'),
(5, 5,  'user5', NULL, NULL, 'admin', '2025-03-05 09:20:00'),
(6, 6,  'user6', NULL, NULL, 'admin', '2025-03-05 09:25:00'),
(7, 7,  'user7', NULL, NULL, 'admin', '2025-03-05 09:30:00'),
(8, 8,  'user8', NULL, NULL, 'admin', '2025-03-05 09:35:00'),
(9, 9,  'user9', NULL, NULL, 'admin', '2025-03-05 09:40:00'),
(10, 10, 'user10', NULL, NULL, 'admin', '2025-03-05 09:45:00'),
(11, 11, 'user11', NULL, NULL, 'admin', '2025-03-05 09:50:00'),
(12, 12, 'user12', NULL, NULL, 'admin', '2025-03-05 09:55:00'),
(13, 13, 'user13', NULL, NULL, 'admin', '2025-03-05 10:00:00'),
(14, 14, 'user14', NULL, NULL, 'admin', '2025-03-05 10:05:00'),
(15, 15, 'user15', NULL, NULL, 'admin', '2025-03-05 10:10:00'),
(16, 16, 'user16', NULL, NULL, 'admin', '2025-03-05 10:15:00'),
(17, 17, 'user17', NULL, NULL, 'admin', '2025-03-05 10:20:00'),
(18, 18, 'user18', NULL, NULL, 'admin', '2025-03-05 10:25:00'),
(19, 19, 'user19', NULL, NULL, 'admin', '2025-03-05 10:30:00'),
(20, 20, 'user20', NULL, NULL, 'admin', '2025-03-05 10:35:00');

INSERT INTO public.line_job 
(
    job_id
    ,line_id
)
VALUES
(1, 1),
(2, 1),
(3, 2),
(4, 3),
(5, 4);

INSERT INTO public.order 
(
    order_id
    ,board_id
    ,name
    ,due_date
    ,quantity
    ,updated_by
    ,updated_dt
    ,created_by
    ,created_dt
)
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

INSERT INTO public.machine 
(
    machine_id
    ,line_id
    ,name
    ,updated_by
    ,updated_dt
    ,created_by
    ,created_dt
)
OVERRIDING SYSTEM VALUE
VALUES
(1, 1, 'Machine 1', NULL, NULL, 'admin', '2025-03-05 09:00:00'),
(2, 2, 'Machine 2', NULL, NULL, 'admin', '2025-03-05 09:05:00'),
(3, 3, 'Machine 3', NULL, NULL, 'admin', '2025-03-05 09:10:00'),
(4, 4, 'Machine 4', NULL, NULL, 'admin', '2025-03-05 09:15:00'),
(5, 5, 'Machine 5', NULL, NULL, 'admin', '2025-03-05 09:20:00'),
(6, 6, 'Machine 6', NULL, NULL, 'admin', '2025-03-05 09:25:00'),
(7, 1, 'Machine 7', NULL, NULL, 'admin', '2025-03-05 09:30:00'),
(8, 2, 'Machine 8', NULL, NULL, 'admin', '2025-03-05 09:35:00'),
(9, 3, 'Machine 9', NULL, NULL, 'admin', '2025-03-05 09:40:00'),
(10, 4, 'Machine 10', NULL, NULL, 'admin', '2025-03-05 09:45:00'),
(11, 5, 'Machine 11', NULL, NULL, 'admin', '2025-03-05 09:50:00'),
(12, 6, 'Machine 12', NULL, NULL, 'admin', '2025-03-05 09:55:00'),
(13, 1, 'Machine 13', NULL, NULL, 'admin', '2025-03-05 10:00:00'),
(14, 2, 'Machine 14', NULL, NULL, 'admin', '2025-03-05 10:05:00'),
(15, 3, 'Machine 15', NULL, NULL, 'admin', '2025-03-05 10:10:00'),
(16, 4, 'Machine 16', NULL, NULL, 'admin', '2025-03-05 10:15:00'),
(17, 5, 'Machine 17', NULL, NULL, 'admin', '2025-03-05 10:20:00'),
(18, 6, 'Machine 18', NULL, NULL, 'admin', '2025-03-05 10:25:00'),
(19, 1, 'Machine 19', NULL, NULL, 'admin', '2025-03-05 10:30:00'),
(20, 2, 'Machine 20', NULL, NULL, 'admin', '2025-03-05 10:35:00');

INSERT INTO public.order_job 
(
    order_id
    ,job_id
)
VALUES
(1, 1),
(2, 1),
(3, 2),
(4, 3),
(5, 4);

INSERT INTO public.cart 
(
    cart_id
    ,machine_id
    ,name
    ,slots
    ,direction
    ,updated_by
    ,updated_dt
    ,created_by
    ,created_dt
)
OVERRIDING SYSTEM VALUE
VALUES 
(1, 1, 'Cart 1', 60, 'right', NULL, NULL, 'admin', '2025-03-05 09:00:00'),
(2, 1, 'Cart 2', 60, 'left', NULL, NULL, 'admin', '2025-03-05 09:05:00'),
(3, 2, 'Cart 3', 60, 'right', NULL, NULL, 'admin', '2025-03-05 09:10:00'),
(4, 2, 'Cart 4', 60, 'left', NULL, NULL, 'admin', '2025-03-05 09:15:00'),
(5, 3, 'Cart 5', 60, 'right', NULL, NULL, 'admin', '2025-03-05 09:20:00'),
(6, 3, 'Cart 6', 60, 'left', NULL, NULL, 'admin', '2025-03-05 09:25:00'),
(7, 4, 'Cart 7', 60, 'right', NULL, NULL, 'admin', '2025-03-05 09:30:00'),
(8, 4, 'Cart 8', 60, 'left', NULL, NULL, 'admin', '2025-03-05 09:35:00'),
(9, 5, 'Cart 9', 60, 'right', NULL, NULL, 'admin', '2025-03-05 09:40:00'),
(10, 5, 'Cart 10', 60, 'left', NULL, NULL, 'admin', '2025-03-05 09:45:00'),
(11, 6, 'Cart 11', 60, 'right', NULL, NULL, 'admin', '2025-03-05 09:50:00'),
(12, 6, 'Cart 12', 60, 'left', NULL, NULL, 'admin', '2025-03-05 09:55:00'),
(13, 7, 'Cart 13', 60, 'right', NULL, NULL, 'admin', '2025-03-05 10:00:00'),
(14, 7, 'Cart 14', 60, 'left', NULL, NULL, 'admin', '2025-03-05 10:05:00'),
(15, 8, 'Cart 15', 60, 'right', NULL, NULL, 'admin', '2025-03-05 10:10:00'),
(16, 8, 'Cart 16', 60, 'left', NULL, NULL, 'admin', '2025-03-05 10:15:00'),
(17, 9, 'Cart 17', 60, 'right', NULL, NULL, 'admin', '2025-03-05 10:20:00'),
(18, 9, 'Cart 18', 60, 'left', NULL, NULL, 'admin', '2025-03-05 10:25:00'),
(19, 10, 'Cart 19', 60, 'right', NULL, NULL, 'admin', '2025-03-05 10:30:00'),
(20, 10, 'Cart 20', 60, 'left', NULL, NULL, 'admin', '2025-03-05 10:35:00');

INSERT INTO public.feeder 
(
    feeder_id
    ,cart_id
    ,package_id
    ,name
    ,type
    ,updated_by
    ,updated_dt
    ,created_by
    ,created_dt
)
OVERRIDING SYSTEM VALUE
VALUES
(1, 1, 1, 'Feeder 1', 'Type A', NULL, NULL, 'admin', '2025-03-05 09:00:00'),
(2, 2, 2, 'Feeder 2', 'Type B', NULL, NULL, 'admin', '2025-03-05 09:05:00'),
(3, 3, 3, 'Feeder 3', 'Type C', NULL, NULL, 'admin', '2025-03-05 09:10:00'),
(4, 4, 4, 'Feeder 4', 'Type D', NULL, NULL, 'admin', '2025-03-05 09:15:00'),
(5, 5, 5, 'Feeder 5', 'Type E', NULL, NULL, 'admin', '2025-03-05 09:20:00'),
(6, 6, 6, 'Feeder 6', 'Type F', NULL, NULL, 'admin', '2025-03-05 09:25:00'),
(7, 7, 7, 'Feeder 7', 'Type G', NULL, NULL, 'admin', '2025-03-05 09:30:00'),
(8, 8, 8, 'Feeder 8', 'Type H', NULL, NULL, 'admin', '2025-03-05 09:35:00'),
(9, 9, 9, 'Feeder 9', 'Type I', NULL, NULL, 'admin', '2025-03-05 09:40:00'),
(10, 10, 10, 'Feeder 10', 'Type J', NULL, NULL, 'admin', '2025-03-05 09:45:00'),
(11, 11, 11, 'Feeder 11', 'Type K', NULL, NULL, 'admin', '2025-03-05 09:50:00'),
(12, 12, 12, 'Feeder 12', 'Type L', NULL, NULL, 'admin', '2025-03-05 09:55:00'),
(13, 13, 13, 'Feeder 13', 'Type M', NULL, NULL, 'admin', '2025-03-05 10:00:00'),
(14, 14, 14, 'Feeder 14', 'Type N', NULL, NULL, 'admin', '2025-03-05 10:05:00'),
(15, 15, 15, 'Feeder 15', 'Type O', NULL, NULL, 'admin', '2025-03-05 10:10:00'),
(16, 16, 16, 'Feeder 16', 'Type P', NULL, NULL, 'admin', '2025-03-05 10:15:00'),
(17, 17, 17, 'Feeder 17', 'Type Q', NULL, NULL, 'admin', '2025-03-05 10:20:00'),
(18, 18, 18, 'Feeder 18', 'Type R', NULL, NULL, 'admin', '2025-03-05 10:25:00'),
(19, 19, 19, 'Feeder 19', 'Type S', NULL, NULL, 'admin', '2025-03-05 10:30:00'),
(20, 20, 20, 'Feeder 20', 'Type T', NULL, NULL, 'admin', '2025-03-05 10:35:00');

CREATE TABLE `user` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_uuid` char(12) NOT NULL,
  `avatar` varchar(255) NOT NULL DEFAULT '' COMMENT '头像',
  `name` varchar(255) NOT NULL COMMENT '姓名',
  `mobile` varchar(20) NOT NULL COMMENT '手机号',
  `password` varchar(255) NOT NULL COMMENT '密码',
  `gender` tinyint unsigned NOT NULL DEFAULT 3 COMMENT '性别: 1男, 2女, 3未知',
  `birth` timestamp NOT NULL COMMENT '生日',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `udx_uuid` (`user_uuid`),
  KEY `idx_name` (`name`),
  KEY `idx_mobile` (`mobile`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
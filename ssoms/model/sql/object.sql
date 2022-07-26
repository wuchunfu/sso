CREATE TABLE `object` (
  `id` int NOT NULL AUTO_INCREMENT,
  `uuid` char(12) NOT NULL,
  `object_name` varchar(255) NOT NULL,
  `identifier` varchar(255) NOT NULL DEFAULT '',
  `key` varchar(255) NOT NULL COMMENT '操作对象的systemCode, 菜单的path, 操作的uri',
  `sort` int NOT NULL DEFAULT '0',
  `type` tinyint(1) NOT NULL DEFAULT '1' COMMENT '类型: 1系统, 2菜单, 3操作(接口)',
  `sub_type` tinyint(1) NOT NULL DEFAULT '1' COMMENT '子分类',
  `extra` varchar(1024) COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '扩展字段，建议封装成 JSON 字符串',
  `icon` varchar(255) NOT NULL DEFAULT '' COMMENT '图标',
  `status` tinyint(1) NOT NULL COMMENT '状态',
  `puuid` char(12) NOT NULL COMMENT '父级uuid',
  `top_key` varchar(255) NOT NULL DEFAULT '' COMMENT '操作对象的所属systemCode',
  `is_delete` tinyint NOT NULL DEFAULT '0' COMMENT '是否删除: 0正常, 1删除',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `udx_uuid` (`uuid`),
  UNIQUE KEY `udx_top_key` (`key`,`type`,`sub_type`,`top_key`),
  KEY `idx_object_name` (`object_name`),
  KEY `idx_key` (`key`),
  KEY `idx_status` (`status`),
  KEY `idx_type` (`type`),
  KEY `idx_sub_type` (`sub_type`),
  KEY `idx_puuid` (`puuid`),
  KEY `idx_is_delete` (`is_delete`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
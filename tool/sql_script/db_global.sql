USE db_global;

-- id池
DROP TABLE IF EXISTS `id_pool`;
CREATE TABLE `id_pool`  (
  `id` bigint(0) NOT NULL AUTO_INCREMENT,
  `invalid` int(0) NOT NULL DEFAULT 0,
  `dispatch_time` int(0) NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX(invalid)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;


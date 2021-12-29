USE db_login_server;

-- 用户表
DROP TABLE IF EXISTS `accounts_info`;
CREATE TABLE `accounts_info`  (
	`id` BIGINT NOT NULL AUTO_INCREMENT,
	`uid` BIGINT NOT NULL DEFAULT 0,
	`account` VARCHAR(255) NOT NULL DEFAULT "",
	`password` VARCHAR(255) NOT NULL DEFAULT "",
	`nick_name` VARCHAR(255) NOT NULL DEFAULT "",
	`channel` int(0) NOT NULL DEFAULT 0,
	`register_time` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`last_login_time` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`ban` INT8(0) NOT NULL DEFAULT 0,
	PRIMARY KEY (`id`) USING BTREE,
	INDEX(uid),INDEX(account),INDEX(nick_name)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;
-- 角色表
DROP TABLE IF EXISTS `role_info`;
CREATE TABLE `role_info`  (
	`id` BIGINT NOT NULL AUTO_INCREMENT,
	`uid` BIGINT NOT NULL DEFAULT 0,
	`role_id` BIGINT NOT NULL DEFAULT 0,
	`role_name` VARCHAR(255) NOT NULL DEFAULT "",
	`head` VARCHAR(255) NOT NULL DEFAULT "",
	`create_time` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`last_login_time` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`ban` INT8(0) NOT NULL DEFAULT 0,
	PRIMARY KEY (`id`) USING BTREE,
	INDEX(uid),INDEX(role_id),INDEX(role_name)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;


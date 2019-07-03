-- 创建数据库
DROP DATABASE IF EXISTS interview;
CREATE DATABASE interview;
USE interview;

-- 创建表
DROP TABLE IF EXISTS question;

-- question 表

CREATE TABLE IF NOT EXISTS question (
  `id`          INT UNSIGNED AUTO_INCREMENT,                   -- primary key
  `title`       VARCHAR(255) NOT NULL DEFAULT '',              -- feed name`s alias
  `difficulty`  INT(8) DEFAULT 0,
  `summary`     VARCHAR(1000) Not NULL DEFAULT '',             -- 概述
  `content`     LONGTEXT DEFAULT NULL,   -- 内容
  `created_at`     DATETIME NOT NULL,
  `updated_at`     DATETIME NOT NULL,
   PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
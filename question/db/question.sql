-- 创建数据库
DROP DATABASE IF EXISTS interview_qu;
CREATE DATABASE interview_qu;
USE interview_qu;

-- 删除表
DROP TABLE IF EXISTS question;
DROP TABLE IF EXISTS topic;
DROP TABLE IF EXISTS qu_topic;

-- question 表
CREATE TABLE IF NOT EXISTS question (
  `id`              INT UNSIGNED AUTO_INCREMENT,                   -- primary key
  `title`           VARCHAR(255) NOT NULL DEFAULT '',              -- feed name`s alias
  `difficulty`      INT(8) DEFAULT 0,
  `summary`         VARCHAR(1000) Not NULL DEFAULT '',             -- 概述
  `content`         LONGTEXT DEFAULT NULL,   -- 内容
  `created_at`      DATETIME NOT NULL,
  `updated_at`      DATETIME NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- topic 表
CREATE TABLE IF NOT EXISTS topic (
  `id`              INT UNSIGNED AUTO_INCREMENT,                   -- primary key
  `slug`            VARCHAR(255) NOT NULL DEFAULT '',              -- feed name`s alias
  `created_at`      DATETIME NOT NULL,
  `updated_at`      DATETIME NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- question topic connect table
CREATE TABLE IF NOT EXISTS qu_topic (
  `id`              INT UNSIGNED AUTO_INCREMENT,                   -- primary key
  `qu_id`           INT(64) NOT NULL,                              -- question.id
  `topic_id`        INT(64) NOT NULL,                              -- topioc.id
  `created_at`      DATETIME NOT NULL,
  `updated_at`      DATETIME NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
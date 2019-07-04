DROP DATABASE IF EXISTS interview_ans;
CREATE DATABASE interview_ans;
USE interview_ans;

-- 删除表
DROP TABLE IF EXISTS ans;

-- answer 表
CREATE TABLE IF NOT EXISTS ans (
  `id`              INT UNSIGNED AUTO_INCREMENT,                   -- primary key
  `qu_id`           INT(64) NOT NULL,                              -- question.id
  `content`         LONGTEXT DEFAULT NULL,   -- 内容
  `created_at`      DATETIME NOT NULL,
  `updated_at`      DATETIME NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
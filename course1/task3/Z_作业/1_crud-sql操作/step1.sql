CREATE TABLE `go-study`.students
(
    `id`          bigint(20) NOT NULL AUTO_INCREMENT,
    `name`        varchar(255)       DEFAULT NULL COMMENT '名称',
    `age`         int(11)        DEFAULT NULL COMMENT '年龄',
    `grade`       varchar(20)        DEFAULT NULL COMMENT '年级',
    `last_update` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后更新时间',
    PRIMARY KEY (`id`) USING BTREE,
    KEY           `idx_name` (`name`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT = 1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci
ROW_FORMAT = DYNAMIC COMMENT ='学生表';


INSERT INTO `go-study`.students
(name, age, grade)
VALUES('张三', 20, '三年级');


select * from students s
where age>18;

update students
set grade = '四年级'
where name ='张三';

delete FROM students
where age < 15;
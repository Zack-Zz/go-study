CREATE TABLE `go-study`.employees
(
    `id`          bigint(20) NOT NULL AUTO_INCREMENT,
    `name`        varchar(255)       DEFAULT NULL COMMENT '名称',
    `department`  varchar(50)        DEFAULT NULL COMMENT '年级',
    `salary`      decimal(18,2)        DEFAULT NULL COMMENT '薪资',
    `last_update` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后更新时间',
    PRIMARY KEY (`id`) USING BTREE,
    KEY           `idx_name` (`name`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT = 1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci
ROW_FORMAT = DYNAMIC COMMENT ='员工表';

INSERT INTO `go-study`.employees
(name, department, salary)
VALUES('张三', '开发部', 25000);

INSERT INTO `go-study`.employees
(name, department, salary)
VALUES('李四', '开发部', 35000);

INSERT INTO `go-study`.employees
(name, department, salary)
VALUES('赵五', '开发部', 45000);

INSERT INTO `go-study`.employees
(name, department, salary)
VALUES('行六', '技术部', 26000);

INSERT INTO `go-study`.employees
(name, department, salary)
VALUES('法七', '技术部', 36000);

INSERT INTO `go-study`.employees
(name, department, salary)
VALUES('笑八', '技术部', 46000);
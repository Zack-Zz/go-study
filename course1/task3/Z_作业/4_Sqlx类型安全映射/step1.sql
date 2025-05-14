CREATE TABLE `go-study`.books
(
    `id`          bigint(20) NOT NULL AUTO_INCREMENT,
    `title`       varchar(255)       DEFAULT NULL COMMENT '书名',
    `author`      varchar(50)        DEFAULT NULL COMMENT '作者',
    `price`       decimal(18, 2)     DEFAULT NULL COMMENT '价格',
    `last_update` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后更新时间',
    PRIMARY KEY (`id`) USING BTREE,
    KEY           `idx_title` (`title`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT = 1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci
ROW_FORMAT = DYNAMIC COMMENT ='书籍表';

INSERT INTO `go-study`.books
    (title, author, price)
VALUES ('A', 'a', 25);

INSERT INTO `go-study`.books
    (title, author, price)
VALUES ('B', 'a', 35);

INSERT INTO `go-study`.books
    (title, author, price)
VALUES ('C', 'a', 45);

INSERT INTO `go-study`.books
    (title, author, price)
VALUES ('D', 'b', 26);

INSERT INTO `go-study`.books
    (title, author, price)
VALUES ('E', 'b', 36);

INSERT INTO `go-study`.books
    (title, author, price)
VALUES ('F', 'b', 46);
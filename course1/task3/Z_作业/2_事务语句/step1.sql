CREATE TABLE `go-study`.accounts
(
    `id`          bigint(20) NOT NULL AUTO_INCREMENT,
    `balance`     decimal(18, 2)     DEFAULT NULL COMMENT '余额',
    `last_update` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后更新时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT = 1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci
ROW_FORMAT = DYNAMIC COMMENT ='账户表';

CREATE TABLE `go-study`.transactions
(
    `id`              bigint(20) NOT NULL AUTO_INCREMENT,
    `from_account_id` bigint(20)       DEFAULT NULL COMMENT '转出账户id',
    `to_account_id`   bigint(20)        DEFAULT NULL COMMENT '转入账户id',
    `amount`          decimal(18, 2)     DEFAULT NULL COMMENT '转账金额',
    `last_update`     timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后更新时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT = 1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci
ROW_FORMAT = DYNAMIC COMMENT ='账户转账表';


## 创建存储过程 - TSql
CREATE PROCEDURE transfer_money(IN from_id INT, IN to_id INT, IN amount DECIMAL(10,2))
BEGIN
    DECLARE from_balance DECIMAL(10,2);

    DECLARE EXIT HANDLER FOR SQLEXCEPTION
BEGIN
ROLLBACK;
END;

START TRANSACTION;

-- 1. 查询转出账户余额
SELECT balance INTO from_balance FROM accounts WHERE id = from_id FOR UPDATE;

-- 2. 判断是否足够
IF from_balance < amount THEN
        ROLLBACK;
ELSE
    -- 3. 更新账户余额
    UPDATE accounts SET balance = balance - amount WHERE id = from_id;
    UPDATE accounts SET balance = balance + amount WHERE id = to_id;

    -- 4. 插入交易记录
    INSERT INTO transactions (from_account_id, to_account_id, amount)
    VALUES (from_id, to_id, amount);

COMMIT;
END IF;
END$$

CALL transfer_money(1, 2, 100);

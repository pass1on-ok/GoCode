CREATE TABLE `category` (
    `id`                    BIGINT          NOT NULL AUTO_INCREMENT,
    `name`                  VARCHAR(255)    NOT NULL,
    `dtm_crt`               TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `dtm_upd`               TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
);

CREATE TABLE `course` (
    `id`                    BIGINT          NOT NULL AUTO_INCREMENT,
    `category_id`           BIGINT          NOT NULL,
    `name`                  VARCHAR(255)    NOT NULL,
    `detail`                TEXT,
    `price`                 DECIMAL(10,2)   NOT NULL,
    `picture`               VARCHAR(1024),
    `dtm_crt`               TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `dtm_upd`               TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
);

CREATE TABLE `transaction` (
    `id`                    BIGINT          NOT NULL AUTO_INCREMENT,
    `user_id`               BIGINT          NOT NULL,
    `course_id`             BIGINT          NOT NULL,
    `price`                 DECIMAL(10,2)   NOT NULL,
    `dtm_crt`               TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `dtm_upd`               TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
);

CREATE TABLE `user` (
    `id`                    BIGINT          NOT NULL AUTO_INCREMENT,
    `email`                 VARCHAR(255)    NOT NULL UNIQUE,
    `password`              VARCHAR(255)    NOT NULL,
    `name`                  VARCHAR(255)    NOT NULL,
    `picture`               VARCHAR(1024),
    `role`                  VARCHAR(255)    NOT NULL,
    `deleted`               TINYINT(1)      NOT NULL,
    `dtm_crt`               TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `dtm_upd`               TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
);
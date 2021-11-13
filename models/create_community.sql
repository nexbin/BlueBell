USE blue_bell;

DROP TABLE IF EXISTS `community`;

CREATE TABLE community
(
    `id`             int(11)                                 NOT NULL AUTO_INCREMENT,
    `community_id`   int(10) UNSIGNED                        NOT NULL,
    `community_name` varchar(128) COLLATE utf8mb4_general_ci NOT NULL,
    `introduction`   varchar(256) COLLATE utf8mb4_general_ci NOT NULL,
    `create_time`    timestamp                               NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time`    timestamp                               NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_community_id` (community_id),
    UNIQUE KEY `idx_community_name` (`community_name`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_general_ci;

# INSERT INTO `community` VALUES ('1','1','Go','Golang','2021-11-13 09:59:59','2021-11-13 09:59:59');
# INSERT INTO `community` VALUES ('2','2','LeetCode','刷题网站','2021-11-13 09:59:59','2021-11-13 09:59:59');
# INSERT INTO `community` VALUES ('3','3','CS:GO','Rush B','2021-11-13 09:59:59','2021-11-13 09:59:59');
# INSERT INTO `community` VALUES ('4','4','Lol','欢迎来到英雄联盟','2021-11-13 09:59:59','2021-11-13 09:59:59');




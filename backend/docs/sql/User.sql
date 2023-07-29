CREATE TABLE `user`
(
    `id`         INT(11)     NOT NULL,
    `created_at` DATETIME    DEFAULT NULL ,
    `name`       VARCHAR(32) NOT NULL,
    `avatarUrl`  VARCHAR(255),
    `phone`      CHAR(13),
    PRIMARY KEY (`id`)
);

-- ----------------------------
-- Records of User
-- ----------------------------
INSERT INTO `user` VALUES (1, NULL, '林涵', 'https://lh3.googleusercontent.com/a/AGNmyxYZi8k72dyr48U9m3vS9V8qd6ibNF70LMuI7wsU=s96-c', NULL);
INSERT INTO `user` VALUES (2, NULL, 'zql', 'https://zql-oss1.oss-cn-nanjing.aliyuncs.com/notes/avatat.png', 18856264667);

CREATE TABLE `wx_user` (
                           `id`  int(11) UNSIGNED NOT NULL AUTO_INCREMENT ,
                           `open_id`  char(28) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '小程序用户openid' ,
                           `nickname`  varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '用户昵称' ,
                           `avatar`  varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '用户头像' ,
                           `gender`  tinyint(1) NULL DEFAULT 0 COMMENT '性别   0 男  1  女  2 人妖' ,
                           `country`  varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '所在国家' ,
                           `province`  varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '省份' ,
                           `city`  varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '城市' ,
                           `language`  varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL ,
                           `ctime`  timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间' ,
                           `update_time` timestamp NULL  DEFAULT CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP,
                           PRIMARY KEY (`id`),
                           UNIQUE KEY `idx_username` (`open_id`) USING BTREE
)
    ENGINE=MyISAM
    DEFAULT CHARACTER SET=utf8 COLLATE=utf8_general_ci
    AUTO_INCREMENT=1
    CHECKSUM=0
    ROW_FORMAT=DYNAMIC
    DELAY_KEY_WRITE=0;




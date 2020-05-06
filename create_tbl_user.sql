CREATE TABLE `tbl_user` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `user_name` varchar(64) NOT NULL DEFAULT '' COMMENT 'username',
    `user_pwd` varchar(256) NOT NULL DEFAULT '' COMMENT 'user\'s encoded password',
    `email` varchar(64) DEFAULT '' COMMENT 'email',
    `phone` varchar(128) DEFAULT '' COMMENT 'phone number',
    `email_validated` tinyint(1) DEFAULT 0 COMMENT 'whether the email is validated',
    `phone_validated` tinyint(1) DEFAULT 0 COMMENT 'whether the phone number is validated',
    `signup_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT 'sign up date',
    `last_active` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'last active timestamp',
    `profile` text COMMENT 'user properties',
    `status` int(11) NOT NULL DEFAULT '0' COMMENT 'account status(enabled/disabled/locked/marked as deleted, etc.)',
    PRIMARY KEY(`id`),
    UNIQUE KEY `idx_phone` (`phone`),
    KEY `idx_status` (`status`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4;
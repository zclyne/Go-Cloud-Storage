CREATE TABLE `tbl_user_file` (
    `id` int(11) NOT NULL PRIMARY KEY AUTO_INCREMENT,
    `user_name` varchar(64) NOT NULL,
    `file_sha1` varchar(64) NOT NULL DEFAULT '' COMMENT 'file hash',
    `file_size` bigint(20) DEFAULT '0' COMMENT 'file size',
    `file_name` varchar(256) NOT NULL DEFAULT '' COMMENT 'file name',
    `upload_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT 'upload time',
    `last_update` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'last update time',
    `status` int(11) NOT NULL DEFAULT '0' COMMENT 'file status (0 for normal, 1 for deleted, 2 for disabled',
    UNIQUE KEY `idx_user_file` (`user_name`, `file_sha1`),
    KEY `idx_status` (`status`),
    KEY `idx_user_id` (`user_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
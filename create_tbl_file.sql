CREATE TABLE `tbl_file` (
  `id` int NOT NULL AUTO_INCREMENT,
  `file_sha1` char(40) NOT NULL DEFAULT '' COMMENT 'file hash',
  `file_name` varchar(256) NOT NULL DEFAULT '' COMMENT 'filename',
  `file_size` bigint DEFAULT '0' COMMENT 'file size',
  `file_addr` varchar(1024) NOT NULL DEFAULT '' COMMENT 'file storage location',
  `create_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT 'create time',
  `update_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'update time',
  `status` int NOT NULL DEFAULT '0' COMMENT 'status(enabled/disabled/deleted, etc.)',
  `ext1` int DEFAULT '0' COMMENT 'extention 1',
  `ext2` text COMMENT 'extention 2',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_file_hash` (`file_sha1`),
  KEY `idx_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
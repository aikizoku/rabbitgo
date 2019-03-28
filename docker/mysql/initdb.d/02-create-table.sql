CREATE TABLE IF NOT EXISTS `sample`.`sample` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL,
  `enabled` tinyint(1) NOT NULL DEFAULT 1,
  `created_at` bigint(20) NOT NULL,
  `updated_at` bigint(20) NOT NULL,
  PRIMARY KEY (`id`),
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4
COMMENT 'sample';

-- 参考
-- CREATE TABLE IF NOT EXISTS `db`.`table` (
--   `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
--   `fuga` varchar(1024) NOT NULL COMMENT 'fugafuga',
--   `aaaa` varchar(255) NOT NULL DEFAULT 'bbb',
--   `enabled` tinyint(1) NOT NULL DEFAULT 1,
--   `dddd` enum('a', 'b', 'c', 'd') NOT NULL DEFAULT 'a',
--   `notes` text,
--   `created_at` bigint(20) NOT NULL,
--   `updated_at` bigint(20) NOT NULL,
--   PRIMARY KEY (`id`),
--   KEY idx_fuga_and_enabled (`fuga`, `enabled`)
-- ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4
-- COMMENT 'comment';

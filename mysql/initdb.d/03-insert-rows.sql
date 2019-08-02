INSERT INTO `sample`.`sample` (
  id, name, enabled, created_at, updated_at
)
VALUES
  (1, "hoge", true, ROUND(UNIX_TIMESTAMP(CURTIME(4)) * 1000), ROUND(UNIX_TIMESTAMP(CURTIME(4)) * 1000));

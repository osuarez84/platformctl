CREATE TABLE IF NOT EXISTS default.my_table (
  id UInt32,
  name String
) ENGINE = MergeTree()
ORDER BY id;

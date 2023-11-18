#!/bin/bash
set -e

clickhouse client -n <<-EOSQL
  CREATE TABLE IF NOT EXISTS default.my_table (
    id UInt32,
    name String
  ) ENGINE = MergeTree()
  ORDER BY id;

  INSERT INTO default.my_table
  SELECT number as id, concat('Name', toString(number)) as name
  FROM numbers(50);
EOSQL
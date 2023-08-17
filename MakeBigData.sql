USE test_db;
DROP TABLE `tests`;
CREATE TABLE `tests` (
  `content` text NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

INSERT INTO `tests`
    (`content`)
VALUES
    ('xxx'),
    ('xxx'),
    ('xxx'),
    ('xxx'),
    ('xxx'),
    ('xxx'),
    ('xxx'),
    ('xxx'),
    ('xxx'),
    ('xxx');

INSERT INTO `tests` (`content`)
SELECT `t1`.`content` FROM `tests` t1, `tests` t2, `tests` t3, `tests` t4, `tests` t5, `tests` t6;

SELECT COUNT(*) FROM tests;
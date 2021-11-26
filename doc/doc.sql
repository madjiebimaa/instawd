CREATE TABLE quote (
    id VARCHAR(30) NOT NULL,
    content VARCHAR(300) NOT NULL,
    author_id VARCHAR(30) NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (author_id) REFERENCES author(id)
);
-- 
CREATE TABLE quote_tag (
    id VARCHAR(30) NOT NULL,
    name VARCHAR(30) NOT NULL
);
-- 
CREATE TABLE author (
    id VARCHAR(30) NOT NULL,
    name VARCHAR(100) NOT NULL,
    link VARCHAR(3000),
    bio VARCHAR(1000),
    description VARCHAR(1000),
    quote_count INT DEFAULT 0,
    PRIMARY KEY (id)
);
-- 
WITH quote_numbered AS (
    SELECT id,
        author_id,
        content,
        row_number() over() AS rn
    FROM quote
)
SELECT id,
    author_id,
    content
FROM quote_numbered
WHERE rn = 1;
-- 
SHOW CREATE table author;
--
ALTER TABLE author
ADD slug VARCHAR(100) NOT NULL;
-- 
UPDATE author
SET link = "test"
WHERE id = "test";
-- 
DELETE FROM quote
WHERE content = "asd";
-- 
SELECT id,
    name,
    quote_count
FROM author
WHERE quote_count = (
        SELECT MAX(quote_count)
        FROM author
    );
-- 
SELECT id,
    author_id
FROM quote
ORDER BY id
LIMIT 10 OFFSET 20;
--
SELECT id,
    author_id,
    content,
    CHAR_LENGTH(content) AS content_length
FROM quote
HAVING content_length > 10
    AND content_length < 30;
--
SELECT id,
    author_id,
    content,
    CHAR_LENGTH(content) AS content_length
FROM quote
SHOW DATABASES;
CREATE TABLE quote (
    id VARCHAR(30) NOT NULL,
    content VARCHAR(300) NOT NULL,
    author_id VARCHAR(30) NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (author_id) REFERENCES author(id)
);
CREATE TABLE quote_tag (
    id VARCHAR(30) NOT NULL,
    name VARCHAR(30) NOT NULL
);
CREATE TABLE author (
    id VARCHAR(30) NOT NULL,
    name VARCHAR(100) NOT NULL,
    link VARCHAR(3000),
    bio VARCHAR(1000),
    description VARCHAR(1000),
    quote_count INT DEFAULT 0,
    PRIMARY KEY (id)
);
USE go_random_quotes;
SHOW TABLES;
SHOW DATABASES;
select *
from quote_tag;
SHOW TABLES;
select *
from author;
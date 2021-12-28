DROP TABLE IF EXISTS `authors`;
CREATE TABLE `authors` (
    `id` varchar(30) NOT NULL,
    `name` varchar(100) NOT NULL,
    `link` varchar(3000) DEFAULT NULL,
    `bio` varchar(1000) DEFAULT NULL,
    `description` varchar(1000) DEFAULT NULL,
    `quote_count` int DEFAULT '0',
    `slug` varchar(100) NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB;
-- 
DROP TABLE IF EXISTS `quotes`;
CREATE TABLE `quotes` (
    `id` varchar(30) NOT NULL,
    `author_id` varchar(30) NOT NULL,
    `content` varchar(300) NOT NULL,
    PRIMARY KEY (`id`),
    KEY `author_id` (`author_id`),
    FOREIGN KEY (`author_id`) REFERENCES `authors` (`id`)
) ENGINE = InnoDB;
-- 
DROP TABLE IF EXISTS `quote_tags`;
CREATE TABLE `quote_tags` (
    `id` varchar(30) NOT NULL,
    `name` varchar(30) NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB;
-- 
DROP TABLE IF EXISTS `quotes_to_quote_tags`;
CREATE TABLE `quotes_to_quote_tags` (
    `id` varchar(30) NOT NULL,
    `quote_id` varchar(30) NOT NULL,
    `quote_tag_id` varchar(30) NOT NULL,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`quote_id`) REFERENCES `quotes` (`id`),
    FOREIGN KEY (`quote_tag_id`) REFERENCES `quote_tags` (`id`)
) ENGINE = InnoDB;
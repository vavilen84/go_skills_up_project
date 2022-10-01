CREATE TABLE posts
(
    id          INT UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
    description text DEFAULT NULL,
    created_at datetime NOT NULL,
    updated_at datetime NOT NULL
) ENGINE=InnoDB CHARSET=utf8;

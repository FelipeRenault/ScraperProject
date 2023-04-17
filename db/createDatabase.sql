USE pelando;

CREATE TABLE products (
    id INT NOT NULL AUTO_INCREMENT,
    url VARCHAR(200),
    title VARCHAR(200),
    image VARCHAR(200),
    price INT,
    description TEXT,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE (url)
);

SET character_set_client = utf8;
SET character_set_connection = utf8;
SET character_set_results = utf8;
SET collation_connection = utf8_general_ci;

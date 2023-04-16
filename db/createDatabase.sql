USE pelando;

CREATE TABLE products (
    id int NOT NULL AUTO_INCREMENT,
    url varchar(200),
    title varchar(200),
    image varchar(200),
    price int,
    description varchar(200),
    created_at datetime not null DEFAULT CURRENT_TIMESTAMP,
    updated_at datetime not null DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE (url)
);

SET character_set_client = utf8;
SET character_set_connection = utf8;
SET character_set_results = utf8;
SET collation_connection = utf8_general_ci;

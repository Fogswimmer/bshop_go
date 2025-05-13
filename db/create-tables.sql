CREATE TABLE book (
    id SERIAL NOT NULL,
    title VARCHAR(128) NOT NULL,
    release_date INT NOT NULL,
    summary TEXT DEFAULT NULL,
    price DECIMAL(5,2) NOT NULL,
    cover VARCHAR(128) NOT NULL,
    author_id INT DEFAULT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (author_id) REFERENCES author(id)
    CREATE INDEX book_author_id_idx ON book (author_id)
);

CREATE TABLE author (
    id SERIAL NOT NULL,
    firstname VARCHAR(128) NOT NULL,
    lastname VARCHAR(128) NOT NULL,
    birthday DATE NOT NULL,
    PRIMARY KEY (id)
);

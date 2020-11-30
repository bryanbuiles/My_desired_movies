CREATE TABLE movie
(
    id           varchar(100) NOT NULL,
    title        varchar(100) NOT NULL,
    caste       varchar(250) NULL,
    release_date DATE         NULL,
    genre        varchar(100) NULL,
    director     varchar(100) NULL,
    PRIMARY KEY(id)
);


CREATE TABLE users
(
    id       varchar(100) NOT NULL,
    username varchar(100) NOT NULL UNIQUE,
    password varchar(100) NOT NULL,
    PRIMARY KEY(id)
);


CREATE TABLE wish_list
(
    user_id  varchar(100) NOT NULL,
    movie_id varchar(100) NOT NULL,
    comment  varchar(300) NULL,
    CONSTRAINT wish_list_PK PRIMARY KEY (user_id, movie_id),
    CONSTRAINT wish_list_FK FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    CONSTRAINT wish_list_FK_1 FOREIGN KEY (movie_id) REFERENCES movie(id)
);

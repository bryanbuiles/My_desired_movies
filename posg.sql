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

INSERT INTO movie (id,title,caste,release_date,genre,director)
VALUES ('1b070506-342a-11eb-adc1-0242ac120002', 'Cape Fear', 'Robert de Niro, Gregory Peck', TO_DATE('13-11-1991', 'DD MM YYYY'), 'suspense, drama', 'Martin scorses');

INSERT INTO movie (id,title,caste,release_date,genre,director)
VALUES ('229a7fa4-3426-11eb-adc1-0242ac120002', 'Blade Runner', 'Harrison Ford, Sean Young', TO_DATE('15-07-1982', 'DD MM YYYY'), 'cs fiction, drama', 'Ridley Scott');


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

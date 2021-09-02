SET FOREIGN_KEY_CHECKS = 0;

CREATE TABLE `user` (
    userid int(11) NOT NULL AUTO_INCREMENT PRIMARY KEY,
    username varchar(255) NULL ,
    email varchar(255) NULL ,
    address text NULL,
    password text NULL,
    token text NULL
);
INSERT INTO user (userid, username, email, address, password)
VALUES 
(1, "iqbal", "iqbal@mail.id", "Jakarta Pusat", "$2y$12$puwDDoIzG34nTM2sth8JhuDW0nbEdiYVqWVOGAPuBK/zcd.mrer86", ""),
(2, "ravi", "ravi@mail.id", "Jakarta Timur", "$2y$12$puwDDoIzG34nTM2sth8JhuDW0nbEdiYVqWVOGAPuBK/zcd.mrer86", "");
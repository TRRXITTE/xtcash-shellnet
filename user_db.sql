-- setup user database
CREATE database users;
\c users;
CREATE TABLE accounts (
IH char(64) NOT NULL,
Verifier char(844) NOT NULL,
Username varchar(64) NOT NULL UNIQUE,
ID  SERIAL PRIMARY KEY,
address char(98) NOT NULL);
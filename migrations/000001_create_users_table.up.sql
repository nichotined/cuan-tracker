CREATE TABLE IF NOT EXISTS users (
   id   	 SERIAL        PRIMARY KEY ,
   first_name    VARCHAR (50) NOT NULL,
   last_name     VARCHAR (50) NOT NULL,
   phone_number  VARCHAR(100) UNIQUE NOT NULL
);